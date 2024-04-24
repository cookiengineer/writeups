
# Natas 18 (Over The Wire CTF)

Credentials:

- Username: `natas18`
- Password: `8NEDUUxg8kFgPV84uLwvZkGn6okJQ6aq`


# Problem

The `PHPSESSID` identifier is weak and can be bruteforced, the identifiers
are incrementing and there's a hardcap limit of `640` which means that it's
very feasible to hijack other sessions.


# Solution

The weak session identifier allows us to hijack and iterate over the session ids,
and setting the `Cookie` HTTP header to gain access to another session that's
already running and actively used by another user.

In our case, the session `119` was logged in with the `admin` user. The
implementation detects that session successfully.

```bash
cd ./solution;

go run main.go;
```

