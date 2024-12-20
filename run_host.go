//go:build windows && !nonhost

package main

import (
	"os/exec"
)

// Run the native WSL command
func command(distribution Distribution, config Config) (*exec.Cmd, error) {
	// Append the additional arguments to the command
	args := append([]string{"-d", distribution.Name}, config.WslArgs...)
	// Set the title of the terminal window
	err := exec.Command("cmd", "/c", "title", distribution.Name).Run()
	if err != nil {
		return nil, err
	}
	return exec.Command("wsl", args...), nil
}
