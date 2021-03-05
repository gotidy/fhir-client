package main

import (
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/client.go.tmpl
var clientTemplate string

//go:embed templates/types.go.tmpl
var typesTemplate string

type Generator struct {
	config Config
}

func NewGenerator(config Config) *Generator {
	return &Generator{config}
}

type Data struct {
	Entities []string
}

func (g *Generator) Run() error {
	tmpl := template.Must(template.New("client").Parse(clientTemplate))

	file, err := os.Create(filepath.Join(g.config.Output, "client.gen.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, Data{Entities: Definitions}); err != nil {
		return err
	}

	tmpl = template.Must(template.New("client").Parse(typesTemplate))

	file, err = os.Create(filepath.Join(g.config.Output, "types.gen.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, Data{Entities: Definitions})

	// for _, def := range Definitions {
	// 	err := func() error {
	// 		file, err := os.Create(filepath.Join(g.config.Output, def+".go"))
	// 		if err != nil {
	// 			return err
	// 		}
	// 		defer file.Close()

	// 		return tmpl.Execute(file, Data{Entity: def})
	// 	}()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}
