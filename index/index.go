package index

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BuildIndex(dict map[string]map[string]int) {
	files := os.Args[1:]

	for _, onefile := range files {
		lines := readLines(onefile)
		words := linesToWords(lines)
		for _, word := range words {
			//если слово встретилось первый раз
			if dict[word] == nil {
				filemap := make(map[string]int)
				filemap[onefile]++
				dict[word] = filemap
			} else {
				dict[word][onefile]++
			}
		}
	}
}

func FindPhrase(dict map[string]map[string]int, phrase []string) map[string]int {
	samewords := make(map[string]map[string]int)
	for item, _ := range dict {
		for _, word := range phrase {
			if item == word {
				samewords[word] = dict[word]
			}
		}
	}
	res := make(map[string]int)
	for _, item := range samewords {
		for name, i := range item {
			res[name] = res[name] + i
		}
	}
	return res
}

func PrintInfo(dict map[string]int) {
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

func ReadPhrase() []string {
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
