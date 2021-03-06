package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
	"path/filepath"
	"os"
)

var (
	listenAddr = flag.String("addr", getenvWithDefault("LISTENADDR", ":8080"), "HTTP address to listen on")
	tmpl = template.New("")
)

func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)

	if val == "" {
		val = defaultValue
	}
	return val
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}


func main() {
	flag.Parse()

	var err error
	_, err = tmpl.ParseGlob(filepath.Join(".", "templates","*.html"))
	if err != nil {
		log.Fatalf("Unable to parse templates: %v\n", err)
	}

	log.Printf("Listening on %s\n", *listenAddr)

	http.HandleFunc("/", handler)
	http.ListenAndServe(*listenAddr, nil)
}
