import subprocess
from flask import Flask, request, send_file
import os

app = Flask(__name__)

def fix_path(user_input):
    user_input = str(user_input)
    if ".." in user_input or "./" in user_input or "./" in user_input:
        return False
    return True

@app.route('/')
def index():
    return send_file(os.path.join(os.path.dirname(__file__), "static", "index.html"))

@app.route("/static/<file_name>")
def get_file(file_name):
    fixed_path = fix_path(file_name)
    file_to_send = os.path.join(os.path.dirname(__file__), "static", file_name)
    return send_file(file_to_send)

@app.route('/recommendations', methods=['GET'])
def show_path():
    user_input = request.args.get('path')
    if not fix_path(user_input):
        return "black listed characters detected"
    prefix = os.path.join(os.path.dirname(__file__), "paths")
    file_to_send = os.path.join(prefix,user_input)
    return send_file(file_to_send)

