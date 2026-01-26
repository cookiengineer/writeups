
# behemoth1

- Username: `behemoth1`
- Password: `8JHFW9vGru`


# Solution

The binary is prone to a buffer overflow, because it uses `gets()` to fill the buffer without a limit.
Inside the binary, `0x44` bytes have been allocated, which means we need at least `68 bytes` of input
before the buffer overflows. Afterwards, we need to get at the `EIP` which is at offset (TODO).


```bash
cd /behemoth;

behemoth1@gibson:/behemoth$ gdb -q ./behemoth1
Reading symbols from ./behemoth1...
(No debugging symbols found in ./behemoth1)
(gdb) set disassembly-flavor intel
(gdb) disas main
Dump of assembler code for function main:
   0x08049196 <+0>:	push   ebp
   0x08049197 <+1>:	mov    ebp,esp
   0x08049199 <+3>:	sub    esp,0x44 # 68 Bytes allocated for buffer
   0x0804919c <+6>:	push   0x804a008
   0x080491a1 <+11>:	call   0x8049050 <printf@plt>
   0x080491a6 <+16>:	add    esp,0x4
   0x080491a9 <+19>:	lea    eax,[ebp-0x43]
   0x080491ac <+22>:	push   eax
   0x080491ad <+23>:	call   0x8049060 <gets@plt>
   0x080491b2 <+28>:	add    esp,0x4
   0x080491b5 <+31>:	push   0x804a014
   0x080491ba <+36>:	call   0x8049070 <puts@plt>
   0x080491bf <+41>:	add    esp,0x4
   0x080491c2 <+44>:	mov    eax,0x0
   0x080491c7 <+49>:	leave  
   0x080491c8 <+50>:	ret                               # return EIP
End of assembler dump.
```

