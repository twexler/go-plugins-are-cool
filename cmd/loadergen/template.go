package main

import "text/template"

var pluginLoaderTemplate = template.Must(template.New("plugin_loader.go").Parse(`
package {{ .PackageName }}

import (
	"errors"
	"fmt"
	"os"
	"path"
	"plugin"
	"strings"

	factory "{{ .FactoryImportPath }}"
)

type factories map[string]factory.{{ .FactoryType }}

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
				factory, ok := factorySym.(*factory.{{ .FactoryType }})
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
`))
