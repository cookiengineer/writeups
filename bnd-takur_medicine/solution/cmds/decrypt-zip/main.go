package main

import "fmt"
import "os"
import "strconv"

var zip_headers [][]byte = [][]byte{
	// See: https://en.wikipedia.org/wiki/ZIP_(file_format)
	[]byte("PK\x03\x04"),
	[]byte("PK\x05\x06"),
	[]byte("PK\x07\x08"),
}

func main() {

	if len(os.Args) == 2 {

		file := os.Args[1]
		buffer, err2 := os.ReadFile(file)

		if err2 == nil {

			// "ransomware.py" uses fixed-size 4 bytes key
			header := buffer[0:4]
			possible_keys := make([][]byte, 0)
			possible_buffers := make([][]byte, 0)

			fmt.Println("Header of encrypted file is:")
			fmt.Println(header)

			fmt.Println("Generating possible keys...")

			for z := 0; z < len(zip_headers); z++ {

				expected_header := zip_headers[z]

				var possible_key = make([]byte, 0)

				for e := 0; e < len(expected_header); e++ {
					possible_key = append(possible_key, expected_header[e] ^ header[e])
				}

				fmt.Println(possible_key)

				possible_keys = append(possible_keys, possible_key)

			}

			for p := 0; p < len(possible_keys); p++ {

				possible_key := possible_keys[p]
				var possible_buffer = make([]byte, 0)

				fmt.Println(len(possible_key))

				for b := 0; b < len(buffer); b++ {
					possible_buffer = append(possible_buffer, buffer[b] ^ possible_key[b % 4])
				}

				possible_buffers = append(possible_buffers, possible_buffer)

			}

			for p := 0; p < len(possible_buffers); p++ {

				buffer := possible_buffers[p]
				name := "geheim-" + strconv.Itoa(p) + ".zip"

				err := os.WriteFile(name, buffer, 0666)

				if err == nil {
					fmt.Println("Possible original ZIP file dumped as: \"" + name + "\"")
				}

			}

		} else {
			fmt.Println("Usage: <program> ./path/to/file.zip.crypt")
		}

	}

}
