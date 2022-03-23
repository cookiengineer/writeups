// var DICTIONARY = ['0a029}', '_again_5', 'this', 'Password\x20Verified', 'Incorrect\x20password', 'getElementById', 'value', 'substring', 'picoCTF{', 'not_this'];
//
// (function(_0x4bd822, _0x2bd6f7) {
//
//     var _0xb4bdb3 = function(_0x1d68f6) {
//         while (--_0x1d68f6) {
//             _0x4bd822['push'](_0x4bd822['shift']());
//         }
//     };
//     _0xb4bdb3(++_0x2bd6f7);
//
// }(DICTIONARY, 0x1b3));

DICTIONARY = [
	'getElementById',      // 0
	'value',               // 1
	'substring',           // 2
	'picoCTF{',            // 3
	'not_this',            // 4
	'0a029}',              // 5
	'_again_5',            // 6
	'this',                // 7
	'Password Verified',   // 8
	'Incorrect password'   // 9
];



console.log(DICTIONARY);



function verify() {

	var checkpass = document.getElementById('pass').value;

	if (checkpass.substr(0, 8) == 'picoCTF{') {
		if (checkpass.substr(8, 16) == 'not_this') {
			if (checkpass.substr(16, 24) == '_again_5') {
				if (checkpass.substr(24, 32) == '0a029}') {
					alert('Password Verified');
				}
			}
		}
	} else {
		alert('Incorrect password');
	}

}
