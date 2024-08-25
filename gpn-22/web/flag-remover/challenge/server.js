const express = require('express');
const puppeteer = require('puppeteer');

const flag = process.env.FLAG || require('fs').readFileSync('flag', 'utf8') || "GPNCTF{fake_flag}";

const app = express();
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
    res.send(`
        <h1>Overriding</h1>
        <form action="/chal" method="post">
            <input type="text" name="flag" placeholder="Flag" id="a">
            <input type="text" name="html" placeholder="HTML" id="b">
            <button type="submit" id="c">Submit to /chal</button>
        </form>
        <hr>
        <form action="/admin" method="post">
            <input type="text" name="solve" placeholder="HTML">
            <button type="submit">Submit to /admin</button>
        </form>
    `);
});

const DOMPurify = require('dompurify')(new (require('jsdom').JSDOM)('').window);

app.post('/chal', (req, res) => {
    let { flag, html } = req.body;
    // don't xss
    html = DOMPurify.sanitize(html);
    // don't close/open the body tag
    html = html.replace(/<body>|<\/body>/g, '');
    res.setHeader("Content-Security-Policy", "default-src 'none'; script-src 'self';");
    res.send(`
        <head>
            <script async defer src="/removeFlag.js"></script>
        </head>
        <body>
            <div class="flag">${flag}</div>
            ${html}
        </body>
    `);
});

app.get('/removeFlag.js', (req, res) => {
    res.type('.js');
    res.send(`try {
        let els = document.body.querySelectorAll('.flag');
        if (els.length !== 1) throw "nope";
        els[0].remove();
    } catch(e) { location = 'https://duckduckgo.com/?q=no+cheating+allowed&iax=images&ia=images' }`);
});

app.post('/admin', async (req, res) => {
    try {
        const { solve } = req.body;
        const browser = await puppeteer.launch({ executablePath: process.env.BROWSER, args: ['--no-sandbox'] });
        // go to localhost:1337, type flag and html, click submit
        const page = await browser.newPage();
        await page.goto('http://localhost:1337/');
        await page.type('input[name="flag"]', flag.trim());
        await page.type('input[name="html"]', solve.trim());
        await page.click('button[type="submit"]');
        await new Promise(resolve => setTimeout(resolve, 5000));
        // make sure js execution isn't blocked
        await page.waitForFunction('true')
        // take a screenshot
        const screenshot = await page.screenshot({ encoding: 'base64' });
        await browser.close();
        res.send(`<img src="data:image/png;base64,${screenshot}" />`);
    } catch(e) {console.error(e); res.send("internal error :( pls report to admins")}
});

app.listen(1337, () => console.log('listening on http://localhost:1337'));
