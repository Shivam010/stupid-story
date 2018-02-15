package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	if err := http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/2018") {
			r.URL.Path = r.URL.Path[5:]
			if r.URL.Path == "/stupid" {
				f, _:=ioutil.ReadFile("stupid.html")
				w.Write(f)
				return
			}
			http.FileServer(http.Dir("")).ServeHTTP(w, r)
		}
	})); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
