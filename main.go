package main

import (
	"boot/internal/bar"
	"boot/internal/foo"
)

func main() {
	bar := &bar.Bar{}
	foo := foo.NewFoo(bar)

	foo.Greet()
}
