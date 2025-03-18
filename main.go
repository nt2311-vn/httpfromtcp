package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatalf("cannot open messages file: %v", err)
	}

	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("cannot read 8 bytes from file: %v", err)
		}

		fmt.Printf("read: %s\n", buf[:n])
	}
}
