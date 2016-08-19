package main

import (
	"fmt"
	"time"
)

func main() {
	switch1 := statusSwitch{7289615, 7289607}
	err := switch1.On()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(1e10)
	err = switch1.Off()
	if err != nil {
		fmt.Println(err)
	}
}
