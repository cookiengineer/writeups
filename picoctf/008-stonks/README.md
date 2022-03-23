
# Stonks

C Format String vulnerability


# Problem

The server's sample program `vuln.c` uses `printf()` which
is vulnerable to overflows.


# Solution

Entering `%X` loooooots of times will get us higher in the
memory stack, allowing us to read the memory that we're not
supposed to be able to read.

