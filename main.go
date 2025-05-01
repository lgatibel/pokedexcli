package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	next     string
	previous string
}

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")
	return nil
}

const (
	apiUrl    = "https://pokeapi.co/api/v2/location-area/"
	pageLimit = 20
)

type ResultLocation struct {
	Count    int
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func commandMapB(config *config) error {
	locations := &ResultLocation{}
	var url string

	if config.previous == "" {
		return fmt.Errorf("you're on the first page")
	} else {
		url = config.previous + fmt.Sprintf("?limit=%d", pageLimit)
	}

	if err := getLocations(url, locations); err != nil {
		return err
	}
	config.next = locations.Next
	config.previous = locations.Previous
	return nil
}

func commandMap(config *config) error {
	locations := &ResultLocation{}
	var url string

	if config.next == "" {
		url = apiUrl + fmt.Sprintf("?limit=%d", pageLimit)
	} else {
		url = config.next + fmt.Sprintf("?limit=%d", pageLimit)
	}

	if err := getLocations(url, locations); err != nil {
		return err
	}
	config.next = locations.Next
	config.previous = locations.Previous

	return nil
}

func getLocations(url string, locations *ResultLocation) error {
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("bad request: %s", err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locations); err != nil {
		return fmt.Errorf("bad parsing of maps: %s", err)
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
}

// "Welcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex",
func main() {
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
	}
	config := config{}
	fmt.Printf("%s", prompt)
	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		firstWord := words[0]
		if command, ok := commands[firstWord]; ok {
			if err := command.callback(&config); err != nil {
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
