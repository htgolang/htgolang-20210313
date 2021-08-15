package utils

import (
	"os/exec"
)

func Exec(cmd string) error {
	return exec.Command("bash", "-c", cmd).Run()
}
