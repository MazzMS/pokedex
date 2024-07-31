package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

		if command, ok := commands[possibleCommand]; ok {
			err := command.Callback(cfg)
			if err != nil {
				fmt.Println(err)
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
