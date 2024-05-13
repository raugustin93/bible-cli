package main

import (
	"fmt"
	"os"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Bible Cli: ")
	fmt.Println("Here are available Commands: ")
	fmt.Println("")
	commands := getCommands()
	for _, curr := range commands {
		fmt.Println("")
		fmt.Printf("	--%s : '%s'\n", curr.name, curr.description)
	}
	fmt.Println("")
	return nil
}

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("Exited From Terminal")
	os.Exit(0)
	return nil
}

func callbackRead(cfg *config, args ...string) error {
	fmt.Println("Reading TODO...")
	return nil
}
