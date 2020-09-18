package main

import "fmt"

// Menu configuration
var decoration string = "#"
var menuWidth int = 70
var lineThicness int = 2

func mainMenu() {
	menuText := []string{
		"Pangramchecker",
		"Welcome to the Go-Pangramchecker!",
		"Please enter a Sentence to begin."}
	menuDescription := []string{"A pangram is a Sentence in which each",
		"Character is present one atleast one Time."}

	fmt.Print(decor(menuText, menuDescription, decoration, lineThicness, menuWidth), "\n")
}

func endMenu(pangram *pangram) {
	var stateString string

	if pangram.perfectPangram == true {
		stateString = "is a perfect Pangram."
	} else if pangram.pangram == true {
		stateString = "is just a normal Pangram."
	} else {
		stateString = "i not a Pangram."
	}

	menuText := []string{
		"Pangramchecker",
		"Your sentence:",
		pangram.sentence}
	menuPangram := []string{stateString}

	fmt.Print(decor(menuText, menuPangram, decoration, lineThicness, menuWidth), "\n")
}
