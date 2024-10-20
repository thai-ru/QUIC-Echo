package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/quic-go/quic-go"
)

func main() {
	// QUIC server address
	addr := "localhost:4242"

	// Establish a QUIC connection
	conn, err := quic.DialAddr(
		context.Background(),
		addr,
		&tls.Config{
			InsecureSkipVerify: true,           // Skip certificate verification (for testing)
			NextProtos:         []string{"h3"}, // h3 - identifier of the application protocol
		},
		&quic.Config{},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.CloseWithError(0, "Normal closure")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')

		if message == "exit\n" {
			break
		}

		// Open a new stream for communication
		stream, err := conn.OpenStreamSync(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Client sent: %s", message)
		_, err = stream.Write([]byte(message))
		if err != nil {
			log.Fatal(err)
		}

		// Read the echoed message from the server
		buf := make([]byte, 1024)
		n, err := stream.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Client received: %s\n", string(buf[:n]))

		// Close the stream after communication
		stream.Close()
	}
}
