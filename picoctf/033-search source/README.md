
# Search Source

Never hide secrets in the client-side code.


# Problem

The challenge is quite easy, you need to scrape the
website and its assets; and then simply grep for the
flag inside the folder(s).

# Solution

```bash
wget --convert-links --page-requisistes --html-extension --span-hosts -e robots=off "http://linktochallenge";
grep -R "picoCTF{" ./
```
