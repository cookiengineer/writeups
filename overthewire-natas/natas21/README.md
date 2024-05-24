
# Natas 21 (Over The Wire CTF)

Credentials:

- Username: `natas21`
- Password: `89OWrTkGmiLZLv12JY4tLj2c4FW0xn56`


# Problem

The session cookies for the `natas21-experimenter` colocated page don't have a correct
domain scope nor a path limitation, which leads to that the session cookies can be
reused on the other subdomain.

# Solution

The (colocated) experimenter page allows to use some debug flags and to override the
admin flag by simply appending the `?debug` parameters to the URL. This will lead to
all parameters in the `POST` body being set into the `$_SESSION` object, which in
return allows us to set any key/value pair.

Then it's possible to reuse the session cookie from the experimenter page to the
colocated CTF page, and the session cookie will lead to the password being revealed.

```bash
cd ./solution;

go run main.go;
```

