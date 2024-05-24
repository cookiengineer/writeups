
# Natas 22 (Over The Wire CTF)

Credentials:

- Username: `natas22`
- Password: `91awVM9oDiUGm33JdzM7RVLBS8bz9n0s`


# Problem

Redirects are not a safe way to fix problems in the following code.

The `GET` parameter `?revelio` leads to the password being printed.

The code also tries to hide the admin password by using a `Location`
HTTP header which leads to a redirect, which is obviously not a secure
way to hide the mistake of printing the password.


# Solution

In the Network Debugger of the Web Browser, there is a field to persist
the network request log (Settings Icon, Persist Logs in Firefox), which
will reveal the password.

The solution, however, just gets the original request without following
the `Location` redirect header.

```bash
cd ./solution;

go run main.go;
```

