
# Natas 23 (Over The Wire CTF)

Credentials:

- Username: `natas23`
- Password: `qjA8cOoKFTzJhtV0Fzvt92fgvxVnVRBj`


# Problem

The comparison for the password depends on two conditions:

`$_REQUEST["passwd"] > 10`:

PHPs type casting system can be abused to bypass this, with a string
starting with "number_iloveyou" where `number` must be a value greater
than 10 as a string. For example, `123iloveyou` would also fulfill
this condition.

`strstr($_REQUEST["passwd"], "iloveyou")`:

This condition only checks whether or not the string contains the
string value `iloveyou`.

# Solution

The `passwd` parameter can be set to `11iloveyou` to bypass the
password check.

```bash
cd ./solution;

go run main.go;
```

