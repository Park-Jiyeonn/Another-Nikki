package service

import "os/exec"
func RunCommand(s string) error {
	cmd := exec.Command("bash", "-c", s)
	err := cmd.Run()
	return err
}