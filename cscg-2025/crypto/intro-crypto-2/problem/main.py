#!/usr/bin/env python3

from hashlib import sha1
from base64 import b64encode, b64decode
from secrets import token_hex

from secret import FLAG


KEY = token_hex(16)


def get_mac(data: bytes) -> str:
    return sha1(KEY.encode("latin1") + data).hexdigest()


def parse_token(token: str) -> dict:
    # Decode token
    token = b64decode(token)

    # Check the MAC
    token, mac = token.split(b"|mac=")
    if get_mac(token) != mac.decode("latin1"):
        return None

    # Parse values
    values = dict()
    for part in token.decode("latin1").split("|"):
        key, value = part.split("=")
        values[key] = value
    return values


def generate_token(values: dict) -> str:
    token = "|".join(f"{key}={value}" for key, value in values.items())
    secure_token = f"{token}|mac={get_mac(token.encode('latin1'))}"

    return b64encode(secure_token.encode("latin1")).decode("latin1")


def handle_register():
    name = input("What is you name? ")
    animal = input("What is your favorite animal? ")

    token = generate_token(
        {
            "name": name,
            "animal": animal,
            "admin": "false",
        }
    )

    print("Here is your access token:", token)


def handle_show_animal_videos():
    user_data = parse_token(input("Enter access token: "))

    if user_data is None:
        print("Invalid token.")
        return

    print(
        f"\nHere are some {user_data['animal']} videos for you: https://www.youtube.com/results?search_query=funny+{user_data['animal']}+video+compilation"
    )


def handle_show_flag():
    user_data = parse_token(input("Enter access token: "))

    if user_data is None:
        print("Invalid token.")
        return

    if user_data["admin"] == "true":
        print("The flag is", FLAG)
    else:
        print("You are not an admin.")


def main():
    while True:
        # Show main menu

        print(
            """
        1. Register
        2. Show animal videos
        3. Show flag
        4. Exit
        """
        )

        try:
            choice = int(input("Enter your choice: "))
        except ValueError:
            print("Please enter a number next time.")
            continue
        except EOFError:
            break

        if choice == 1:
            handle_register()
        elif choice == 2:
            handle_show_animal_videos()
        elif choice == 3:
            handle_show_flag()
        elif choice == 4:
            break
        else:
            print("Please enter a valid choice.")


if __name__ == "__main__":
    main()
