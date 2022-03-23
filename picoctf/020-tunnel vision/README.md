
# Tunnel Vision

Corrupted BMP file


# Problem

The BMP file is corrupted, and the challenge is to find the bytes
that are corrupted and to modify them, so that the picture is
visible again.


# Solution

This is a manual process in a Hex Editor. I used `ImHex`, because
it offers `patterns` for specific file formats such as `BMP`, and
this makes it super-easy to identify corrupted values that don't
match the expected patterns and/or static headers.

