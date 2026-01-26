
# Leviathan 1 (Over The Wire CTF)

Credentials:

- Username: `leviathan1`
- Password: `3QJ3TgzHDq`


# Problem

The `check` binary checks against the password, so we need to download and analyze it.

```bash
cd ./challenge;

sshpass -p "3QJ3TgzHDq" scp -P 2223 leviathan1@leviathan.labs.overthewire.org:/home/leviathan1/check ./
```

# Solution

Using `ltrace ./check` we can see that the program compares the entered value `123` against `sex`,
which is the password.

```bash
leviathan1@gibson:~$ ltrace ./check
__libc_start_main(0x80490ed, 1, 0xffffd494, 0 <unfinished ...>
printf("password: ")                                             = 10
getchar(0, 0, 0x786573, 0x646f67password: 123456
)                                = 49
getchar(0, 49, 0x786573, 0x646f67)                               = 50
getchar(0, 0x3231, 0x786573, 0x646f67)                           = 51
strcmp("123", "sex")                                             = -1
puts("Wrong password, Good Bye ..."Wrong password, Good Bye ...
)                             = 29
+++ exited (status 0) +++
```

Executing the binary again with the password `sex`, we are now the user `leviathan2` and can
access the flag located at `/etc/leviathan_pass/leviathan2`.

```bash
leviathan1@gibson:~$ ./check 
password: sex

$ whoami
leviathan2

$ cat /etc/leviathan_pass/leviathan2
NsN1HwFoyN
```
