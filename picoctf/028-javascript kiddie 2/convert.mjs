
import child_process from 'child_process';
import fs            from 'fs';



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

	let values = [];

	for (let i = 0; i < shifted_table[column].length; i++) {

		if (shifted_table[column][i] === png_header[column]) {
			if (i >= 0 && i <= 9) {
				values.push(i);
			}
		}

	}

	shift.push(values);

}

fs.writeFileSync('bytes.bin', Buffer.from(shifted_bytes));



const toCombinations = function(array) {

	if (array.length === 0) return [[]];

	let result            = [];
	let [ first, ...rest] = array;
	let remaining         = toCombinations(rest);

	first.forEach((number) => {

		remaining.forEach((smaller) => {
			result.push([ number ].concat(smaller));
		});

	});

	return result;

};

let permutations = toCombinations(shift);

console.log('Found ' + permutations.length + ' possible combinations...');

if (permutations.length > 0) {

	permutations.forEach((permutation, p) => {

		let key = permutation.map((v) => {
			return '' + v + v;
		}).join('');

		console.log('shift code #' + p + ' is: ' + key);

		let result = [];
		for (let i = 0; i < 16; i++) {

			let shifter = Number(key.slice((i * 2), (i * 2) + 1));
			for (let j = 0; j < (shifted_bytes.length / 16); j++) {
				result[(j * 16) + i] = shifted_bytes[(((j + shifter) * 16) % shifted_bytes.length) + i];
			}

		}

		while (result[result.length - 1] === 0) {
			result = result.slice(0, result.length - 1);
		}

		fs.writeFileSync('cracked/permutation-' + p +'.png', Buffer.from(result));


		let flag = null;
		try {
			flag = child_process.spawnSync('zbarimg', [
				process.cwd() + '/cracked/permutation-' + p + '.png'
			], {
				cwd: process.cwd() + '/cracked'
			}).stdout;
		} catch (err) {
			flag = null;
		}

		if (flag !== null && flag.length > 0) {
			console.log('Found Flag! ' + flag.toString('utf8'));
		}

	});

}

