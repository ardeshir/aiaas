package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	listenAddr = flag.String("addr", getenvWithDefault("LISTENADDR", ":8080"), "HTTP address to listen on")
)

func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)

	if val == "" {
		val = defaultValue
	}
	return val
}

func main() {
	flag.Parse()

	log.Printf("Listening on %s\n", *listenAddr)
	http.ListenAndServe(*listenAddr, nil)
}
