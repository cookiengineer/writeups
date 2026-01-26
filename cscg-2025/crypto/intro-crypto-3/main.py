import os
import struct

BITS = 56

FLAG = os.getenv("FLAG", "CSCG{TESTFLAG}")

A = int.from_bytes(os.urandom(BITS//8), "little")
B = int.from_bytes(os.urandom(BITS//8), "little")
SEED = int.from_bytes(os.urandom(BITS//8), "little")

def rng(x, size):
    return (x*A+B) & ((2**size)-1)

def gen_random(seed, bits, mask):
    state = seed
    while True:
        state = rng(state, bits)
        yield state & mask

def main():
    print("Here are some random numbers, now guess the flag")
    rng = gen_random(SEED, BITS, 0xFF)
    for i in range(len(FLAG)):
        print(next(rng) ^ ord(FLAG[i]))

if __name__ == "__main__":
    main()