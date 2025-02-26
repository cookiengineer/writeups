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

	http.HandleFunc("/api/test", func(response http.ResponseWriter, request *http.Request) {

		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("{\"message\": \"success\"}"))

	})

	err := http.ListenAndServe(":3000", nil)

	if err == nil {
		fmt.Println("Exiting...")
	}

}
