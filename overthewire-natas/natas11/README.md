
# Natas 11 (Over The Wire CTF)

Credentials:

- Username: `natas11`
- Password: `1KFqoJXi6hRaPluAmk8ESDW4fSysRoIg`


# Problem

The Cookie is XOR encrypted, which is pretty unsecure as a cipher
because it allows to find out the used repeating key.


# Solution

Two helper scripts were needed to solve this.

The first script [01_find_key.php](./solution/01_find_key.php) finds
the used original key by applying the XOR cipher on the decoded data.

In the generated output we can see that `KNHL` is a repeating pattern,
so it is most likely the used key to encrypt the cookie.

The second script [02_generate_cookie.php](./solution/02_generate_cookie.php)
generates the cookie with the necessary flags being set to pass the
challenge.

This needs to be set in the Browser (Network Tab or via `document.cookie`)
to pass the challenge. It can also be used via `curl`.

