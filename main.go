package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := os.Args[1:]

	cmd := exec.Command(command[0], command[1:]...)

	// parallel chalana or output dikhana
	// output, err := cmd.Output()
	go func() {
		output, err := cmd.Output()

		if err != nil {
			fmt.Println("runnning command error")
			fmt.Println(err)
			return
		}
	}()

	fmt.Println("hello")

	// monitor for CPU usage
	// thread to check cpu

	// monitor for disk usage
	// thread to check disk

	// monitor for mem usage
	// thread to check mem

}
