package main

import (
	"fmt"
	index "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
)

func main() {
	dict := make(map[string]map[string]int)
	result := make(map[string]int)
	var phrase []string
	index.BuildIndex(dict)
	fmt.Println("Enter search phrase:")
	phrase = index.ReadPhrase()
	result = index.FindPhrase(dict, phrase)
	index.PrintInfo(result)
}
