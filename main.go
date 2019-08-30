package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", r.URL.Path[1:])
	log.Println(r.RemoteAddr, r.Method, r.URL)
}
func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
