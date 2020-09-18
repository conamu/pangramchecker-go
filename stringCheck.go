package main

import (
	"fmt"
	"regexp"
	"strings"
)

func (pangram *pangram) cleanInput() {
	reg, _ := regexp.Compile("[^a-zA-Z]+")
	pangram.data = strings.ToLower(pangram.sentence)
	pangram.data = reg.ReplaceAllString(pangram.data, "")
	fmt.Println(pangram.data)
}

func (pangram *pangram) checkPangram() {

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

	for _, c := range pangram.data {
		for k := range pangram.presentChars {
			if string(c) == string(k) {
				pangram.presentChars[k]++
			}
		}

	}

}

func (pangram *pangram) checkPerfectPangram() {

}
