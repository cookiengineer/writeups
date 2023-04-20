
# Cowsay As a Service

Never use unsanitized strings as a command line argument


# Problem

The `cowsay ${message}` command is invoked on the server,
and the contents of that shell command are returned as a
response.


# Solution

Our task is to invoke the `ls` command to get an overview
of the folder structure, and then to use the `cat` command
in order to download the necessary files.

The `scrape.mjs` implements a simple approach that automatically
downloads all files it finds on the server.


# Further Notes

There's probably a way to invoke shellshock, too.

