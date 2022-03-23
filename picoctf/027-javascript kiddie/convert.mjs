
import fs from 'fs';


let shifted_bytes = fs.readFileSync('bytes').toString('utf8').trim().split(' ').map((v) => parseInt(v, 10));
let shifted_table = [
	[], [], [], [], [], [], [], [],
	[], [], [], [], [], [], [], []
];
let shift         = [];
let png_header    = [
	0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A,
	0x00, 0x00, 0x00, 0x0D, 0x49, 0x48, 0x44, 0x52
];

shifted_bytes.forEach((value, b) => {
	shifted_table[b % 16].push(value);
});

for (let column = 0; column < shifted_table.length; column++) {

	let index = shifted_table[column].indexOf(png_header[column]);

	if (index >= 0 && index <= 9) {
		shift.push(index);
	} else {
		console.error('Fatal Error: Cannot crack the shift code!');
		break;
	}

}

fs.writeFileSync('bytes.bin', Buffer.from(shifted_bytes));

if (shift.length === 16) {

	// shift code was: 6696705967835463
	console.log('password / shift code is: ' + shift.join(''));


	// from the index.html
	let key    = shift.join('');
	let result = [];
	for (let i = 0; i < 16; i++) {

		let shifter = key.charCodeAt(i) - 48;
		for (let j = 0; j < (shifted_bytes.length / 16); j++) {
			result[(j * 16) + i] = shifted_bytes[(((j + shifter) * 16) % shifted_bytes.length) + i];
		}

	}

	while (result[result.length - 1] === 0) {
		result = result.slice(0, result.length - 1);
	}

	fs.writeFileSync('bytes.png', Buffer.from(result));

}

