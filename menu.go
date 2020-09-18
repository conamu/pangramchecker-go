package main

import "fmt"

// Menu configuration
var decoration string = "#"
var menuWidth int = 70
var lineThicness int = 2

func mainMenu() {
	menuText := []string{"Welcome to the Go-Pangramchecker!",
		"Please enter a Sentence to begin."}
	menuDescription := []string{"A pangram is a Sentence in which each",
		"Character is present one atleast one Time."}

	fmt.Print(decor(menuText, menuDescription, decoration, lineThicness, menuWidth))
}

kljh
