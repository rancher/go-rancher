package main

import (
	"encoding/json"
	"github.com/rancherio/go-rancher/client"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func capitalize(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func getTypeMap(schema client.Schema) map[string]string {
	result := map[string]string{}

	for name, field := range schema.ResourceFields {
		fieldName := capitalize(name)
		if strings.HasPrefix(field.Type, "reference") {
			result[fieldName] = "string"
		} else {
			result[fieldName] = field.Type
		}
	}

	return result
}

func generateSchema(schema client.Schema) error {
	template, err := template.ParseFiles("type.template")
	if err != nil {
		return err
	}

	output, err := os.Create("generated/" + schema.Id + ".go")
	if err != nil {
		return err
	}

	defer output.Close()

	data := map[string]interface{}{
		"schema":          schema,
		"typeCapitalized": strings.ToUpper(schema.Id[:1]) + schema.Id[1:],
		"structFields":    getTypeMap(schema),
	}

	err = template.Execute(output, data)

	return err
}

func generateFiles() error {
	schemaBytes, err := ioutil.ReadFile("schemas.json")
	if err != nil {
		return err
	}

	var schemas client.Schemas
	err = json.Unmarshal(schemaBytes, &schemas)
	if err != nil {
		return err
	}

	for _, schema := range schemas.Data {
		err = generateSchema(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

func generate() {
	err := generateFiles()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
