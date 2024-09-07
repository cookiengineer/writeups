package main

import "github.com/yeka/zip"
import _ "embed"
import "fmt"
import "io/ioutil"
import "log"
import "os"
import "strconv"
import "strings"
import "sync"

//go:embed passwords.txt
var EMBEDDED_PASSWORDS []byte
var PASSWORDS []string

func init() {

	lines := strings.Split(string(EMBEDDED_PASSWORDS), "\n")

	for l := 0; l < len(lines); l++ {

		line := strings.TrimSpace(lines[l])

		if line != "" {
			PASSWORDS = append(PASSWORDS, line)
		}

	}

}

func decrypt(reader *zip.ReadCloser, password string) bool {

	var result bool = false

	for _, file := range reader.File {

		if file.IsEncrypted() {
			file.SetPassword(password)
		}

		stream, err1 := file.Open()

		if err1 == nil {

			buffer, err2 := ioutil.ReadAll(stream)

			if err2 == nil {
				fmt.Println("Successfully decrypted " + strconv.Itoa(len(buffer)) + " bytes")
				fmt.Println("Password was \"" + password + "\"")
				result = true
			}

			defer stream.Close()

		}

	}

	return result

}

func main() {

	var channel = make(chan bool)

	if len(os.Args) == 2 {

		file := os.Args[1]

		reader, err1 := zip.OpenReader(file)

		if err1 != nil {
			log.Fatal(err1)
		}

		var wait_group sync.WaitGroup

		for p := 0; p < len(PASSWORDS); p++ {

			var password = PASSWORDS[p]
			wait_group.Add(1) // bool size is 1 Byte

			go func (password string) {

				defer wait_group.Done()
				result := decrypt(reader, password)
				channel <- result

			}(password)

			result := <-channel

			if result == true {
				fmt.Println("Password found:", password)
				break
			} else {
				fmt.Println("Tried out:", password)
			}

			wait_group.Wait()

		}

	}

}
