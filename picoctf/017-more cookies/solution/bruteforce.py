
import requests
import base64


session = requests.Session()
session.get('http://mercury.picoctf.net:21553/')

cookie_encoded = session.cookies['auth_name']
cookie_decoded = base64.b64decode(base64.b64decode(cookie_encoded))

print(cookie_decoded)


for bit in range(0, 128):

    position = bit // 8
    flipped_decoded = cookie_decoded[0:position] + chr(ord(cookie_decoded[position]) ^ (1 << (bit % 8))) + cookie_decoded[position + 1:]
    flipped_encoded = base64.b64encode(base64.b64encode(flipped_decoded))

    response = requests.get('http://mercury.picoctf.net:21553/', cookies = { 'auth_name': flipped_encoded })

    if 'picoCTF' in response.text:
        print('The #' + str(bit) + ' was correct!')
        print(response.text)
    else:
        print('The #' + str(bit) + ' was incorrect...')

