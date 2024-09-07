
# BND's Takur Medicine Ransomware Challenge

## First Problem: Broken XOR encryption/rotation with fixed key size

XOR encryption/rotation of the `zip` file uses a fixed `4 bytes length` key which
is enough to infer the possible keys from the `file headers` that a zip file usually
should have. In this case `PK\0x03\0x04` was the initial file header.

```bash
cd ./solution;

# Generates 3 possible zip files
go run ./cmds/decrypt/main.go ../challenge/geheim.zip.crypt;

# The first one is the correct one
mv ./geheim-0.zip ../challenge/geheim.zip;
```

## Second Problem: ZIP encryption password is bruteforceable

The password of the ZIP file needs to be bruteforced, and it was assumed to have a weak
character set (alphanumeric, whitespace, underscore) and standard decryption due to
AES128/AES256 probably taking too long for a challenge like this.

(The challenge duration is less than 2 weeks, and I guess they don't assume everybody
has an OpenCL compatible GPU(s) available).

**Results**:

The password of decrypted zip file (the original `geheim.zip`) is `roneisha2209` which
is part of the `rockyou.txt` file. It therefore can be either reproduced with my
self-written password cracker, or johntheripper.

```bash
cd ./solution;

# Option 1: Crack the password with my own tool
wget https://github.com/brannondorsey/naive-hashcat/releases/download/data/rockyou.txt -O ./cmds/crack-dictionary/passwords.txt;
go run ./cmds/crack-dictionary/main.go ../geheim.zip;

# Option 2: Crack the password with johntheripper
zip2john geheim.zip > zip_hashes.txt;
john --wordlist=/path/to/rockyou.txt zip_hashes.txt;

# If you miss johntheripper being done, show password via:
john --show zip_hashes.txt;
```

## Third Problem: KeePass Master Password findable in memory dump

KeePass had a [CVE-2023-32784](https://nvd.nist.gov/vuln/detail/CVE-2023-32784) earlier this year,
which allowed to find the offsets in a memory dump to find KeePass'es password lengths due to
KeePass using a special character to display the UI/hide the passwords.

The character that is looked for is `0xCF` followed by a `0x25` byte value, and characters for
KeePass passwords have a limited charset from `0x20` to `0xFF` because they're ASCII only.

```bash
cd ./solution;

# Find potential master keys in the memory dump
go run ./cmds/find-keypass-masterkey.go ../challenge/KeePass.DMP;
```


## Solution

Hashes of the cracked/bruteforced/recovered passwords, all generated via `sha512sum` to prove
that they were recovered and hashed again:

Decrypted ZIP file password: `1b52f373a7324b8c344928b997874a9fb5a174cebc841b53a42ce2a0f872cb0d10066a1cc1ec3566ef030ba654af5db25d6341d95404a8615ff9529f1fcaa604`
Unpacked KeePass file password: `b99b89217dad01a5992247a220f5226716dca14c813d8516095cd23240f2ce345fd22883ed4358c441e643e8e98fa09e37ebaec6285ef22e569b6633f263e7a1`
Hash of Final password to submit: `8695447c00858cdb3c0614727963135c0d28abbc9a0c938885dcbd569582b921b6ba1bc876714153ae116d0b84085593297456b5269f465fc282d749d619a64f`

