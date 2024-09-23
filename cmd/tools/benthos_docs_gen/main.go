package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/benthosdev/benthos/v4/internal/api"
	tdocs "github.com/benthosdev/benthos/v4/internal/cli/test/docs"
	"github.com/benthosdev/benthos/v4/internal/docs"
	"github.com/benthosdev/benthos/v4/internal/log"
	"github.com/benthosdev/benthos/v4/internal/template"
	"github.com/benthosdev/benthos/v4/public/service"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
)

func create(t, path string, resBytes []byte) {
	if existing, err := os.ReadFile(path); err == nil {
		if bytes.Equal(existing, resBytes) {
			return
		}
	}
	if err := os.WriteFile(path, resBytes, 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Documentation for '%v' has changed, updating: %v\n", t, path)
}

func main() {
	docsDir := "./website/docs/components"
	flag.StringVar(&docsDir, "dir", docsDir, "The directory to write docs to")
	flag.Parse()

	service.GlobalEnvironment().WalkInputs()
	service.GlobalEnvironment().WalkOutputs()
}

func doBloblang(dir string) {
	mdSpec, err := docs.BloblangFunctionsMarkdown()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for bloblang functions: %v", err))
	}

	create("bloblang functions", filepath.Join(dir, "..", "guides", "bloblang", "functions.md"), mdSpec)

	if mdSpec, err = docs.BloblangMethodsMarkdown(); err != nil {
		panic(fmt.Sprintf("Failed to generate docs for bloblang methods: %v", err))
	}

	create("bloblang methods", filepath.Join(dir, "..", "guides", "bloblang", "methods.md"), mdSpec)
}

func doTestDocs(dir string) {
	mdSpec, err := tdocs.DocsMarkdown()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for unit tests: %v", err))
	}

	create("test docs", filepath.Join(dir, "..", "configuration", "unit_testing.md"), mdSpec)
}

func doHTTP(dir string) {
	mdSpec, err := api.DocsMarkdown()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for http: %v", err))
	}

	create("http docs", filepath.Join(dir, "http", "about.md"), mdSpec)
}

func doLogger(dir string) {
	mdSpec, err := log.DocsMarkdown()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for logger: %v", err))
	}

	create("logger docs", filepath.Join(dir, "logger", "about.md"), mdSpec)
}

func doTemplates(dir string) {
	mdSpec, err := template.DocsMarkdown()
	if err != nil {
		panic(fmt.Sprintf("Failed to generate docs for templates: %v", err))
	}

	create("template docs", filepath.Join(dir, "..", "configuration", "templating.md"), mdSpec)
}
