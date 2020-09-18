package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Function to get User Input
func getIO() string {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getRandom(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// Function to Decorate text Menus
func decor(allText, allMenuPoints []string, decorChar string, lineWidth int, menuWidth int) string {

	// Declare menu measurements:
	menuHeight := len(allMenuPoints) + len(allText) + (lineWidth * 2) + 1

	completeArray := make([][]string, menuHeight)

	// Prepare the Horizontal decor lines to prepend to the Text and Menu points
	decorLine := make([]string, ((lineWidth * 2) + (lineWidth)))
	for i := 0; i < (lineWidth * 2); i++ {
		decorLine[i] = decorChar
	}

	for i := (lineWidth * 2); i < (lineWidth*2)+1; i++ {
		decorLine[i] = " "
	}

	// Place Decor lines on the top
	for i := 0; i < lineWidth; i++ {
		completeArray[i] = make([]string, menuWidth)
		for j := 0; j < menuWidth; j++ {
			completeArray[i][j] = decorChar
		}
	}

	// Place Decor lines at the bottom
	for i := (len(completeArray) - lineWidth); i < menuHeight; i++ {
		completeArray[i] = make([]string, menuWidth)
		for j := 0; j < menuWidth; j++ {
			completeArray[i][j] = decorChar
		}
	}

	// Place the Text
	for i := lineWidth; i < (len(allText) + lineWidth); i++ {
		completeArray[i] = make([]string, menuWidth)
		for j := 0; j < 1; j++ {
			completeArray[i][j] = allText[i-lineWidth]
			// Prepend decorlines to Text
			completeArray[i] = append(decorLine, completeArray[i][j])
			// Create a string to dynamically fill the rest of the menu, append that.
			dynamicFiller := make([]string, (menuWidth - len(allText[i-lineWidth]) - ((lineWidth * 2) + 1)))
			dynamicFiller[0] = " "
			for k := 1; k < len(dynamicFiller); k++ {
				dynamicFiller[k] = decorChar
			}

			dynamicFillString := strings.Join(dynamicFiller, "")
			completeArray[i] = append(completeArray[i], dynamicFillString)
		}
	}

	// Place a separator line
	completeArray[(len(allText) + lineWidth)] = make([]string, menuWidth)
	for i := 0; i < menuWidth; i++ {
		completeArray[(len(allText) + lineWidth)][i] = decorChar
	}

	// Place the Menu points
	for i := (lineWidth + len(allText) + 1); i < (len(allMenuPoints) + lineWidth + len(allText) + 1); i++ {
		completeArray[i] = make([]string, menuWidth)
		for j := 0; j < 1; j++ {
			completeArray[i][j] = allMenuPoints[i-lineWidth-len(allText)-1]
			// Prepend decorlines to Menu Point Text
			completeArray[i] = append(decorLine, completeArray[i][j])
			// Create a string to dynamically fill the rest of the menu, append that.
			dynamicFiller := make([]string, (menuWidth - len(allMenuPoints[i-lineWidth-len(allText)-1]) - ((lineWidth * 2) + 1)))
			dynamicFiller[0] = " "
			for k := 1; k < len(dynamicFiller); k++ {
				dynamicFiller[k] = decorChar
			}

			dynamicFillString := strings.Join(dynamicFiller, "")
			completeArray[i] = append(completeArray[i], dynamicFillString)
		}
	}

	// Unify all the arrays to a single string to be printed
	stringArray := make([]string, menuHeight)

	for i := 0; i < len(completeArray); i++ {
		stringArray[i] = strings.Join(completeArray[i], "")
	}

	fullString := strings.Join(stringArray, "\n")

	// Return full string to print on console.
	return fullString

}
