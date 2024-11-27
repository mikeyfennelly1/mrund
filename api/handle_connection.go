package api

import (
	"bufio"
	"fmt"
	"net"
)

func StartListener(listener *net.Listener) {
	for {
		conn, connectionErr := (*listener).Accept()
		if connectionErr != nil {
			fmt.Printf("Connection error: %s\n", connectionErr)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from connection %s\n", err)
			return
		}

		fmt.Printf("Received message: %s", message)
		response := fmt.Sprintf("Echo: %s", message)
		conn.Write([]byte(response))
	}
}
