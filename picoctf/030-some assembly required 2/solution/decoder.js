
let buffer_encoded  = Buffer.from('eGFrZ0tcTnNsPDg/bm1pOjxpOzBqOTo7P25tOGk9MD8/Oj1uam49OXUAAA==', 'base64');
let buffer_original = buffer_encoded.map((value) => {

	// from $2
	// return (($1(1024 | 0, 1072 | 0) | 0 | 0) != (0 | 0) ^ -1 | 0) & 1 | 0 | 0;

	return value ^ 8;

});

console.log(buffer_original.toString('utf8'));

