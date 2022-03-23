
# Cookies

Incremented Session IDs are not secure


# Problem

The server-side implementation uses guessable Session IDs,
which means that the `Cookie: <Number>` header represents
logged in users.


# Solution

As the Cookies are easily spoofable, it's also easy to
bruteforce them to get access. In my `bruteforce.mjs`
implementation I just looked at the first 64 sessions
to find a flag and got it.

