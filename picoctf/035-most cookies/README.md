
# Most Cookies

Flask App Secrets should be completely randomized


# Problem

The server-side implementation uses a dictionary to choose
the Flask App Secret upon running the `server.py` file.

So this is basically a dictionary attack, as the amount of
encryption keys is not really random; and the salt for the
generation of the hashes is also static and `cookie-session`
by default.


# Solution

1. Open up the Browser to the CTF index page.
2. Copy the `session` Cookie from the DevTools (Network Tab/Response).
3. Run the `python crack.py <session_cookie>` command, which will do the dictionary attack and then send the spoofed cookie to the `/display` route to get the flag.

