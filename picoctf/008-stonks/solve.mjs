
import fs from 'fs';


const buffer = Buffer.from(fs.readFileSync('flag').toString('utf8'), 'hex');
const bytes  = [];


console.log(buffer);
console.log(buffer.toString('ascii'));

for (let b = 2 + 40, bl = buffer.length; b < bl; b += 4) {

	let chunk = [
		buffer[b + 3],
		buffer[b + 2],
		buffer[b + 1],
		buffer[b + 0]
	];

	bytes.push.apply(bytes, chunk);

}

console.log(Buffer.from(bytes).toString('ascii'));

