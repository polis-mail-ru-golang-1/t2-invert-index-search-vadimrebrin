package console

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//PrintInfo prints statistics of search
func PrintInfo(dict map[string]int) string {
	if len(dict) == 0 {
		return "Phrase not found\n\r"
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
	var res string
	for i := 0; i < len(filearr); i++ {
		res += ("File " + string(filearr[i]) + " contains " + strconv.Itoa((countarr[i])) + " words of requested phrase\n\r")
	}
	return res
}

//ReadPhrase reads phrase from stdin and converts it to slice of words
func ReadPhrase() []string {
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

//ReadLines reads file and converts it to slice of lines
func ReadLines(arg string) []string {
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

//ReadFiles reads files and returns map of these files
func ReadFiles(args []string) map[string][]string {
	files := make(map[string][]string)
	for _, file := range args {
		files[file] = ReadLines(file)
	}
	return files
}
