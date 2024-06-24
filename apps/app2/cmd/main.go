package main

import (
	"fmt"
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

	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		panic(err)
	}
}
