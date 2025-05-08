package main

import (
	"flag"

	"github.com/brookemitchell/google-go-style/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	// Don't use it: just to not crash on -unsafeptr flag from go vet
	flag.Bool("unsafeptr", false, "")

	// Single checker for Google Go Style Guide rules
	// - No snake_case names (use camelCase or PascalCase)
	singlechecker.Main(analyzer.Analyzer)
}
