
enc = open("enc").read()
dec = "";

for e in enc:
    dec += hex(ord(e)).lstrip("0x")

print(bytes.fromhex(dec).decode("utf-8"))

