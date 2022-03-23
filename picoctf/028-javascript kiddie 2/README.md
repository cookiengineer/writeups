
# JavaScript Kiddie 2

Byte-Shift-Based Encryption


# Problem

As in `JavaScript Kiddie 1`, the `bytes` file contains all shifted
bytes that represent a PNG image. Again, the PNG header is fixed
and the first 16 bytes are static, and contain the PNG header followed
by the `IDHR` header.


As the `bytes` file this time contains multiple possible offsets
that could lead to potential solutions, the challenge here is more
complex in the sense that there are multiple permutations of those
offsets possible.


# Solution

Instead of manually typing in all 45 possible permutations, I decided
to write a `cracker.mjs` that does this automatically by using
`zbarimg` to detect whether or not the generated PNG images contain
a valid QR code.

1. Convert the bytes into a Buffer.
2. Split the buffer into 16 columns (possible "arm bandit rows").
3. For each column, look for possible matching indexes that match the expected value of the static PNG header.
4. For each possible index, generate permutations of input keys.
5. Use all possible permutations in a loop and generate a decrypted PNG image.
6. Use `zbarimg` to find a potential QR code embedded in the PNG image.

