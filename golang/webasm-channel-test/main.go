package main

import "fmt"
import "syscall/js"
import "time"

func main() {

	messages := make(chan string)

	callback := js.FuncOf(func(this js.Value, args []js.Value) any {

		fmt.Println("timeout fired!")
		messages <- "ping"

		return nil

	})

	js.Global().Call("setTimeout", callback, js.ValueOf(1000))
	js.Global().Call("setTimeout", callback, js.ValueOf(2000))
	js.Global().Call("setTimeout", callback, js.ValueOf(3000))

	message1 := <-messages
	fmt.Println(message1)

	message2 := <-messages
	fmt.Println(message2)

	message3 := <-messages
	fmt.Println(message3)

	for true {
		time.Sleep(1 * time.Second)
	}

}
