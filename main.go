package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// getLinesChannel reads from an io.ReadCloser in chunks and sends full lines to a channel.
func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer close(lines)
		defer f.Close()

		var currentLine string
		buf := make([]byte, 8)

		for {
			n, err := f.Read(buf)
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
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	for line := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", line)
	}
}
