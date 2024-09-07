import random

def encrypt(file_path, output_path):
  with open(file_path, 'rb') as file:
    m = file.read()
  k = bytearray(random.getrandbits(8) for _ in range(4))
  s = 4
  c = bytearray()
  for i in range(0, len(m), s):
    b = m[i:i+s]
    c.extend(bytearray(b[i] ^ k[i % 4] for i in range(len(b))))

  with open(output_path, 'wb') as file:
    file.write(c)

if __name__ == "__main__":
  input_file = 'geheim.zip'
  encrypted_file = 'geheim.zip.crypt'
  encrypt(input_file, encrypted_file)