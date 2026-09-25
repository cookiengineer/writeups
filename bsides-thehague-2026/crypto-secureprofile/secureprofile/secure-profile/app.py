import os
import secrets
import urllib.parse

from flask import Flask, request, redirect, make_response, render_template_string
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad

app = Flask(__name__)

BLOCK_SIZE = 16
KEY = secrets.token_bytes(16)
FLAG = os.getenv("FLAG", "FLAG{dummy_flag_change_me}")


PAGE = """
<!doctype html>
<html>
<head>
    <meta charset="utf-8">
    <title>Secure Profile Factory</title>
    <style>
        body {
            font-family: sans-serif;
            max-width: 850px;
            margin: 40px auto;
            padding: 0 20px;
            line-height: 1.5;
        }
        code, pre, textarea, input {
            font-family: monospace;
        }
        .box {
            border: 1px solid #ccc;
            padding: 18px;
            margin-bottom: 20px;
            border-radius: 10px;
        }
        input[type=text], textarea {
            width: 100%;
            box-sizing: border-box;
            padding: 10px;
        }
        button {
            padding: 10px 16px;
            cursor: pointer;
        }
        .good {
            background: #eaffea;
            border-color: #7bc67b;
        }
        .bad {
            background: #fff4f4;
            border-color: #e19999;
        }
    </style>
</head>
<body>
    <h1>Secure Profile Factory</h1>
    <p>
        Create an encrypted profile token and load it back later.
        Profiles are encrypted so they can't be tampered with. Probably.
    </p>

    <div class="box">
        <h2>Create profile</h2>
        <form method="post" action="/register">
            <label for="name">Display name</label>
            <input type="text" id="name" name="name" maxlength="128"
                   placeholder="alice">
            <p>
                Tip: this field accepts URL-encoded bytes too, so
                <code>%41%42%43</code> becomes <code>ABC</code>.
            </p>
            <button type="submit">Create encrypted profile</button>
        </form>
    </div>

    <div class="box">
        <h2>Current session</h2>
        <p>
            Your encrypted profile is stored in the <code>profile</code> cookie.
            Reload this page with a different cookie to test another profile.
        </p>
    </div>

    {% if error %}
    <div class="box bad">
        <strong>Error:</strong> {{ error }}
    </div>
    {% endif %}

    {% if token %}
    <div class="box">
        <h2>Current token</h2>
        <textarea rows="4" readonly>{{ token }}</textarea>
    </div>
    {% endif %}

    {% if raw_profile %}
    <div class="box">
        <h2>Decrypted profile</h2>
        <pre>{{ raw_profile }}</pre>
    </div>
    {% endif %}

    {% if profile %}
    <div class="box {% if is_admin %}good{% endif %}">
        <h2>Parsed profile</h2>
        <p><strong>username:</strong> {{ profile.get("username", "") }}</p>
        <p><strong>role:</strong> {{ profile.get("role", "") }}</p>

        {% if is_admin %}
            <h3>Admin panel</h3>
            <p>Welcome, administrator.</p>
            <p><strong>Flag:</strong> <code>{{ flag }}</code></p>
        {% else %}
            <p>You are a normal user.</p>
        {% endif %}
    </div>
    {% endif %}
</body>
</html>
"""


def ecb_encrypt(data: bytes) -> bytes:
    cipher = AES.new(KEY, AES.MODE_ECB)
    return cipher.encrypt(pad(data, BLOCK_SIZE))


def ecb_decrypt(data: bytes) -> bytes:
    cipher = AES.new(KEY, AES.MODE_ECB)
    return unpad(cipher.decrypt(data), BLOCK_SIZE)


def profile_for(name_raw: str) -> bytes:
    name = urllib.parse.unquote_to_bytes(name_raw)
    name = name.replace(b"&", b"").replace(b"=", b"")
    return b"username=" + name + b"&role=user"


def encode_token(name_raw: str) -> str:
    profile = profile_for(name_raw)
    return ecb_encrypt(profile).hex()


def decode_token(token_hex: str) -> bytes:
    data = bytes.fromhex(token_hex)
    if len(data) == 0 or len(data) % BLOCK_SIZE != 0:
        raise ValueError("Ciphertext must be a non-empty multiple of 16 bytes.")
    return ecb_decrypt(data)


def parse_profile(raw: bytes) -> dict[str, str]:
    result = {}
    for chunk in raw.split(b"&"):
        if b"=" in chunk:
            k, v = chunk.split(b"=", 1)
            result[k.decode("latin-1", errors="ignore")] = v.decode("latin-1", errors="ignore")
    return result


@app.route("/", methods=["GET"])
def index():
    token = request.cookies.get("profile")
    raw_profile = None
    parsed = None
    is_admin = False
    error = None

    if token:
        try:
            raw = decode_token(token)
            raw_profile = raw.decode("latin-1", errors="replace")
            parsed = parse_profile(raw)
            is_admin = parsed.get("role") == "admin"
        except Exception:
            error = "Stored token is invalid."

    return render_template_string(
        PAGE,
        token=token,
        raw_profile=raw_profile,
        profile=parsed,
        is_admin=is_admin,
        flag=FLAG if is_admin else None,
        error=error,
    )


@app.route("/register", methods=["POST"])
def register():
    name = request.form.get("name", "")
    token = encode_token(name)

    resp = make_response(redirect("/"))
    resp.set_cookie("profile", token)
    return resp


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=3000, debug=False)