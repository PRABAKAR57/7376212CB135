package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "PRABAKAR M")
	})
	http.HandleFunc("/registernumber", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212cb135")
	})
	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CSBS")
	})
	http.HandleFunc("/favcolor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "BLUE")
	})

	fmt.Printf("Server running (port=8080), route: go\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
