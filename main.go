package main

import (
	"boot/internal/bar"
	"boot/internal/foo"
)

func main() {
	bar := &bar.bar{}
	foo := foo.NewFoo(bar)

	foo.Greet()
}
