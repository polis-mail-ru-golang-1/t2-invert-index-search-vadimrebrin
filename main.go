package main

import (
	"bufio"
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
	"net/http"
	"os"
	"strings"
)

var dict map[string]map[string]int

func main() {
	dict = make(map[string]map[string]int)
	files := os.Args[1:]

	http.HandleFunc("/", handleConnection)
	index.BuildIndex(dict, readFiles(files))
	fmt.Println("Starting server at :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	name := r.RemoteAddr
	fmt.Println(name + " connected")
	text := r.FormValue("q")
	if text != "" {
		text = strings.ToLower(text)
		fmt.Println(name + " entered " + text)
		phrase := strings.Fields(text)
		fmt.Fprintln(w, index.FindPhrase(dict, phrase))
	}
}

//readLines reads file and converts it to slice of lines
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

//readFiles reads files and returns map of these files
func readFiles(args []string) map[string][]string {
	files := make(map[string][]string)
	for _, file := range args {
		files[file] = readLines(file)
	}
	return files
}

//localhost/?q=good+morning
