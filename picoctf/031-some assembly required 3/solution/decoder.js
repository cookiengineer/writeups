
let buffer1_encoded = Buffer.from('nW6TyLK5QYvCkItkx57JiGKVkZDaY8WVldgyxMWSjmWSlpeMYcSTkpAAAA==', 'base64');
let buffer2_encoded = Buffer.from('8afwB+0=', 'base64');

let buffer1_decoded = buffer1_encoded.map((value, index) => {

	// function copy(a:int, b:int) {
	//   (...)
	//   var g:int = 4;
	//   var h:int = e[2];
	//   var i:int = 5;
	//   var j:int = h % i;
	//   var k:ubyte_ptr = g - j;
	//   var l:int = k[1067];     // [4 - index % 5]
	//   (...)
	// }

	return value ^ buffer2_encoded[4 - (index % 5)];

});

console.log(buffer1_decoded.toString('utf8'));
