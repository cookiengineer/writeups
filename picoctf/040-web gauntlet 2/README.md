
# Web Gauntlet 2

SQL injection.


# Problem

The challenge is to bypass the filters in order to get a successful login via the SQL queries.
Each time you're successful, the SQL query is printed out.

The password for each round doesn't matter, you can always enter `whatever`.

# Solution

## Round 1:

Username: `adm'||'in`
Password: `a' IS NOT 'b`

