package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"syscall"
)

const (
	MUST_BE_EUID      = 0
	MRUND_SOCKET_PATH = "/run/mrund.sock"
)

func main() {
	hasPermissions := checkEUID(MUST_BE_EUID)
	if hasPermissions == false {
		fmt.Printf("You dont have the necessary permissions. \nEUID: %d \n", os.Geteuid())
		os.Exit(1)
	}

	err := deleteSocketPathIfExists(MRUND_SOCKET_PATH)
	if err != nil {
		fmt.Printf("Unable to delete socket at path: %s\n", MRUND_SOCKET_PATH)
		os.Exit(1)
	}

	listener, err := tryCreateUnixSocket(MRUND_SOCKET_PATH)
	if err != nil {
		fmt.Printf("Unable to create socket at path %s\n", MRUND_SOCKET_PATH)
		os.Exit(1)
	}

	startListener(listener)
}

func startListener(listener *net.Listener) {
	for {
		conn, connectionErr := (*listener).Accept()
		if connectionErr != nil {
			fmt.Printf("Connection error: %s\n", connectionErr)
			continue
		}

		go handleConnection(conn)
	}
}

func tryCreateUnixSocket(socketPath string) (*net.Listener, error) {
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Printf("Error creating socket: %v\n", err)
		return nil, err
	}

	fmt.Printf("Listening on socket: %s\n", socketPath)
	return &listener, nil
}

func checkEUID(euidToCheck int) bool {
	euid := syscall.Geteuid()
	if euid != euidToCheck {
		return false
	}
	return true
}

func deleteSocketPathIfExists(socketPath string) error {
	_, err := os.Stat(socketPath)
	socketExists := err == nil
	if socketExists {
		err := os.Remove(socketPath)
		if err != nil {
			return err
		}
		return nil
	}
	return err
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
