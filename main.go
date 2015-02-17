package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, _ := ioutil.ReadFile("index.html")
		w.Write(f)
	})); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
