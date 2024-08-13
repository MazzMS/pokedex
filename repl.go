package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mazzms/pokedex/commands"
)

func startRepl(cfg *commands.Config) {
	fmt.Println("Starting the pokedex...")
	scanner := bufio.NewScanner(os.Stdin)
	commands := commands.GetCommands()
	fmt.Println("---------------------------")
	fmt.Println("Welcome to the BEST pokedex")
	fmt.Println("Use 'exit' to quit")
	fmt.Println("Use 'help' to known possible commands")
	fmt.Println()
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := clearInput(input)

		if len(words) == 0{
			continue
		}

		possibleCommand := words[0]
		args := words[1:]

		if command, ok := commands[possibleCommand]; ok {
			err := command.Callback(cfg, args...)
			if err != nil {
				// print the error in default color initially
				fmt.Printf("%s\r", err)
				time.Sleep(100 * time.Millisecond)
				// print the error in red for ~half a second
				fmt.Printf("\033[31m%s\033[0m\r", err)
				time.Sleep(400 * time.Millisecond)
				//  left the error printed in the initial color 
				fmt.Printf("%s\r\n", err)
				fmt.Println()
			}
		} else {
			fmt.Printf("%q is not a valid command, please try again\n", input)
		}
	}
}

func clearInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered)
	return words
}
