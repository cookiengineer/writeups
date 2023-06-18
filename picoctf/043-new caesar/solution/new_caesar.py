import string

LOWERCASE_OFFSET = ord("a")
ALPHABET = string.ascii_lowercase[:16]

def b16_encode(plain):
    enc = ""
    for c in plain:
        binary = "{0:08b}".format(ord(c))
        enc += ALPHABET[int(binary[:4], 2)]
        enc += ALPHABET[int(binary[4:], 2)]
    return enc

def b16_decode(enc):
    plain = ""
    for c in range(0, len(enc), 2):
        tmp = ""
        tmp += "{0:b}".format(ALPHABET.index(enc[c])).zfill(4)
        tmp += "{0:b}".format(ALPHABET.index(enc[c+1])).zfill(4)
        plain += chr(int(tmp, 2))
    return plain

def shift(c, k):
    t1 = ord(c) - LOWERCASE_OFFSET
    t2 = ord(k) - LOWERCASE_OFFSET
    return ALPHABET[(t1 + t2) % len(ALPHABET)]

def unshift(c, k):
    t1 = ord(c) - LOWERCASE_OFFSET
    t2 = ord(k) - LOWERCASE_OFFSET
    return ALPHABET[(t1 - t2) % len(ALPHABET)]

encoded_and_shifted = "ihjghbjgjhfbhbfcfjflfjiifdfgffihfeigidfligigffihfjfhfhfhigfjfffjfeihihfdieieih"

print("encoded and shifted", encoded_and_shifted)

for key in ALPHABET:

    unshifted = ""
    for i,c in enumerate(encoded_and_shifted):
        unshifted += unshift(c, key[i % len(key)])

    decoded = b16_decode(unshifted)
    print("key \"" + key + "\"", decoded)

# ORIGINAL CODE
#key = "p"
#assert all([k in ALPHABET for k in key])
#assert len(key) == 1
# 
# b16 = b16_encode(flag)
# enc = ""
# for i, c in enumerate(b16):
# 	enc += shift(c, key[i % len(key)])
# print(enc)
