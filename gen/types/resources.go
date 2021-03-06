// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/gotidy/fhir-client/gen/types/models"
)

type Resource struct {
	ResourceType string
	URL          *string
	Version      *string
	Name         *string
}

const packageName = "models"

func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}

type ResourceMap = map[string]map[string][]byte

var licenseComment = strings.Split(strings.Trim(`
Copyright 2021 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
`, "\n"), "\n")

var namePattern = regexp.MustCompile("^[A-Z]([A-Za-z0-9_]){0,254}$")

type Generator struct {
	config Config

	definitions []string
}

func NewGenerator(config Config) *Generator {
	return &Generator{config: config}
}

func (g *Generator) addDefinition(name string) {
	g.definitions = append(g.definitions, name)
}

func (g *Generator) Run() {
	resources := make(ResourceMap)
	resources["StructureDefinition"] = make(map[string][]byte)
	resources["ValueSet"] = make(map[string][]byte)
	resources["CodeSystem"] = make(map[string][]byte)

	err := filepath.Walk(g.config.Input, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".json") {
			return nil
		}
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		fmt.Printf("Generate Go sources from file: %s\n", path)
		resource, err := UnmarshalResource(bytes)
		if err != nil {
			return err
		}
		if resource.ResourceType == "Bundle" {
			bundle, err := models.UnmarshalBundle(bytes)
			if err != nil {
				return err
			}
			for _, entry := range bundle.Entry {
				entryResource, err := UnmarshalResource(entry.Resource)
				if err != nil {
					return err
				}
				switch entryResource.ResourceType {
				case "StructureDefinition":
					if entryResource.Name != nil {
						resources[entryResource.ResourceType][*entryResource.Name] = entry.Resource
					}
				case "ValueSet":
					fallthrough
				case "CodeSystem":
					if entryResource.URL != nil {
						if entryResource.Version != nil {
							resources[entryResource.ResourceType][*entryResource.URL+"|"+*entryResource.Version] = entry.Resource
							resources[entryResource.ResourceType][*entryResource.URL] = entry.Resource
						} else {
							resources[entryResource.ResourceType][*entryResource.URL] = entry.Resource
						}
					}
				}
			}
		}
		switch resource.ResourceType {
		case "StructureDefinition":
			if resource.Name != nil {
				resources[resource.ResourceType][*resource.Name] = bytes
			}
		case "ValueSet":
			fallthrough
		case "CodeSystem":
			if resource.URL != nil {
				if resource.Version != nil {
					resources[resource.ResourceType][*resource.URL+"|"+*resource.Version] = bytes
					resources[resource.ResourceType][*resource.URL] = bytes
				} else {
					resources[resource.ResourceType][*resource.URL] = bytes
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	requiredTypes := make(map[string]bool, 0)
	requiredValueSetBindings := make(map[string]bool, 0)

	for _, bytes := range resources["StructureDefinition"] {
		structureDefinition, err := models.UnmarshalStructureDefinition(bytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if (structureDefinition.Kind == models.StructureDefinitionKindResource) &&
			structureDefinition.Name != "Element" &&
			structureDefinition.Name != "BackboneElement" {
			g.addDefinition(structureDefinition.Name)

			goFile, err := g.generateResourceOrType(resources, requiredTypes, requiredValueSetBindings, structureDefinition)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			err = goFile.Save(filepath.Join(g.config.Output, FirstLower(structureDefinition.Name)+".gen.go"))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	err = g.generateTypes(resources, make(map[string]bool, 0), requiredTypes, requiredValueSetBindings)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for url := range requiredValueSetBindings {
		bytes := resources["ValueSet"][url]
		if bytes == nil {
			fmt.Printf("Missing ValueSet `%s`.\n", url)
			os.Exit(1)
		}
		valueSet, err := models.UnmarshalValueSet(bytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if valueSet.Name == nil {
			fmt.Println("Skip ValueSet without name.")
			continue
		}
		if !namePattern.MatchString(*valueSet.Name) {
			fmt.Printf("Skip ValueSet with non-conforming name `%s`.\n", *valueSet.Name)
			continue
		}
		goFile, err := g.generateValueSet(resources, valueSet)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = goFile.Save(filepath.Join(g.config.Output, FirstLower(*valueSet.Name)+".gen.go"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	g.generateDefinitionsList()
}

func (g *Generator) generateDefinitionsList() {
	fmt.Println("Definitions list saving", g.config.Definitions, len(g.definitions))
	if len(g.definitions) == 0 || g.config.Definitions == "" {
		return
	}

	file := jen.NewFile("main")
	g.setHeader(file)

	var defs = make([]jen.Code, 0, len(g.definitions))
	sort.Strings(g.definitions)
	for _, def := range g.definitions {
		defs = append(defs, jen.Lit(def))
	}
	multiLine := jen.Options{
		Close:     "}",
		Multi:     true,
		Open:      "{",
		Separator: ",",
	}
	file.Var().Id("Definitions").Op("=").Index().String().Custom(multiLine, defs...)

	if err := file.Save(g.config.Definitions); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Definitions list saved to:", g.config.Definitions)
}

func FirstLower(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func (g *Generator) generateTypes(resources ResourceMap, alreadyGeneratedTypes map[string]bool, types map[string]bool, requiredValueSetBindings map[string]bool) error {
	moreRequiredTypes := make(map[string]bool, 0)
	for name := range types {
		bytes := resources["StructureDefinition"][name]
		if bytes == nil {
			return fmt.Errorf("missing StructureDefinition with name `%s`", name)
		}
		structureDefinition, err := models.UnmarshalStructureDefinition(bytes)
		if err != nil {
			return err
		}
		goFile, err := g.generateResourceOrType(resources, moreRequiredTypes, requiredValueSetBindings, structureDefinition)
		if err != nil {
			return err
		}
		err = goFile.Save(filepath.Join(g.config.Output, FirstLower(structureDefinition.Name)+".gen.go"))
		if err != nil {
			return err
		}
		alreadyGeneratedTypes[name] = true
	}
	for name := range alreadyGeneratedTypes {
		delete(moreRequiredTypes, name)
	}
	if len(moreRequiredTypes) > 0 {
		return g.generateTypes(resources, alreadyGeneratedTypes, moreRequiredTypes, requiredValueSetBindings)
	}
	return nil
}

func (g *Generator) generateResourceOrType(resources ResourceMap, requiredTypes map[string]bool, requiredValueSetBindings map[string]bool, definition models.StructureDefinition) (*jen.File, error) {
	elementDefinitions := definition.Snapshot.Element
	if len(elementDefinitions) == 0 {
		return nil, fmt.Errorf("missing element definitions in structure definition `%s`", definition.Name)
	}

	fmt.Printf("Generate Go sources for StructureDefinition: %s\n", definition.Name)
	file := jen.NewFile(packageName)
	g.setHeader(file)

	// generate structs
	file.Commentf("%s is documented here %s", definition.Name, definition.URL)
	var err error
	file.Type().Id(definition.Name).StructFunc(func(rootStruct *jen.Group) {
		_, err = g.appendFields(resources, requiredTypes, requiredValueSetBindings, file, rootStruct, definition.Name, elementDefinitions, 1, 1)
	})
	if err != nil {
		return nil, err
	}

	// generate marshal
	if definition.Kind == models.StructureDefinitionKindResource {
		file.Type().Id("Other" + definition.Name).Id(definition.Name)
		file.Commentf("MarshalJSON marshals the given %s as JSON into a byte slice", definition.Name)
		file.Func().Params(jen.Id("r").Id(definition.Name)).Id("MarshalJSON").Params().
			Params(jen.Op("[]").Byte(), jen.Error()).Block(
			jen.Return().Qual("encoding/json", "Marshal").Call(jen.Struct(
				jen.Id("Other"+definition.Name),
				jen.Id("ResourceType").String().Tag(map[string]string{"json": "resourceType"}),
			).Values(jen.Dict{
				jen.Id("Other" + definition.Name): jen.Id("Other" + definition.Name).Call(jen.Id("r")),
				jen.Id("ResourceType"):            jen.Lit(definition.Name),
			})),
		)
	}

	// generate unmarshal
	if definition.Kind == models.StructureDefinitionKindResource {
		file.Commentf("Unmarshal%s unmarshals a %s.", definition.Name, definition.Name)
		file.Func().Id("Unmarshal"+definition.Name).
			Params(jen.Id("b").Op("[]").Byte()).
			Params(jen.Id(definition.Name), jen.Error()).
			Block(
				jen.Var().Id(FirstLower(definition.Name)).Id(definition.Name),
				jen.If(
					jen.Err().Op(":=").Qual("encoding/json", "Unmarshal").Call(
						jen.Id("b"),
						jen.Op("&").Id(FirstLower(definition.Name)),
					),
					jen.Err().Op("!=").Nil(),
				).Block(
					jen.Return(jen.Id(FirstLower(definition.Name)), jen.Err()),
				),
				jen.Return(jen.Id(FirstLower(definition.Name)), jen.Nil()),
			)
	}

	return file, nil
}

func (g *Generator) setLicenseComment(file *jen.File) {
	for _, line := range licenseComment {
		file.HeaderComment(line)
	}
}

func (g *Generator) setGeneratorComment(file *jen.File) {
	file.HeaderComment("Code generated by go generate; DO NOT EDIT.")
	file.HeaderComment("This file was generated by robots at")
	file.HeaderComment(time.Now().UTC().String())
}

func (g *Generator) setHeader(file *jen.File) {
	g.setLicenseComment(file)
	file.HeaderComment("")
	g.setGeneratorComment(file)
}

func (g *Generator) appendFields(resources ResourceMap, requiredTypes map[string]bool, requiredValueSetBindings map[string]bool, file *jen.File, fields *jen.Group, parentName string, elementDefinitions []models.ElementDefinition, start, level int) (int, error) {
	//fmt.Printf("appendFields parentName=%s, start=%d, level=%d\n", parentName, start, level)
	for i := start; i < len(elementDefinitions); i++ {
		element := elementDefinitions[i]
		pathParts := strings.Split(element.Path, ".")
		if len(pathParts) == level+1 {
			// direct childs
			name := normalizeName(pathParts[level])

			// support contained resources later
			if name != "Contained" {
				switch len(element.Type) {
				case 0:
					if element.ContentReference != nil && (*element.ContentReference)[:1] == "#" {
						statement := fields.Id(name)

						if *element.Max == "*" {
							statement.Op("[]")
						} else if *element.Min == 0 {
							statement.Op("*")
						}

						typeIdentifier := ""
						for _, pathPart := range strings.Split((*element.ContentReference)[1:], ".") {
							typeIdentifier += strings.Title(pathPart)
						}
						statement.Id(typeIdentifier).Tag(map[string]string{"json": pathParts[level] + ",omitempty", "bson": pathParts[level] + ",omitempty"})
					}
				// support polymorphic elements later
				case 1:
					statement := fields.Id(name)

					switch element.Type[0].Code {
					case "code":
						if *element.Max == "*" {
							statement.Op("[]")
						} else if *element.Min == 0 {
							statement.Op("*")
						}

						if url := g.requiredValueSetBinding(element); url != nil {
							if bytes := resources["ValueSet"][*url]; bytes != nil {
								valueSet, err := models.UnmarshalValueSet(bytes)
								if err != nil {
									return 0, err
								}
								if name := valueSet.Name; name != nil {
									if !namePattern.MatchString(*name) {
										fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet has a non-conforming name.\n", *name)
										statement.Id("string")
									} else if len(valueSet.Compose.Include) > 1 {
										fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet includes more than one CodeSystem.\n", *valueSet.Name)
										statement.Id("string")
									} else if codeSystemURL := canonical(valueSet.Compose.Include[0]); resources["CodeSystem"][codeSystemURL] == nil {
										fmt.Printf("Skip generating an enum for a ValueSet binding to `%s` because the ValueSet includes the non-existing CodeSystem with canonical URL `%s`.\n", *valueSet.Name, codeSystemURL)
										statement.Id("string")
									} else {
										requiredValueSetBindings[*url] = true
										statement.Id(*name)
									}
								} else {
									return 0, fmt.Errorf("missing name in ValueSet with canonical URL `%s`", *url)
								}
							} else {
								statement.Id("string")
							}
						} else {
							statement.Id("string")
						}
					case "Resource":
						statement.Qual("encoding/json", "RawMessage")
					default:
						if *element.Max == "*" {
							statement.Op("[]")
						} else if *element.Min == 0 {
							statement.Op("*")
						}

						var typeIdentifier string
						if parentName == "Element" && name == "ID" ||
							parentName == "Extension" && name == "URL" {
							typeIdentifier = "string"
						} else {
							typeIdentifier = typeCodeToTypeIdentifier(element.Type[0].Code)
						}
						if typeIdentifier == "Element" || typeIdentifier == "BackboneElement" {
							backboneElementName := parentName + name
							statement.Id(backboneElementName)
							var err error
							file.Type().Id(backboneElementName).StructFunc(func(childFields *jen.Group) {
								//var err error
								i, err = g.appendFields(resources, requiredTypes, requiredValueSetBindings, file, childFields, backboneElementName, elementDefinitions, i+1, level+1)
							})
							if err != nil {
								return 0, err
							}
							i--
						} else {
							if unicode.IsUpper(rune(typeIdentifier[0])) && !isPredefinedType(typeIdentifier) {
								requiredTypes[typeIdentifier] = true
							}
							statement.Id(typeIdentifier)
						}
					}

					if *element.Min == 0 {
						statement.Tag(map[string]string{"json": pathParts[level] + ",omitempty", "bson": pathParts[level] + ",omitempty"})
					} else {
						statement.Tag(map[string]string{"json": pathParts[level], "bson": pathParts[level]})
					}
				}
			}
		} else {
			// index of the next parent sibling
			return i, nil
		}
	}
	return 0, nil
}

func (g *Generator) requiredValueSetBinding(elementDefinition models.ElementDefinition) *string {
	if elementDefinition.Binding != nil {
		binding := *elementDefinition.Binding
		if binding.Strength == models.BindingStrengthRequired {
			return binding.ValueSet
		}
	}
	return nil
}

var nameMapping = map[string]string{
	"id":  "ID",
	"url": "URL",
	"uri": "URI",
}

func normalizeName(name string) string {
	if name, ok := nameMapping[name]; ok {
		return name
	}
	return strings.Title(name)
}

func isPredefinedType(t string) bool {
	return t == "DateTime" || t == "Time" || t == "Decimal"
}

func typeCodeToTypeIdentifier(typeCode string) string {
	switch typeCode {
	case "base64Binary":
		return "string"
	case "boolean":
		return "bool"
	case "canonical":
		return "string"
	case "code":
		return "string"
	case "date":
		return "DateTime" // "string"
	case "dateTime":
		return "DateTime" // "string"
	case "decimal":
		return "Decimal"
	case "id":
		return "string"
	case "instant":
		return "string"
	case "integer":
		return "int"
	case "markdown":
		return "string"
	case "oid":
		return "string"
	case "positiveInt":
		return "int"
	case "string":
		return "string"
	case "time":
		return "Time" // "string"
	case "unsignedInt":
		return "int"
	case "uri":
		return "string"
	case "url":
		return "string"
	case "uuid":
		return "string"
	case "xhtml":
		return "string"
	case "http://hl7.org/fhirpath/System.String":
		return "string"
	default:
		return typeCode
	}
}
