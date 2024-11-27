package utils

import (
	"fmt"
	"os"
	"syscall"
)

func CheckEUID(euidToCheck int) bool {
	euid := syscall.Geteuid()
	if euid != euidToCheck {
		return false
	}
	return true
}

func ExitIfErr(err *error, messageIfErr string) {
	if (*err) != nil {
		fmt.Printf(messageIfErr)
		os.Exit(1)
	}
}
