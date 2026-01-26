
# behemoth0

- Username: `behemoth0`
- Password: `behemoth0`


# Solution

If you use `ltrace` to find out what methods are being called, you can observe
that the entered password is compared to `eatmyshorts`, which is stored as a
char encoded array in the memory as `OK^GSYBEX^Y`.

```bash
cd /behemoth;
ltrace ./behemoth0; # Enter "whatever"
```

```bash
__libc_start_main(0x804921d, 1, 0xffffd694, 0 <unfinished ...>
printf("Password: ")                                             = 10
__isoc99_scanf(0x804a054, 0xffffd58f, 0x804a008, 0x804a020Password: whatever
)      = 1
strlen("OK^GSYBEX^Y")                                            = 11
strcmp("whatever", "eatmyshorts")                                = 1
puts("Access denied.."Access denied..
)                                          = 16
+++ exited (status 0) +++
```

```bash
./behemoth0; # Enter "eatmyshorts"
behemoth0@gibson:/behemoth$ ./behemoth0
Password: eatmyshorts
Access granted..
$ cat /etc/behemoth_pass/behemoth1
8JHFW9vGru
```
