package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/lgatibel/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func startRepl(config *pokeapi.Config) {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next maps",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous maps",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Display the pokemon list in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon",
			callback:    commandCatch,
		},
	}
	fmt.Printf("%s", prompt)
	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		firstWord := words[0]
		var param string
		if len(words) > 1 {
			param = words[1]
		}
		if command, ok := commands[firstWord]; ok {
			if err := command.callback(config, param); err != nil {
				fmt.Println(err.Error())
			}

			if command.name == "help" {
				for name, command := range commands {
					fmt.Printf("%s: %s\n", name, command.description)
				}
			}
		} else {
			fmt.Printf("Unknown command: %s\n", firstWord)
		}
		fmt.Printf("%s", prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
