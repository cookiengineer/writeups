
# PicoBrowser


# Problem

The code on the server-side uses HTTP headers to verify whether
or not the `User-Agent` is set to `picobrowser`.


# Solution

Simply use `curl -H 'Header: Value'` to override the HTTP headers.

