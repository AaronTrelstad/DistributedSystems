package container

import (
	"os/exec"
	"syscall"
	"os"
)

func Run(command string, args []string) error {
	cmd := exec.Command(command, args...);
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	return cmd.Run()
}
