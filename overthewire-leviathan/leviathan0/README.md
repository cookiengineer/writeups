
# Leviathan 0 (Over The Wire CTF)

Credentials:

- Username: `leviathan0`
- Password: `leviathan0`


# Problem

The `~/.backup/bookmarks.html` contains a bookmark for the `leviathan1` user.


# Solution

```bash
ssh leviathan0@leviathan.labs.overthewire.org -p 2223;

cat .backup/bookmarks.html | grep leviathan;

<DT><A HREF="http://leviathan.labs.overthewire.org/passwordus.html | This will be fixed later, the password for leviathan1 is 3QJ3TgzHDq" ADD_DATE="1155384634" LAST_CHARSET="ISO-8859-1" ID="rdf:#$2wIU71">password to leviathan1</A>
```
