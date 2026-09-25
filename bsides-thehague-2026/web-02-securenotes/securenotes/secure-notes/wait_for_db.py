import mysql.connector
import time
import os


def wait_for_db(host: str, port: int, user: str, password: str, database: str) -> None:
    while True:
        try:
            conn = mysql.connector.connect(
                host=host,
                port=port,
                user=user,
                password=password,
                database=database
            )
            conn.close()
            print("Database is ready!")
            break
        except Exception as e:
            print(f"Waiting for database... ({e})")
            time.sleep(2)


def grant_ctf_all_privileges(
    host: str,
    port: int,
    root_user: str,
    root_password: str,
    ctf_user: str,
) -> None:
    conn = mysql.connector.connect(
        host=host,
        port=port,
        user=root_user,
        password=root_password,
    )
    cursor = conn.cursor()
    cursor.execute(f"GRANT ALL PRIVILEGES ON *.* TO '{ctf_user}';")
    cursor.execute("FLUSH PRIVILEGES;")
    conn.commit()
    cursor.close()
    conn.close()
    print(f"Granted all privileges to {ctf_user}")


if __name__ == "__main__":
    MYSQL_HOST = os.environ.get("MYSQL_HOST", "mysql")
    MYSQL_PORT = int(os.environ.get("MYSQL_PORT", "3306"))
    MYSQL_USER = os.environ.get("MYSQL_USER", "ctf")
    MYSQL_PASSWORD = os.environ.get("MYSQL_PASSWORD", "ctf")
    MYSQL_DATABASE = os.environ.get("MYSQL_DATABASE", "ctf")
    MYSQL_ROOT_USER = os.environ.get("MYSQL_ROOT_USER", "root")
    MYSQL_ROOT_PASSWORD = os.environ.get("MYSQL_ROOT_PASSWORD", "")

    wait_for_db(MYSQL_HOST, MYSQL_PORT, MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE)

    if MYSQL_ROOT_PASSWORD:
        grant_ctf_all_privileges(
            MYSQL_HOST,
            MYSQL_PORT,
            MYSQL_ROOT_USER,
            MYSQL_ROOT_PASSWORD,
            MYSQL_USER,
        )
    else:
        print("Skipping GRANT: MYSQL_ROOT_PASSWORD is not set")
    conn = mysql.connector.connect(
        host=MYSQL_HOST,
        port=MYSQL_PORT,
        user=MYSQL_USER,
        password=MYSQL_PASSWORD,
        database=MYSQL_DATABASE
    )
    cursor = conn.cursor()
    cursor.execute("DROP TABLE IF EXISTS users;")
    cursor.execute("DROP TABLE IF EXISTS notes;")
    conn.commit()
    cursor.close()
    conn.close()
    print("Cleared users and notes tables")
