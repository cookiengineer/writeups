package selects

import "time"

func Racer(url_a string, url_b string) (winner string, error error) {
	return WebsiteRacer(url_a, url_b, 10 * time.Second)
}
