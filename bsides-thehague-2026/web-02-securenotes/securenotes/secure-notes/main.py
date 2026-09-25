import os
import mysql.connector
from flask import Flask, redirect, render_template, request, session, url_for

BASE_DIR = os.path.dirname(os.path.abspath(__file__))

MYSQL_HOST = os.environ.get("MYSQL_HOST", "127.0.0.1")
MYSQL_PORT = int(os.environ.get("MYSQL_PORT", "3306"))
MYSQL_USER = os.environ.get("MYSQL_USER", "ctf")
MYSQL_PASSWORD = os.environ.get("MYSQL_PASSWORD", "ctf")
MYSQL_DATABASE = os.environ.get("MYSQL_DATABASE", "ctf")

NOTES_KEY = os.urandom(32).hex()
ADMIN_PASSWORD = os.urandom(32).hex()

FLAG = os.environ.get("FLAG", "flag{dev_flag_change_me}")

app = Flask(__name__)
app.secret_key = os.urandom(24)


def get_db_connection():
    return mysql.connector.connect(
        host=MYSQL_HOST,
        port=MYSQL_PORT,
        user=MYSQL_USER,
        password=MYSQL_PASSWORD,
        database=MYSQL_DATABASE
    )


def block_malicious(value: str) -> bool:
    blacklist = ["'", '"']
    return any(part in value.lower() for part in blacklist)


def ensure_db() -> None:
    conn = get_db_connection()
    try:
        cur = conn.cursor()
        cur.execute(
            """
            CREATE TABLE IF NOT EXISTS users (
                id INT AUTO_INCREMENT PRIMARY KEY,
                username VARCHAR(255) UNIQUE NOT NULL,
                password VARCHAR(255) NOT NULL
            )
            """
        )
        cur.execute(
            """
            CREATE TABLE IF NOT EXISTS notes (
                id INT AUTO_INCREMENT PRIMARY KEY,
                owner VARCHAR(255) NOT NULL,
                title VARCHAR(255) NOT NULL,
                tag VARCHAR(255) NOT NULL,
                body BLOB NOT NULL
            )
            """
        )
        cur.execute("DROP FUNCTION IF EXISTS KDF")
        cur.execute("""
DELIMITER //

CREATE FUNCTION KDF(p_passphrase VARCHAR(255), p_salt VARCHAR(255), p_iterations INT)
RETURNS CHAR(64)
DETERMINISTIC
BEGIN
  DECLARE v_result VARCHAR(255);
  DECLARE v_i INT DEFAULT 0;

  -- ensure we always do at least 1 round and avoid NULL
  SET p_iterations = GREATEST(1, COALESCE(p_iterations, 1));

  SET v_result = COALESCE(p_passphrase, '');

  WHILE v_i < p_iterations DO
    SET v_result = SHA2(CONCAT(v_result, COALESCE(p_salt, '')), 256);
    SET v_i = v_i + 1;
  END WHILE;

  RETURN v_result; -- 64 hex chars
END//

DELIMITER ;
        """)

        cur.execute("SELECT 1 FROM users WHERE username=%s", ("admin",))
        if not cur.fetchone():
            cur.execute(
                "INSERT INTO users (username, password) VALUES (%s, SHA2(%s, 256))",
                ("admin", ADMIN_PASSWORD),
            )
            conn.commit()

        cur.execute("SELECT 1 FROM notes WHERE owner = 'admin' AND title = 'flag'")
        if not cur.fetchone():
            save_note("admin", "flag", FLAG, "secret")
    finally:
        conn.close()

def save_note(owner: str, title: str, body: str, tag: str) -> None:
    conn = get_db_connection()
    try:
        cur = conn.cursor()
        cur.execute("INSERT INTO notes (owner, title, body, tag) VALUES (%s, %s, AES_ENCRYPT(%s, KDF(%s, 'salty', 3000000)), %s)",
                        (owner, title, body, NOTES_KEY, tag))
        conn.commit()
    finally:
        conn.close()

def read_note(note_id: int, user: str) -> tuple[str, str, str] | None:
    conn = get_db_connection()
    try:
        cur = conn.cursor()
        cur.execute("SELECT title, tag, AES_DECRYPT(body, KDF(%s, 'salty', 3000000)) FROM notes WHERE id = %s AND owner = %s",
                        (NOTES_KEY, note_id, user))
        row = cur.fetchone()
        if row and row[0]:
            return row
    finally:
        conn.close()
    return None

@app.get("/")
def index():
    if session.get("user"):
        return redirect(url_for("home"))
    return render_template("login.html")


@app.post("/login")
def login():
    username = request.form.get("username", "")
    password = request.form.get("password", "")

    conn = get_db_connection()
    try:
        with conn.cursor() as cur:
            cur.execute("SELECT username FROM users WHERE username=%s AND password=SHA2(%s, 256)", (username, password))
            row = cur.fetchone()
    finally:
        conn.close()

    if not row:
        return render_template("login.html", error="Invalid credentials.", username=username)

    session["user"] = row[0]
    return redirect(url_for("home"))

@app.post("/register")
def register():
    username = request.form.get("username", "")
    password = request.form.get("password", "")

    if block_malicious(username) or block_malicious(password):
        return render_template("login.html", error="No quotes allowed!", username=username)

    conn = get_db_connection()
    try:
        with conn.cursor() as cur:
            cur.execute("INSERT INTO users (username, password) VALUES (%s, SHA2(%s, 256))", (username, password))
            conn.commit()
    except Exception as e:
        return render_template("login.html", error="Username already taken.", username=username)
    finally:
        conn.close()

    session["user"] = username
    return redirect(url_for("home"))


@app.get("/register")
def register_get():
    return render_template("register.html")


@app.get("/home")
def home():
    if not session.get("user"):
        return redirect(url_for("index"))
    return render_template("home.html", user=session["user"])


@app.get("/create")
def create_get():
    if not session.get("user"):
        return redirect(url_for("index"))
    return render_template("create.html")


@app.post("/create")
def create_post():
    if not session.get("user"):
        return redirect(url_for("index"))

    title = request.form.get("title", "")
    body = request.form.get("body", "")
    tag = request.form.get("tag", "general")
    owner = session["user"]
    save_note(owner, title, body, tag)
    conn = get_db_connection()
    try:
        cur = conn.cursor()
        cur.execute("SELECT id FROM notes WHERE owner = %s AND title = %s AND tag = %s ORDER BY id DESC LIMIT 1", (owner, title, tag))
        row = cur.fetchone()
        if row:
            note_id = row[0]
            return redirect(url_for("view_note", note_id=note_id))
    finally:
        conn.close()
    return redirect(url_for("home"))


@app.get("/search")
def search_get():
    if not session.get("user"):
        return redirect(url_for("index"))
    return render_template("search.html", rows=None)


@app.post("/search")
def search_post():
    if not session.get("user"):
        return redirect(url_for("index"))

    title = request.form.get("title", "")
    tag = request.form.get("tag", "")

    if block_malicious(title) or block_malicious(tag):
        return render_template("search.html", error="No quotes allowed!", rows=None)

    conn = get_db_connection()
    rows = []
    try:
        cur = conn.cursor()
        query = f"SELECT id, owner, title FROM notes WHERE owner = %s AND (title = '{title}' OR tag = '{tag}');"
        cur.execute(query.split(";")[0], (session["user"],))
        rows = cur.fetchall()
    except Exception as e:
        return render_template("search.html", error=f"MySQL error: {e}", rows=None)
    finally:
        conn.close()
    return render_template("search.html", rows=rows)


@app.post("/logout")
def logout():
    session.clear()
    return redirect(url_for("index"))


@app.get("/note/<int:note_id>")
def view_note(note_id):
    if not session.get("user"):
        return redirect(url_for("index"))

    content = read_note(note_id, session["user"])
    if not content:
        return "Note not found or decryption failed.", 404

    return render_template("note.html", note=content)

if __name__ == "__main__":
    ensure_db()
    app.run(host="0.0.0.0", port=3000, debug=False, threaded=True)