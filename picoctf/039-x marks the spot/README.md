
# X Marks The Spot

XML XPath Injection and vulnerabilities.


# Problem

The challenge is to bypass and abuse the XSLT processor on the backend in order to
gather intelligence about the users, the password lengths, and then about the passwords
themselves.

Additionally the login form and responses are so crappy that the "not true" conditions
are a little hacky, so you have to write conditions that allow to reason in case the
login failed (e.g. using `>` and `<` instead of `>=` or `<=` for comparison).


## Data Structure

Due to how the API is being called, and by what was found out during tryouts of multiple
XPath queries, we assume that the XML data structure on the server-side looks something
like this:

```xml
<users>
    <user id="1"> <!-- example user id -->
        <name>... username</name>
        <pass>... password</pass>
    </user>
</users>
```

This implies that we can craft multiple XPath queries that abuse the data structures,
and use things like `string-length()`, `text()` and others for the queries; because
they are being processed by the XPath query selector on the server-side.

In order to do so, it is generally recommended to use the first text input field for
this; in order to have a better validation. As generally most database queries are
selecting the username first; and then matching it against the password, it's more
likely to succeed when using the injection strings in the username text field.


# Solution

## Round 1: Find out usernames (user ids) with short passwords

We can craft an XPath injection that returns a success/failure response when we use a
query that compares the password length. If the password length is longer than the query
compares against, we get a response of failure.

Initially I tried out the length of `30` for the sake of tryout, and in the resulting
`users.json` there were two user ids missing, which implied that they had a longer password.
When trying it out with a maximum length of `64`, all user ids appeared and they were
appearing in the resulting `users.json` database.

```bash
node bruteforce-password-lengths.mjs;

# users with password length==0 are likely to not exist
vim password-lengths.json;
```


## Round 2: Bruteforce usernames

In Round 1 we found out that there are 3 possibly valid user ids, each with their own
password length represented in the `users.json`.

Now it's time to found out the usernames. We craft an XPath injection that compares the
text value of the `<name>` element in the XML with the bruteforced username(s).

```bash
node bruteforce-usernames.mjs;
vim usernames.json;
```


## Round 3: Find the Flag

One of the users must be the flag. By obversing the results in Round 2, we can assume
that the user with the longest password (`length` being `50` characters) is likely to
be the flag.

We can verify this by crafting an XPath query that compares the `<pass>` element's `text()`
with the prefix `picoCTF{`.

```bash
node find-flag.mjs;
# user with id=3 has a password that starts with picoCTF{...
```


## Round 4: Bruteforce passwords

The job would now be to bruteforce only the password of the user with id `3`. But because
we know what to do, we can just generally start to bruteforce the passwords for all users.

```bash
node bruteforce-passwords.mjs;
vim passwords.json;
```
