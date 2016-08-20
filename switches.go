package main

import "os/exec"

type statusSwitch struct {
	onCode  string
	offCode string
}

func (s *statusSwitch) On() error {
	return sendCode(s.onCode)
}
func (s *statusSwitch) Off() error {
	return sendCode(s.offCode)
}

//This command needs to be run as root, so setuid it.

const path string = "/home/pi/433Utils/RPi_utils/codesend"

func sendCode(code string) error {
	c := exec.Command(path, code)
	c.Run()
	return nil
}
