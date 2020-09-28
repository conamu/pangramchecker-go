package main

type customMap map[string]int

// Custom type for pangram object
type pangram struct {
	sentence          string
	missing           []string
	presentCharN      int
	presentUniqeCharN int
	pangram           bool
	perfectPangram    bool
	data              string
	presentChars      customMap
}
