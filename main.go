package main

import (
	"fmt"
	"github.com/mikeyfennelly1/mrund/api"
	"github.com/mikeyfennelly1/mrund/container"
	"github.com/mikeyfennelly1/mrund/utils"
	"os"
)

const (
	MUST_BE_EUID      = 0
	MRUND_SOCKET_PATH = "/run/mrund.sock"
)

// ensures that the user running the binary has the
// necessary permissions and starts server
func main() {
	hasPermissions := utils.CheckEUID(MUST_BE_EUID)
	if hasPermissions == false {
		fmt.Printf("You dont have the necessary permissions. \nEUID: %d \n", os.Geteuid())
		os.Exit(1)
	}

	// delete socket path if it exists already
	err := api.DeleteSocketPathIfExists(MRUND_SOCKET_PATH)
	if err != nil {
		panic(fmt.Sprintf("Unable to delete socket at path: %s\n", MRUND_SOCKET_PATH))
	}

	// create listening server on unix socket at path MRUND_SOCKET_PATH
	//listener, err := api.TryCreateUnixSocket(MRUND_SOCKET_PATH)
	utils.ExitIfErr(&err, fmt.Sprintf("Unable to create socket at path %s\n", MRUND_SOCKET_PATH))

	container.CreateVethPairAndBridge()
}
