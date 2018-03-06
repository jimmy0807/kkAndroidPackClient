package sh

import (
	"fmt"
	"os/exec"
	"runtime"
)

func ExecuteShell(command string) bool {
	cmd := exec.Command("/bin/sh", "-c", command)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
		fmt.Println("Windows command")
		fmt.Println(command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}

	_, err := cmd.Output()
	if err == nil {
		print("成功了")
		return true
	} else {
		print(err)
	}

	return false
}

func ExecuteShellWithResultString(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("/bin/sh", "-c", command)
	}

	bytes, err := cmd.Output()
	if err == nil {
		return string(bytes)
	}

	fmt.Println(err)
	return ""
}
