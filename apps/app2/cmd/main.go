package main

import (
	"time"

	"github.com/xurvan/monorepo-template/tools/foo"
)

func main() {
	for {
		foo.LogFoo()
		time.Sleep(5 * time.Second)
	}
}
