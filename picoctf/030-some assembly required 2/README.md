
# Some Assembly Required 2

WebAssembly Reverse Engineering


# Problem

On the client-side, there's a code running that was compiled
and obfuscated into WebAssembly. The `index.html` uses this
code and calls it. However, WebAssembly is hard to read and
the challenge is to reverse engineer the `copy_char` and `strcmp`
methods in the WebAssembly file.


# Solution

- Install `binaryen` on your machine.
- Use `wasm2js` to decompile the WebAssembly file into `JavaScript` again.

The WebAssembly code exports three methods (`$1`, `$2` and `$3`), and includes
a statically embedded `base64` string, which is masked. The demasking function
is `$2`, which uses a `byte ^ 8` mask. Luckily, `XOR`s inverse function is `XOR`.

I implemented the `decoder.js` file which shows how to do that in `node.js`, and
which returns the decoded flag.

