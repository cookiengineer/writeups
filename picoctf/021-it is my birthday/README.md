
# It Is My Birthday

MD5 is not a collision-safe hash algorithm


# Problem

On the server-side, the code expects two different `PDF` which
are expected to be two files with the same `MD5` hash.


# Solution

As the code on the server-side is too stupid to validate the
metadata of a PDF file, it's possible to use the demo collisions
from the paper about MD5 collisions and renaming the `.exe` files
into `.pdf` to get the flag.

Another method would've been to use the `evilize` tool from the
paper to generate two separate PDF files with a colliding hash.

