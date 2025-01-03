package main

import (
	"fmt"
	"log"
	"os/exec"
)

func executeCommand(command string, args ...string) error {
    cmd := exec.Command(command, args...)
    
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("command failed: %s, error: %v, output: %s", 
            command, err, string(output))
    }
    
    return nil
}

func main() {
    // Execute a docker command
    if err := executeCommand("docker", "ps"); err != nil {
        log.Fatal(err)
    }
    
    // Copy a file
    if err := executeCommand("cp", "source.txt", "dest.txt"); err != nil {
        log.Fatal(err)
    }
}