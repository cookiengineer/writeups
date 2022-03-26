
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
two statically embedded `base64` strings, where one is the masked string and
the other one is the mask itself.

This time, the mask is computed in a more complicated manner, and the challenge
is to follow through the code and understand what it's doing.

The `WebAssembly` memory contains at offset `1024` the masked string, and at offset
`1067` the mask for decoding it. The `copy` function in the decompiled `.dcmp` file
contains the algorithm that is used to decode the masked string.

I implemented the `decoder.js` file which shows how to do that in `node.js`, and
which returns the decoded flag.

