package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	fmt.Println("Starting the pokedex...")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("---------------------------")
	fmt.Println("Welcome to the BEST pokedex")
	fmt.Println("Use 'exit' to quit")
	fmt.Print("pokedex > ")
	for scanner.Scan() {
		command := scanner.Text()

		fmt.Print("pokedex > ", command)
	}
}
