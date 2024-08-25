package main

import "encoding/hex"
import "fmt"
import "os"
import "strings"

func decrypt(message []byte, key []byte) []byte {

	var result []byte

	for m := 0; m < len(message); m++ {
		result = append(result, message[m] ^ key[m % len(key)])
	}

	return result

}

func main() {

	buffer, err1 := os.ReadFile("FLAG.enc")

	if err1 == nil {

		decoded, err2 := hex.DecodeString(strings.TrimSpace(string(buffer)))

		if err2 == nil {

			// Prefix is GPNCTF{

			expect := []byte("GPNCTF{")
			prefix := decoded[0:5]

			key := make([]byte, 5)

			for p := 0; p < len(prefix); p++ {
				key[p] = prefix[p] ^ expect[p]
			}

			fmt.Println("Key is: ")
			fmt.Println(key)

			fmt.Println("Decrypted message is:")
			fmt.Println(string(decrypt(decoded, key)))


		}

	}

}
