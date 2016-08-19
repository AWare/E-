package main

import (
	"os"
	"strconv"
	"syscall"
)

type statusSwitch struct {
	onCode  int
	offCode int
}

func (s *statusSwitch) On() error {
	return sendCode(s.onCode)
}
func (s *statusSwitch) Off() error {
	return sendCode(s.offCode)
}

//This command needs to be run as root, so setuid it.

const path string = "/home/pi/433Utils/RPi_utils/codesend"

func sendCode(code int) error {
	env := os.Environ()
	args := []string{"codesend", strconv.Itoa(code)}
	return syscall.Exec(path, args, env)
}
