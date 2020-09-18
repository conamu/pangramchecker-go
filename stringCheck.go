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

}

func (pangram *pangram) checkPerfectPangram() {

}
