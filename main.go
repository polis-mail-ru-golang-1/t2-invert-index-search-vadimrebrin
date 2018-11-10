package main

import (
	"encoding/json"
	"fmt"
	"github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
	"log"
	"net/http"
	"os"
	"strings"
	"bufio"
	"text/template"
	"time"
)

type configuration struct {
	Files []string
}

var dict map[string]map[string]int
var start time.Time
var l log.Logger

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		zl.Debug().
			Str("method", r.Method).
			Str("remote", r.RemoteAddr).
			Str("path", r.URL.Path).
			Int("duration", int(time.Since(start))).
			Msgf("Called url %s", r.URL.Path)
	})
}

func main() {
	//configuration
	conf, _ := os.Open("conf.json")
	defer conf.Close()
	decoder := json.NewDecoder(conf)
	configuration := configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}

	//logging
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zl.Logger = zl.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	//index initialization 
	dict = make(map[string]map[string]int)
	index.BuildIndex(dict, readFiles(configuration.Files))
	
	//starting server
	zl.Printf("Staring server at :80")
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/search", searchHandler)
	siteMux.HandleFunc("/", staticHandler)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)

	siteMux.Handle("/data/", staticHandler)
	siteHandler := logMiddleware(siteMux)
	http.ListenAndServe(":80", siteHandler)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	tml := template.Must(template.ParseFiles("static/index.html"))
	tml.Execute(w, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	zl.Printf(r.RemoteAddr + " searches " + r.FormValue("q"))
	if r.FormValue("q") != "" {
		text := r.FormValue("q")
		var phrase []string

		for _, str := range strings.Fields(text) {
			str = strings.ToLower(str)
			phrase = append(phrase, str)
		}

		response := index.FindPhrase(dict, phrase)
		tml := template.Must(template.ParseFiles("static/search.html"))
		tml.Execute(w, response)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
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