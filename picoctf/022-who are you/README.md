
# Who Are You?


# Problem

The code on the server-side uses a series of HTTP headers and clues
to verify whether or not the client is allowed to access the website.


# Solution

Simply use `curl -H 'Header: Value'` to override the HTTP headers.

As all HTTP headers that are used in this Challenge are spec-compliant
with RFC2616, there's not that many options to choose from :)

