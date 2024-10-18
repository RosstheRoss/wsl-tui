//go:build !windows || nonhost

package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
)

func command(distribution Distribution, config Config) (*exec.Cmd, error) {
	if distribution.Username == "" || distribution.Port == 0 {
		return nil, errors.New("SSH username or port not set, not running.")
	}
	args := append([]string{fmt.Sprintf("%s@%s", distribution.Username, config.HostIP), "-p", strconv.Itoa(distribution.Port)}, config.SshArgs...)
	return exec.Command("ssh", args...), nil
}
