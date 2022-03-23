
import crypto from 'crypto';

const KEY_PREFIX   = "picoCTF{1n_7h3_|<3y_of_"
const KEY_SUFFIX   = "}";
const KEY_USERNAME = 'SCHOFIELD';
const SHA_USERNAME = crypto.createHash('sha256').update(KEY_USERNAME).digest('hex');


let indexes = [ 4, 5, 3, 6, 2, 7, 1, 8 ];
let strings = [];

strings.push(KEY_PREFIX);

indexes.forEach((index) => {
	strings.push(SHA_USERNAME[index]);
});

strings.push(KEY_SUFFIX);


console.log(strings.join(''));

