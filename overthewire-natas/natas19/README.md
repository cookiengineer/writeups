
# Natas 19 (Over The Wire CTF)

Credentials:

- Username: `natas19`
- Password: `8LMJEhKFbMKIL2mxQKjv0aEDdk7zpT0s`


# Problem

The `PHPSESSID` identifier is weak and can be bruteforced, the identifiers
are incrementing and there's a hardcap limit of `640` which means that it's
very feasible to hijack other sessions.

The previous CTF [Natas 18](../natas18) used only the number as the session
identifier, but this one uses a hexadecimal encoded variant of the format
`<number>-<user>` which is equivalently unsafe as an identifier format.


# Solution

The weak session identifier allows us to hijack and iterate over the session ids,
and setting the `Cookie` HTTP header to gain access to another session that's
already running and actively used by another user.

In our case, the session `bin2hex("281-admin")` was logged in with the `admin` user.
The implementation detects that session successfully.

```bash
cd ./solution;

go run main.go;
```

