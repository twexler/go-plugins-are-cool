package main

import (
	"fmt"

	"github.com/twexler/go-plugins-are-cool/pkg/hello"
	"gopkg.in/AlecAivazis/survey.v1"
)

var (
	PluginName string             = "plugin2"
	Factory    hello.HelloFactory = NewHello
)

type myHello struct {
	Name string
}

func (h myHello) SayHello() {
	fmt.Printf("hello from %s\n", h.Name)
}

func NewHello() hello.Hello {
	h := myHello{}
	if err := survey.Ask([]*survey.Question{
		{
			Name: "Name",
			Prompt: &survey.Input{
				Message: "What's your name?",
			},
		},
	}, &h); err != nil {
		panic(err)
	}
	return &h
}
