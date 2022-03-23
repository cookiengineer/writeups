
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
script, but somehow I messed it up and crashed the server :(


1. Decode the base64 encoded cookie
2. Decode the base64 encoded cookie again (yeah, so safe)
3. For the first 128 bits (assuming CBC 128 bit key length) flip each bit, double-encode the cookie again and do a request with the spoofed cookie.


Currently I'm not sure whether the script still crashes the server,
but while the server restarted I reimplemented it via the
`bruteforce.py` script which eventually worked and got me the flag.

