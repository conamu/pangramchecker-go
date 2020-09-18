package main

type customMap map[string]int

type pangram struct {
	sentence       string
	missing        []string
	pangram        bool
	perfectPangram bool
	data           string
	presentChars   customMap
}
