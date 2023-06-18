package main

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/webview/webview"
)

var w webview.WebView

func main() {
	debug := true
	w = webview.New(debug)
	defer w.Destroy()
	w.SetTitle("WebView Example with Bindings")
	w.SetSize(800, 600, webview.HintNone)

	// Load the start page, index.html
	w.Navigate("file://" + pathToStartPage())

	// Expose the Go function printSomething as "ps"
	// in our WebView window. A JavaScript call to "ps"
	// will invoke the Go function printSomething.
	w.Bind("ps", printSomething)

	// Run the app.
	w.Run()
}

// printSomething prints the name from the JavaScript form
// to STDOUT and back to the WebView HTML page.
func printSomething(name string) {
	now := time.Now().Format(time.Stamp)
	fmt.Println(textMessage(name, now))
	w.Eval(fmt.Sprintf("setDivContent('output', '%s')", htmlMessage(name, now)))
}

// Returns a plain text message to display in the console.
func textMessage(name, time string) string {
	return fmt.Sprintf(">>> [%s] Hello, %s.", time, name)
}

// Returns an HTML message to display in the webview.
func htmlMessage(name, time string) string {
	escapedName := strings.ReplaceAll(name, "'", "&#39;")
	return fmt.Sprintf(`Hello, <b>%s</b>. The current time is <span class="blue">%s</span>. Check your console window for a message.`, escapedName, time)
}

// pathToStartPage returns the absolute path the index.html page
// we want to show when the app starts up. This works when we're
// running the app through 'go run main.go' because runtime.Caller()
// returns the path of this source file, not the path of the temp
// file created by 'go run'.
func pathToStartPage() string {
	_, currentFile, _, _ := runtime.Caller(0)
	dir := path.Dir(currentFile)
	return path.Join(dir, "index.html")
}
