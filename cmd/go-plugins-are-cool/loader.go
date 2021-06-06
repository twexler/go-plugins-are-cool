package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"plugin"
	"strings"

	factory "github.com/twexler/go-plugins-are-cool/pkg/hello"
)

type factories map[string]factory.HelloFactory

func NewPluginLoader() PluginLoader {
	return &loader{
		errors: make([]error, 0),
	}
}

type PluginLoader interface {
	Load(pluginDir string) factories
	Errors() []error
}

func (fs factories) merge(other factories) {
	for k, v := range other {
		fs[k] = v
	}
}

type loader struct {
	errors []error
}

func (l *loader) Load(pluginDir string) factories {
	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		panic(err)
	}
	fs := make(factories)
	for _, entry := range entries {
		path := path.Join(pluginDir, entry.Name())
		if entry.IsDir() {
			fs.merge(l.Load(path))
		} else {
			if strings.TrimSuffix(path, ".so") != path {
				// this is a plugin!
				fmt.Println("loading", path)
				p, err := plugin.Open(path)
				if err != nil {
					l.errors = append(l.errors, err)
					continue
				}
				nameSym, err := p.Lookup("PluginName")
				if err != nil {
					l.errors = append(l.errors, err)
					continue
				}
				name, ok := nameSym.(*string)
				if !ok {
					l.errors = append(l.errors, errors.New("plugin name is not a string"))
				}
				factorySym, err := p.Lookup("Factory")
				if err != nil {
					l.errors = append(l.errors, err)
					continue
				}
				factory, ok := factorySym.(*factory.HelloFactory)
				if !ok {
					l.errors = append(l.errors, errors.New("does not implement factory type"))
					continue
				}
				fs[*name] = *factory
			}
		}
	}
	return fs
}

func (l *loader) Errors() []error {
	return l.errors
}
