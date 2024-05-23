
# Natas 25 (Over The Wire CTF)

Credentials:

- Username: `natas25`
- Password: `O9QD9DZBDq1YpswiTM5oqMDaOtuZtAcx`


# Problem

The error logging method which is used to log path traversal attempts for the `natas_webpass`
file paths uses the unsanitized `User-Agent` HTTP header and appends it into the log file.

Combined with the path traversal detection mechanism only using a string replacement once
(and not until no match is found) means that this very log file can be included and executed
via PHP.

The path traversal method can be bypassed by prepending `..` and appending a `/` like so:

- `....//` will lead to a `../`
- `..//` will lead to a `./`

The PHP session identifier (`PHPSESSID`) is also known to the end-user because it is set via
the `Set-Cookie` header, which makes the file names in `/var/www/natas/natas25/logs` predictable.


# Solution

It is necessary to do 3 HTTP requests:

1. Get the index file and parse the `Cookie` header to get the `PHPSESSID` identifier.
2. Get the `?lang=natas_webpass` file to trigger the `logRequest()` method, and use `User-Agent: <?php (...) ?>` and `Cookie: (...)`.
3. Get the `?lang=....//....//....//....//....//var/www/natas/natas25/logs/natas25_(...).log`.

```bash
cd ./solution;

go run main.go;
```

