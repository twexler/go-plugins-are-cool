package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("specify a plugin directory")
	}
	loader := NewPluginLoader()
	helloFactories := loader.Load(os.Args[1])
	if len(loader.Errors()) > 0 {
		for _, err := range loader.Errors() {
			fmt.Printf("error: %s\n", err.Error())
		}
		os.Exit(1)
	}
	for _, hello := range helloFactories {
		h := hello()
		h.SayHello()
	}
}
