package main

import "fmt"
import "syscall/js"
import "time"

func main() {

	messages := make(chan string)

	fetch_options := make(map[string]any)
	fetch_options["mode"] = "same-origin"

	on_success := js.FuncOf(func(this js.Value, args []js.Value) any {

		fmt.Println("on_success fired!")
		messages <- "success"

		return nil

	})

	on_failure := js.FuncOf(func(this js.Value, args []js.Value) any {

		fmt.Println("on_failure fired!")
		messages <- "failure"

		return nil

	})

	go js.Global().Call("fetch", js.ValueOf("http://localhost:3000/api/test"), js.ValueOf(fetch_options)).Call("then", on_success).Call("catch", on_failure)

	message1 := <-messages
	fmt.Println(message1)

	for true {
		time.Sleep(1 * time.Second)
	}

}
