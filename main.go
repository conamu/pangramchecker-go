package main

import "fmt"

func pangramchecker(pangram *pangram) {
	pangram.initPangramData()
	mainMenu()
	pangram.sentence = getIO()
	pangram.cleanInput()
	switch lenData := len(pangram.data); {
	case lenData < 24:
		pangram.checkMissing()
		fmt.Println(pangram.missing)
		endMenu(pangram)
	case lenData >= 24:
		pangram.checkPangram()
		pangram.checkMissing()
		endMenu(pangram)
		pangramchecker(pangram)
	}
}

func main() {
	var p pangram
	var pp = &p
	pangramchecker(pp)
}
