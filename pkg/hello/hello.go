package hello

type Hello interface {
	SayHello()
}

type HelloFactory func() Hello
