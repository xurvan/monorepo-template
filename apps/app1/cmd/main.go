package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Print("Hello, world!")
		time.Sleep(2 * time.Second)
	}
}
