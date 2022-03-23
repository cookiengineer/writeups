
# Static ain't always noise

Binary containing plaintext strings


# Problem

The sample program / binary contains the flag,
and it is up to the user to find it.


# Solution

Simply use `objdump` and `strings` to find the
strings in the `.text` section of the binary.

Either that, or execute the sample script `ltdis.sh`
that does that for you.

```bash
objdump -sj .text -D ./path/to/static;
strings -a -t x ./path/to/static;
```

