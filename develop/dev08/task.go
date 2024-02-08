package dev

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

func shell() {
	for {
		fmt.Print(getPrompt())
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "\\quit" {
			break
		}

		runCommand(input)

	}
}

func getPrompt() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s@%s:%s$ ", user.Username, "shell", getWorkingDirectory())
}

func getWorkingDirectory() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wd
}

func runCommand(input string) {
	input = strings.TrimSpace(input)

	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "cd":
		if len(args) > 1 {
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("cd:", err)
			}
		} else {
			fmt.Println("cd: missing argument")
		}
	case "pwd":
		fmt.Println(getWorkingDirectory())
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) > 1 {
			pid := args[1]
			process, err := os.FindProcess(pidToInt(pid))
			if err != nil {
				fmt.Println("kill:", err)
				return
			}
			err = process.Kill()
			if err != nil {
				fmt.Println("kill:", err)
			}
		} else {
			fmt.Println("kill: missing argument")
		}
	case "ps":
		psCommand := exec.Command("ps", "-e")
		psCommand.Stdout = os.Stdout
		psCommand.Stderr = os.Stderr
		psCommand.Run()
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Command failed:", err)
		}
	}
}

func pidToInt(pid string) int {
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return -1
	}
	return pidInt
}
