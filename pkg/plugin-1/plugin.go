package main

import (
	"fmt"

	"github.com/twexler/go-plugins-are-cool/pkg/hello"
)

var (
	PluginName string             = "plugin1"
	Factory    hello.HelloFactory = NewHello
)

func NewHello() hello.Hello {
	return &PluginHello{}
}

type PluginHello struct{}

func (p PluginHello) SayHello() {
	fmt.Println("hello from the plugin")
}
