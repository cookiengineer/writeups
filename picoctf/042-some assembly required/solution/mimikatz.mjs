
import fs from 'fs';

const buffer = fs.readFileSync('./ZoRd23o0wd');
const imports = {};

const CHARSET = '_abcdefghijklmnopqrstuvwxyz0123456789';

function logIt(whatever) {

	let keys = Object.keys(whatever);

	console.log(keys);

	for (let k = 0; k < keys.length; k++) {

		console.log(keys[k], whatever[keys[k]])

	}

}

WebAssembly.instantiate(buffer, imports).then((global) => {

	const check_flag = global.instance.exports.check_flag;
	const memory     = Buffer.from(global.instance.exports.memory.buffer);

	const search = Buffer.from('picoCTF', 'utf8');

	for (let m = 0; m < memory.length; m++) {

		if (memory[m] == search[0]) {

			console.log(m);

		}

	}

	// picoCTF{32 hex}
	// 41 characters

	// let buffer = new Array(32);
	// let result = 0;

	// buffer.fill('_');

	// while (result === 0) {

	// 	for (let b = 0; b < buffer.length; b++) {

	// 		for (let c = 0; c < CHARSET.length; c++) {

	// 			buffer[b] = CHARSET[c];
	// 			result = check_flag('picoCTF{' + buffer.join('') + '}');

	// 		}

	// 	}

	// }

	// console.log('Flag is probably?', 'picoCTF{' + buffer.join('') + '}');

});

