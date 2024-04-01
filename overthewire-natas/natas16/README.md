
# Natas 16 (Over The Wire CTF)

Credentials:

- Username: `natas16`
- Password: `TRD7iZrd5gATjj9PkPEuaOlfEjHqj32V`


# Problem

Shell Injection for `username` parameter.

The website's PHP code executes `grep -i "$username" dictionary.txt`
which leads to potential shell escape techniques, or embedded shell
execution by executing e.g. `$(echo whatever)zigzags`.


# Solution

This gives us the possibility to write a Bruteforce script which
uses the `username` set to `$(grep ^<prefix> /etc/natas_webpass/natas17)zigzags`
which can be used to see whether or not the password starts with the
given prefix.

- If the output is nothing (`<pre>\n</pre>`), it means that the password matched.
- If the output is `zigzags`, it means that the password didn't match.

The solution is implemented in `Go`: [solution](./solution/main.go)

```bash
cd ./solution;

go run main.go;
```
