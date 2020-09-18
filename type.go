package main

type customMap map[string]int

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
