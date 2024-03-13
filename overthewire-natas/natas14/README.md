
# Natas 14 (Over The Wire CTF)

Credentials:

- Username: `natas14`
- Password: `qPazSJBmrmU7UQJv17MHk1PGC4DxZMEP`


# Problem

MySQL Injection for both `username` and `password` parameter.

The website's PHP code uses the `mysqli_num_rows()` method
which expects a boolean result, which makes the crafting aspect
of the injected SQL statement a little harder.

# Solution

- Use the `?debug` parameter first
- Second parameter is `&username=1" or "1"`
- Third parameter is `&password="1" or "1" = "1`

