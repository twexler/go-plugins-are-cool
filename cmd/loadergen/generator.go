package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagForce               = flag.Bool("force", false, "force creation if file already exists")
	flagOutput              = flag.String("output", "plugin_loader.go", "the path to the newly generated loader code")
	flagFactoryImport       = flag.String("factory-import-path", "", "the import path where the factory function lives")
	flagPackageName         = flag.String("package-name", "loader", "the package name of the generated loader (default: loader)")
	flagFactoryFuncTypeName = flag.String("factory-function-type-name", "", "the name of the factory function type name")
)

func main() {
	flag.Parse()
	if *flagFactoryImport == "" {
		fmt.Println("missing factory import path")
		os.Exit(1)
	}
	if *flagFactoryFuncTypeName == "" {
		fmt.Println("missing factory function type name")
		os.Exit(1)
	}
	if _, err := os.Stat(*flagOutput); err == nil && !*flagForce {
		fmt.Printf("%s already exists, and -force not set\n", *flagOutput)
		os.Exit(1)
	}
	f, err := os.Create(*flagOutput)
	if err != nil {
		fmt.Printf("unable to create %s: %s", *flagOutput, err.Error())
		os.Exit(1)
	}
	if err := pluginLoaderTemplate.Execute(f, struct {
		PackageName       string
		FactoryImportPath string
		FactoryType       string
	}{
		PackageName:       *flagPackageName,
		FactoryImportPath: *flagFactoryImport,
		FactoryType:       *flagFactoryFuncTypeName,
	}); err != nil {
		fmt.Printf("error evaluating template: %s", err)
		os.Exit(1)
	}
}
