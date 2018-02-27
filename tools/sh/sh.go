package sh

import (
	"fmt"
	"os/exec"
)

func ExecuteShell(command string) bool {
	cmd := exec.Command("/bin/sh", "-c", command)
	_, err := cmd.Output()
	if err == nil {
		return true
	}

	return false
}

func ExecuteShellWithResultString(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	bytes, err := cmd.Output()
	if err == nil {
		return string(bytes)
	}

	fmt.Println(err)
	return ""
}
