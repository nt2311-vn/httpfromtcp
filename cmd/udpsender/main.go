package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Resolve the UDP address
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	// Dial UDP connection
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("Failed to dial UDP: %v", err)
	}
	defer conn.Close()

	// Create a reader to get input from stdin
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("UDP sender ready. Type messages and hit Enter to send.")
	fmt.Println("Run 'nc -u -l 42069' in another terminal to receive messages.")

	for {
		// Print prompt
		fmt.Print("> ")

		// Read user input
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading input: %v", err)
			continue
		}

		// Send the message over UDP
		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error sending UDP packet: %v", err)
		}
	}
}
