package main

import (
	"encoding/json"
	"fmt"
	console "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/console"
	index "github.com/polis-mail-ru-golang-1/t2-invert-index-search-vadimrebrin/index"
	"github.com/rs/zerolog"
	zl "github.com/rs/zerolog/log"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

type Configuration struct {
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
	//Configuration
	conf, _ := os.Open("conf.json")
	defer conf.Close()
	decoder := json.NewDecoder(conf)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
	}

	//Logging
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zl.Logger = zl.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	dict = make(map[string]map[string]int)
	index.BuildIndex(dict, console.ReadFiles(configuration.Files))
	
	zl.Printf("Staring server at :80")
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/search", searchHandler)
	siteMux.HandleFunc("/", myHandler)

	staticHandler := http.StripPrefix(
		"/data/",
		http.FileServer(http.Dir("./static")),
	)

	siteMux.Handle("/data/", staticHandler)
	siteHandler := logMiddleware(siteMux)
	http.ListenAndServe(":80", siteHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
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

		res := index.FindPhrase(dict, phrase)
		response := console.PrintInfo(res)
		tml := template.Must(template.ParseFiles("static/search.html"))
		tml.Execute(w, response)
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
