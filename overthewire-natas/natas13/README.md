
# Natas 13 (Over The Wire CTF)

Credentials:

- Username: `natas13`
- Password: `lW3jYRI02ZKDBb8VtQBU1f6eDRo6WEj9`


# Problem

The website's PHP code is written in a way so that it reuses
the client-side `filename` POST parameter as the designated
file name's extension.

Additional security measurement is to check against EXIF file
headers. The problem is though, that PHP files can be executed
no matter what kind of text (or bytes) come before the `<?php`
header.


# Solution

Craft an `image.php` file which executes code on the server-side,
and modify the `<input type=hidden name=filename value=randomstuff.php>`
so that the `php` file extension is used when the file is written
in the `upload/` folder.

Use a valid EXIF header before the `<?php` as the first bytes of the file.
The JPEG header `0xFF 0xD8 0xFF 0xEE` worked for me.

