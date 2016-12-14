package main

import (
	"fmt"
	"os/exec"
	"time"
)

type statusSwitch struct {
	onCode  string
	offCode string
	name    string
}

type switcher interface {
	On() error
	Off() error
}

type multiswitch struct{ switches []statusSwitch }

func (m multiswitch) On() error {
	for _, s := range m.switches {
		s.On()
	}
	return nil
}

func (m multiswitch) Off() error {
	time.Sleep(10 * time.Second)
	for _, s := range m.switches {
		s.Off()
	}
	return nil
}

func (s statusSwitch) On() error {
	return sendCode(s.onCode)
}
func (s statusSwitch) Off() error {
	return sendCode(s.offCode)
}

//This command needs to be run as root, so setuid it.

const path string = "/home/pi/433Utils/RPi_utils/codesend"

func sendCode(code string) error {
	for i := 1; i <= 3; i++ {
		c := exec.Command(path, code)
		c.Run()
		fmt.Println(c)
	}

	return nil
}
