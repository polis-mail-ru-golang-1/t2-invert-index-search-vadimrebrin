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
	index.BuildIndex(dict, readFiles(files))
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
	for _, str := range strings.Fields(scanner.Text()) {
		str = strings.ToLower(str)
		phrase = append(phrase, str)
	}
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
		for _, str := range strings.Fields(scanner.Text()) {
			str = strings.Trim(str, ".,?!-\"")
			str = strings.ToLower(str)
			lines = append(lines, str)
		}
	}
	return lines
}

func readFiles(args []string) map[string][]string {
	files := make(map[string][]string)
	for _, file := range args {
		files[file] = readLines(file)
	}
	return files
}
