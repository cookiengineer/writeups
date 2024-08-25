package concurrency

import "reflect"
import "testing"

func mockWebsiteChecker(url string) bool {

	if url == "http://localhost:3000/" {
		return false
	}

	return true

}

func TestCheckWebsites(t *testing.T) {

	urls := []string{
		"https://google.com/",
		"https://cookie.engineer/",
		"https://tholian.network/en/index.html",
		"http://localhost:3000/",
	}

	want := map[string]bool{
		"https://google.com/":                   true,
		"https://cookie.engineer/" :             true,
		"https://tholian.network/en/index.html": true,
		"http://localhost:3000/":                false,
	}

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %v but got %v", want, got)
	}

}
