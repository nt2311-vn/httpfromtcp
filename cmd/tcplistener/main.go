package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(conn io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)
		defer conn.Close()

		var currentLine string
		buf := make([]byte, 8)

		for {
			n, err := conn.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println("Error reading file:", err)
				return
			}

			chunk := string(buf[:n])
			parts := strings.Split(chunk, "\n")

			for _, part := range parts[:len(parts)-1] {
				lines <- currentLine + part
				currentLine = "" // Reset after sending a complete line
			}

			currentLine += parts[len(parts)-1]
		}

		if currentLine != "" {
			lines <- currentLine
		}
	}()

	return lines
}

func main() {
	ln, err := net.Listen("tcp", ":42069")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer ln.Close()

	fmt.Println("Listening on port :42069...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error accepting connection")
		}

		fmt.Println("Connection Accepted")

		go func(c net.Conn) {
			for line := range getLinesChannel(c) {
				fmt.Println(line)
			}
		}(conn)
	}
}
