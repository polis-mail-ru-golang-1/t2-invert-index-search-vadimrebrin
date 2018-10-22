package main

import (
	"bufio"
	"fmt"
	index "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
	"os"
	"strings"
)

func main() {
	dict := make(map[string]map[string]int)
	result := make(map[string]int)
	files := os.Args[1:]

	var phrase []string
	index.BuildIndex(dict, formatFiles(files))
	fmt.Println("Enter search phrase:")
	phrase = readPhrase()
	result = index.FindPhrase(dict, phrase)
	printInfo(result)
}

func printInfo(dict map[string]int) {
	if len(dict) == 0 {
		fmt.Println("Phrase not found")
		return
	}
	var filearr []string
	var countarr []int
	for name, count := range dict {
		filearr = append(filearr, name)
		countarr = append(countarr, count)
	}
	for i := 0; i < len(filearr); i++ {
		for j := i; j < len(filearr); j++ {
			if countarr[i] < countarr[j] {
				tempcount := countarr[i]
				countarr[i] = countarr[j]
				countarr[j] = tempcount
				tempfile := filearr[i]
				filearr[i] = filearr[j]
				filearr[j] = tempfile
			}
		}
	}
	for i := 0; i < len(filearr); i++ {
		fmt.Printf("File %s contains %d words of requested phrase\n", filearr[i], countarr[i])
	}
}

func readPhrase() []string {
	var phrase []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "" {
		panic("Empty phrase")
	}
	phrase = append(phrase, scanner.Text())
	phrase = linesToWords(phrase)
	return phrase
}

func readLines(arg string) []string {
	file, err := os.Open(arg)
	if err != nil {
		panic("File not found")
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func linesToWords(lines []string) []string {
	var result []string
	for _, item := range lines {
		arr := strings.Split(item, " ")
		for _, item := range arr {
			result = append(result, item)
		}
	}
	return result
}

func formatFiles(args []string) map[string][]string {
	files := make(map[string][]string)
	for _, file := range args {
		files[file] = linesToWords(readLines(file))
	}
	return files
}
