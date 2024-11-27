package core

import (
	"os/exec"
	"syscall"
)

// creates isolated process with namespaces
func createIsolatedProcess() {
	proc := exec.Cmd{}
	proc.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS | syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWCGROUP | syscall.CLONE_NEWUSER,
	}
	proc.Run()
}
