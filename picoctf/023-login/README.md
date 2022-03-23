
# Login

base64 is not an encryption


# Problem

The code does client-side validation of the password, and tries to obfuscate
it a little by using `base64` as an encoding algorithm.


# Solution

This is a manual process, simply copy the encoded value and decode it via
the `base64` command.

```bash
# username is admin
echo "YWRtaW4=" | base64 -d;

# password/flag is picoCTF{53rv3r_53rv3r_53rv3r_53rv3r_53rv3r}
echo "cGljb0NURns1M3J2M3JfNTNydjNyXzUzcnYzcl81M3J2M3JfNTNydjNyfQ==" | base64 -d;
```

