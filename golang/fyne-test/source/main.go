package main

import fyne_app "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2/canvas"
import "fyne.io/fyne/v2/container"
import "fyne.io/fyne/v2/layout"
import "fyne.io/fyne/v2/widget"
import "fmt"
import "image/color"
import "time"

func destroy() {
	fmt.Println("Exited!")
}

func main() {

	app := fyne_app.New()
	win := app.NewWindow("Hello World")

	var label = widget.NewLabel("Welcome")

	text1 := canvas.NewText("Hello", color.Black)
	text2 := canvas.NewText("World", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer())

	go func() {

		for range time.Tick(time.Second) {
			formatted := time.Now().Format("Time: 03:04:05")
			label.SetText(formatted)
		}

	}()

	text3 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text3, layout.NewSpacer())
	win.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	win.Show()

	app.Run()
	destroy()

}
