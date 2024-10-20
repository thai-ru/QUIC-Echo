# **QUIC Echo Server & Client**

## 📄 **Project Overview**

This project is an implementation of a simple **QUIC**-based echo server and client in **Go**. QUIC (Quick UDP Internet Connections) is a transport protocol designed to provide fast, reliable, and secure communication over the internet. The echo server and client use QUIC to establish low-latency, secure connections, making it ideal for scenarios like real-time messaging, gaming, financial services, and IoT applications.

## ✨ **Features**
- **Low-latency communication** using QUIC protocol.
- Secure connections via **TLS encryption**.
- Simple echo functionality for testing QUIC-based communication.
- Written in **Go**, ensuring high performance and ease of use.
- Easily extendable for various use cases, such as gaming, fintech, or IoT.

## 🚀 **Use Cases**
- **Banking/Fintech**: Secure, real-time payment processing.
- **Gaming**: Low-latency game servers and real-time multiplayer.
- **IoT Systems**: Reliable communication for devices with constrained networks.
- **Chat Apps**: Fast messaging with minimal overhead.
- **Secure Communication**: Built-in encryption for privacy-first applications.

## 📦 **Installation**

### Prerequisites
Make sure you have the following installed:
- **Go** 1.18+ [Install Go](https://golang.org/doc/install)
- **QUIC-go** library (automatically included in `go.mod`)
- Certificates for **TLS encryption** (generate if needed)

### Steps

1. Clone the repository:
   ```bash
   git clone https://github.com/thai-ru/QUIC-Echo.git
   cd QUIC-Echo
   ```

2. Build the project:
   ```bash
   go build cmd/server/main.go
   go build cmd/client/main.go
   ```

3. Generate self-signed certificates for TLS:
   ```bash
   openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
   ```

## ⚙️ **Usage**

### 1. Start the server
To run the QUIC echo server, execute the following command:
```bash
go run cmd/server/main.go
```
You should see the output:
```bash
QUIC server listening on localhost:4242
```

### 2. Start the client
In a new terminal, run the client:
```bash
go run cmd/client/main.go
```
You'll be prompted to send messages to the server. Example:
```bash
Enter message (or type 'exit' to quit): Hello QUIC server!
Client sent: Hello QUIC server!
Client received: Echo: Hello QUIC server!
```

### 3. Exiting the client
To close the client, type `exit`:
```bash
Enter message (or type 'exit' to quit): exit
```

## 🛠 **Project Structure**

```
QUIC-Echo/
│
├── cmd/
│   ├── client/
│   │   └── main.go         # Client-side QUIC logic
│   └── server/
│       └── main.go         # Server-side QUIC logic
│
├── server.crt              # TLS certificate (example)
├── server.key              # TLS private key (example)
├── go.mod                  # Go module file
└── README.md               # Project documentation
```

## 🤝 **Contributing**

We welcome contributions! If you'd like to contribute to the project, follow these steps:

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -am 'Add a new feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a Pull Request.

Please ensure that your changes pass existing tests and that you write new tests if applicable.

## 📝 **License**

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

## ✉️ **Contact**

For any questions or support, please reach out via:
- **Email**: robinsonthairu@gmail.com
- **GitHub Issues**: [Issues](https://github.com/thai-ru/QUIC-Echo/issues)