package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearBash() {
	var cmd *exec.Cmd

	// Check the operating system
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	} else {
		cmd = exec.Command("clear") // For Unix/Linux
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}