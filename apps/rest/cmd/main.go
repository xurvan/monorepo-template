package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := fmt.Fprintf(w, `{"message": "hello"}`)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	log.Println("Serving HTTP on 0.0.0.0:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
