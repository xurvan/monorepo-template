package main

import (
	"time"

	"github.com/xurvan/monorepo-template/tools/hello"
)

func main() {
	for {
		hello.LogHello()
		time.Sleep(5 * time.Second)
	}
}
