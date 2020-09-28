package main

import (
	"fmt"
	"os"
	"strings"
)

// Menu configuration
var decoration string = "@"
var menuWidth int = 70
var lineThicness int = 2

func mainMenu() {
	menuText := []string{
		"Pangramchecker",
		"Welcome to the Go-Pangramchecker!",
		"Please enter a Sentence to check."}
	menuDescription := []string{
		"A pangram is a Sentence in which each",
		"Character is present one atleast one Time."}

	fmt.Print(decor(menuText, menuDescription, decoration, lineThicness, menuWidth), "\n")
}

func endMenu(pangram *pangram) {
	var stateString string
	var missingString string

	if pangram.perfectPangram == false && pangram.pangram == true {
		missingString = "There are no Characters missing."
	} else if pangram.perfectPangram == false && pangram.pangram == false {
		missingString = "Missing: " + strings.Join(pangram.missing, ",")
	}

	if pangram.perfectPangram == true {
		stateString = "is a perfect Pangram."
	} else if pangram.pangram == true {
		stateString = "is just a normal Pangram."
	} else {
		stateString = "is not a Pangram."
	}

	menuText := []string{
		"Pangramchecker",
		"Your sentence:",
		pangram.sentence}
	menuPangram := []string{
		stateString,
		missingString,
		"Do you want to Exit?",
		"1) Yes",
		"2) No"}

	fmt.Print(decor(menuText, menuPangram, decoration, lineThicness, menuWidth), "\n")
	switch getIO() {
	case "1":
		os.Exit(0)
	case "2":
		pangramchecker(pangram)
	}
}
