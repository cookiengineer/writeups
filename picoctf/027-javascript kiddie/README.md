
# JavaScript Kiddie

Byte-Shift-Based Encryption


# Problem

The `bytes` file contains all shifted bytes that represent a PNG image.
The PNG header is fixed and the first 16 bytes are static, and contain
the PNG header followed by the `IDHR` header.


As the `bytes` file contains only a single-digit offset for each column,
there's only one valid solution (which is automatically the first index
by coincidence).


# Solution

Instead of analyzing the file manually in a hex editor, I decided to
automate this process and write a simple `convert.mjs` script.

1. Convert the bytes into a Buffer.
2. Split the buffer into 16 columns (possible "arm bandit rows").
3. For each column, look for the index that matches the expected value of the static PNG header.
4. Generate the key and decrypt the PNG image.
5. Use `zbarimg` to find the QR code embedded in the PNG image.

