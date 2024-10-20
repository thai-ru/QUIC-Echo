package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/quic-go/quic-go"
	"io"
	"log"
)

func generateTLSConfig() *tls.Config {
	/*
		generate self-signed certs using open ssl :
		openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes
	*/

	// You can also automate this e.g:
	/*
		key, err := rsa.GenerateKey(rand.Reader, 1024)
			if err != nil {
				panic(err)
			}
			template := x509.Certificate{SerialNumber: big.NewInt(1)}
			certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
			if err != nil {
				panic(err)
			}
			keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
			certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	*/

	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")

	if err != nil {
		log.Fatal(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h3"}, // h3 - identifier of the application protocol
	}
}

func main() {
	// Listen on QUIC port
	addr := "localhost:4242"
	lis, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("QUIC server listening on %s\n", addr)

	for {
		// Accept a new connection
		conn, err := lis.Accept(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		go handleSession(conn)
	}
}

func handleSession(conn quic.Connection) {
	defer conn.CloseWithError(0, "closing connection")

	for {
		// Accept a stream (client sends data over this stream)
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			if err == io.EOF {
				fmt.Println("No more streams; client closed connection.")
				return
			}
			log.Fatal(err)
		}
		defer stream.Close()

		// Read message from the stream
		buf := make([]byte, 1024)
		n, err := stream.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Stream closed by client.")
				return
			}
			log.Fatal(err)
		}
		fmt.Printf("Received: %s\n", string(buf[:n]))

		// Echo the message back to the client
		_, err = stream.Write([]byte(fmt.Sprintf("Echo: %s", string(buf[:n]))))
		if err != nil {
			log.Fatal(err)
		}
	}
}
