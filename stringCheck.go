package main

import (
	"regexp"
	"strings"
)

func (pangram *pangram) initPangramData() {
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
	pangram.missing = make([]string, 26)
}

func (pangram *pangram) cleanInput() {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	pangram.data = strings.ToLower(pangram.sentence)
	pangram.data = reg.ReplaceAllString(pangram.data, "")
}

func (pangram *pangram) checkPangram() {
	for _, c := range pangram.data {
		for k := range pangram.presentChars {
			if string(c) == string(k) {
				pangram.presentChars[k]++
			}
		}

	}

	for _, i := range pangram.presentChars {
		if i == 0 || i > 1 {
			pangram.presentCharN++
		} else if i == 1 {
			pangram.presentUniqeCharN++
		} else if i == 0 {
			pangram.pangram = false
			pangram.perfectPangram = false
		}
	}

	if (pangram.presentCharN + pangram.presentUniqeCharN) == 26 {
		pangram.pangram = true
	}

	if pangram.presentUniqeCharN == 26 {
		pangram.perfectPangram = true
	}
}

func (pangram *pangram) checkMissing() {
	g := 0
	for i, v := range pangram.presentChars {
		if v == 0 {
			pangram.missing[g] = i
			g++
		}
	}
}
