package main

import (
	"log"
	"os/exec"
	"time"
)

// var threeSecond = time.Duration(time.Duration(2) * time.Second)
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func main() {
	go func() {
		for {
			output, err := ExecuteCommand("dir")
			if err != nil {
				log.Printf("Error: %v", err)
			} else {
				log.Printf("Output: %v", output)
			}
			time.Sleep(2 * time.Second)
			log.Printf("Run and wait 2 second !")
		}
	}()
	select {}
}
