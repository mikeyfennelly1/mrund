package container

import (
	"os"
	"os/exec"
	"syscall"
)

func createBasicContainer(hostPort int, internalPort int) {
	containerCMD := exec.Cmd{}
	containerCMD.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	containerCMD.Stdin = os.Stdin
	containerCMD.Stdout = os.Stdout
	containerCMD.Stderr = os.Stderr

	containerCMD.Run()
}
