package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Split(strings.Trim(strings.ToLower(text), " "), " ")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	baseMesage := "Your command was: %s\n"
	fmt.Printf("%s", prompt)
	for scanner.Scan() {
		text := scanner.Text()
		words := cleanInput(text)
		firstWord := words[0]
		fmt.Printf(baseMesage, firstWord)
		fmt.Printf("%s", prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
}
