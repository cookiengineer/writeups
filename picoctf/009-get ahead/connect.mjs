
import fs   from 'fs';
import http from 'http';

const VERBS = [
	'CONNECT',
	'DELETE',
	'GET',
	'HEAD',
	'OPTIONS',
	'PATCH',
	'POST',
	'PUT',
	'TRACE',
	'UPDATE'
];



VERBS.forEach((verb) => {

	let request = http.request('http://mercury.picoctf.net:45028/', {
		method: verb
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);
			let headers = Buffer.from(JSON.stringify(response.headers, null, '\t'), 'utf8');

			fs.writeFileSync('flags/' + verb + '.payload', payload);
			fs.writeFileSync('flags/' + verb + '.headers', headers);

		});

	});

	request.on('error', (err) => {
		console.error(err);
		console.error('ERROR requesting "' + verb + '"');
	});

	request.end();

});

