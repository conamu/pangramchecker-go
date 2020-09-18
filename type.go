package main

type customMap map[string]bool

type pangram struct {
	sentence       string
	missing        []string
	pangram        bool
	perfectPangram bool
	data           string
	presentChars   customMap
}
