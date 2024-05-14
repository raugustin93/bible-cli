package main

import (
	"errors"
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
	valid := 1 < len(args) && len(args) < 4
	if !valid {
		return errors.New("there must be 2 or 3 arguments for read command. *Book* *Chapter* *Verse?*")
	}
	_, err := cfg.bibleApiClient.GetVerse(args[0], args[1], args[2])
	if err != nil {
		return err
	}

	return nil
}
