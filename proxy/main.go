package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	path := "/home/pi/433Utils/RPi_utils/codesend"
	args := []string{"codesend", "7289615"}
	env := os.Environ()
	err := syscall.Exec(path, args, env)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1e10)
	argsOff := []string{"codesend", "7289607"}
	err = syscall.Exec(path, argsOff, env)
	if err != nil {
		fmt.Println(err)
	}
}
