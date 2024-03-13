
# Natas 15 (Over The Wire CTF)

Credentials:

- Username: `natas15`
- Password: `TTkaI7AWG4iDERztBcEyKV7kRXH1EZRB`


# Problem

MySQL Injection for `username` parameter.

The website's PHP code for checking whether or not a user exists
gives us a hint for when the executed MySQL expression returns a
true result.


# Solution

This gives us the possibility to write a Bruteforce script which
uses the `username` set to `natas16" AND password LIKE "<prefix>%`
where the `<prefix>` parameter will be set incrementally by the
tool to find out whether or not the password starts with the given
prefix.

The solution is implemented in `Go`: [solution](./solution/main.go)

```bash
cd ./solution;

go run main.go;
```
