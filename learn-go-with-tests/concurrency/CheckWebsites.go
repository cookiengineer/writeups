package concurrency

type WebsiteChecker func(string) bool
type WebsiteCheckerResult struct {
	string
	bool
}

func CheckWebsites(checker WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	channel := make(chan WebsiteCheckerResult)

	for _, url := range urls {

		go func(url string) {
			channel <- WebsiteCheckerResult{url, checker(url)}
		}(url)

	}

	for i := 0; i < len(urls); i++ {
		tmp := <-channel
		results[tmp.string] = tmp.bool
	}

	return results

}
