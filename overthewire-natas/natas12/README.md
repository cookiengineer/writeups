
# Natas 12 (Over The Wire CTF)

Credentials:

- Username: `natas12`
- Password: `YWqo0pjpcXzSIl5NMAVxg12QxeC1w9QG`


# Problem

The website's PHP code is written in a way so that it reuses
the client-side `filename` POST parameter as the designated
file name's extension.


# Solution

Craft an `image.php` file which executes code on the server-side,
and modify the `<input type=hidden name=filename value=randomstuff.php>`
so that the `php` file extension is used when the file is written
in the `upload/` folder.
