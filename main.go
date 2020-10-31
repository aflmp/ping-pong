package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}

func main() {
	http.HandleFunc("/ping", pong)
	log.Printf("server listening on port: %v\n", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal(err)
}

func pong(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v - %v - %v - %v\n", r.Proto, r.Method, r.URL.Path, r.UserAgent())
	fmt.Fprintf(w, "pong - %v\n", time.Now().Format(time.ANSIC))
}
