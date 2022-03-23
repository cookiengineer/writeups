
import fs from 'fs';



const buffer = fs.readFileSync('flag');
const bytes  = [];
const lines  = buffer.toString('utf8').trim().split('\n');


lines.forEach((line) => {
	bytes.push(parseInt(line, 10));
});

console.log(Buffer.from(bytes).toString('utf8'));


