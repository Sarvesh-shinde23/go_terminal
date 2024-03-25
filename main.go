package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func getUserInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}
	return strings.TrimSpace(input), nil // Remove trailing newline
}
func main() {
	for {
		// Get user input
		command, err := getUserInput()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if command == "exit" {
			break // Exit the loop on "exit" command
		}

		// Execute the command using os/exec
		cmd := exec.Command("powershell", "-c", command) // Use a shell for flexibility
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("Error creating stdout pipe:", err)
			continue
		}

		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting command:", err)
			continue
		}

		// Read and print output
		output, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("Error reading output:", err)
			continue
		}
		fmt.Println(string(output))

		cmd.Wait() // Wait for the command to finish
	}
	fmt.Println("Exiting...")
}
