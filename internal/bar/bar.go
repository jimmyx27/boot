package bar

import (
	"fmt"
)

type Bar struct {
}

func (b *Bar) SayHello() {
	fmt.Println("hello")
}
func (b *Bar) SayHola() {
	fmt.Println("hola")
}
func (b *Bar) SayPrivet() {
	fmt.Println("privet")
}
func (b *Bar) SayGoddog() {
	fmt.Println("goddog")
}
