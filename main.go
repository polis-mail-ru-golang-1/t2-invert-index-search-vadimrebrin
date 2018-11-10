package main

import (
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
	index.BuildIndex(dict, files)
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

//localhost/?q=good+morning
