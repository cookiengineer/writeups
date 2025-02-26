package main

import "embed"
import "fmt"
import "io/fs"
import "net/http"

//go:embed public/*
var embedded_filesystem embed.FS

func main() {

	fsys, _ := fs.Sub(embedded_filesystem, "public")

	fmt.Println("Listening on http://localhost:3000")
	http.Handle("/", http.FileServer(http.FS(fsys)))

	err := http.ListenAndServe(":3000", nil)

	if err == nil {
		fmt.Println("Exiting...")
	}

}
