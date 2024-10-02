package foo

type bar interface {
	SayHello()
}

type foo struct {
	bar bar
}
