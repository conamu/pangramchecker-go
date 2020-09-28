package main

import (
	"regexp"
	"strings"
)

func (pangram *pangram) initPangramData() {
	// Reset states for the new sentence.
	pangram.pangram = false
	pangram.perfectPangram = false
	pangram.presentCharN = 0
	pangram.presentUniqeCharN = 0
	// Create a map to track chars and their use count
	pangram.presentChars = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}
	// Create the Array for the missing chars to be fileld in
	pangram.missing = make([]string, 26)
}

// Clean the string beeing input and make it all lowercase
func (pangram *pangram) cleanInput() {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	pangram.data = strings.ToLower(pangram.sentence)
	pangram.data = reg.ReplaceAllString(pangram.data, "")
}

// Check the data if it results beeing a perfect pangram,
// a normal pangram or no pangram
func (pangram *pangram) checkPangram() {

	// Compare every character in the data against the map
	// of present characters. If it matches a character, count one up to that char.
	for _, c := range pangram.data {
		for k := range pangram.presentChars {
			if string(c) == string(k) {
				pangram.presentChars[k]++
			}
		}

	}

	// Count the unique chars and the chars which are there multiple times.
	// If one of the chars results 0, so not beeing there,
	// it cannot be a pangram so we immediatly set it to false.
	for _, i := range pangram.presentChars {
		if i > 1 {
			pangram.presentCharN++
		} else if i == 1 {
			pangram.presentUniqeCharN++
		} else if i == 0 {
			pangram.pangram = false
			pangram.perfectPangram = false
		}
	}

	// if the unique chars and multiple chars equal 26, set pangram to true
	if (pangram.presentCharN + pangram.presentUniqeCharN) == 26 {
		pangram.pangram = true
	}

	// if the number of truly unique characters is exactly 26,
	// it's a perfect pangram
	if pangram.presentUniqeCharN == 26 {
		pangram.perfectPangram = true
	}
}

// Check which characters are missing
func (pangram *pangram) checkMissing() {
	g := 0

	// input the number of characters to the presentChars map
	for _, d := range pangram.data {
		for i := range pangram.presentChars {
			if string(d) == string(i) {
				pangram.presentChars[i]++
				break
			}
		}
	}

	// Where there is 0 in the presentChar map, add that entry into the missing
	// char array
	for i, v := range pangram.presentChars {
		if v == 0 {
			pangram.missing[g] = i
			g++
		}
	}
}
