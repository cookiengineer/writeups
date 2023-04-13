
# More Cookies

AES CBC Mode Encryption is not secure


# Problem

The server-side implementation uses AES encryption in CBC Mode,
which means that it is vulnerable to a Bit Flipping Attack.

In AES, the cypher is cyclic, which also implies that the first
128 bits in this case will lead to flipping the bits in the
decrypted plaintext.


# Solution

I first tried to implement it in node.js via the `bruteforce.mjs`
script, but somehow I messed it up and crashed the server after
a bunch of requests. So be careful not to run it too often :(


1. Decode the base64 encoded cookie
2. Decode the base64 encoded cookie again (yeah, so safe)
3. For the first 128 bits (assuming CBC 128 bit key length) flip each bit.
4. Encode the spoofed cookie via base64.
5. Encode the spoofed cookie again via base64.
6. Send the spoofed cookie to the server, which will behave differently (with a redirect) in case of success.

