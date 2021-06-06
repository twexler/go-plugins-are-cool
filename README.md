# go-plugins-are-cool

This is an experimental package utilizing the Go [`plugin`](https://pkg.go.dev/plugin) package. It contains a simple plugin architecture, and a code generator to automatically load all plugins of a certain type.

## Generating a plugin loader

`cmd/go-plugins-are-cool/loader.go` is an example of a generated loader

```
$ go run ./cmd/loadergen -help
Usage of loadergen:
  -factory-function-type-name string
        the name of the factory function type name
  -factory-import-path string
        the import path where the factory function lives
  -force
        force creation if file already exists
  -output string
        the path to the newly generated loader code (default "plugin_loader.go")
  -package-name string
```