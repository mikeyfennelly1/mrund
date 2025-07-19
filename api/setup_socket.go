package api

import (
	"fmt"
	"net"
	"os"
)

func TryCreateUnixSocket(socketPath string) (*net.Listener, error) {
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Printf("Error creating socket: %v\n", err)
		return nil, err
	}

	fmt.Printf("Listening on socket: %s\n", socketPath)
	return &listener, nil
}

func DeleteSocketPathIfExists(socketPath string) error {
	_, err := os.Stat(socketPath)
	socketExists := err == nil
	if socketExists {
		err := os.Remove(socketPath)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
