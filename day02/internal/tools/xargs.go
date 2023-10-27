package tools

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func RunXargs() {
	scanner := bufio.NewScanner(os.Stdin)

	command := os.Args[1]
	args := os.Args[2:]

	for scanner.Scan() {

		text := scanner.Text()
		if text == "" {
			break
		}

		command := exec.Command(command, append(args, text)...)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()

		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
