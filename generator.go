package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/rancherio/go-rancher/client"
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

		if strings.HasPrefix(field.Type, "reference") || strings.HasPrefix(field.Type, "date") || strings.HasPrefix(field.Type, "enum") {
			result[fieldName] = "string"
		} else if strings.HasPrefix(field.Type, "array") {
			if strings.Contains(field.Type, "reference") {
				result[fieldName] = "[]string"
			} else if strings.Contains(field.Type, "resource") {
				result[fieldName] = "[]interface{}"
			}
		} else if strings.HasPrefix(field.Type, "map") {
			result[fieldName] = "map[string]interface{}"
		} else if strings.HasPrefix(field.Type, "json") {
			result[fieldName] = "interface{}"
		} else if strings.HasPrefix(field.Type, "boolean") {
			result[fieldName] = "bool"
		} else if strings.HasPrefix(field.Type, "extensionPoint") {
			result[fieldName] = "interface{}"
		} else {
			result[fieldName] = field.Type
		}
	}

	return result
}

func generateSchema(schema client.Schema) error {
	output, err := os.Create("client/" + schema.Id + ".go")

	if err != nil {
		return err
	}

	defer output.Close()

	data := map[string]interface{}{
		"schema":          schema,
		"typeCapitalized": strings.ToUpper(schema.Id[:1]) + schema.Id[1:],
		"structFields":    getTypeMap(schema),
		"resourceActions": schema.ResourceActions,
	}

	funcMap := template.FuncMap{
		"toLowerCamelCase": ToLowerCamelCase,
		"toUpperCamelCase": ToUpperCamelCase,
	}

	templ := template.New("type.template")

	if err != nil {
		return err
	}

	return template.Must(templ.Funcs(funcMap).ParseFiles("type.template")).Execute(output, data)

}

func ToLowerCamelCase(input string) string {
	return (strings.ToLower(input[:1]) + input[1:])
}

func ToUpperCamelCase(input string) string {
	return (strings.ToUpper(input[:1]) + input[1:])
}

func generateClient(schema []client.Schema) error {
	template, err := template.ParseFiles("client.template")
	if err != nil {
		return err
	}

	output, err := os.Create("client/clientgenerator.go")
	if err != nil {
		return err
	}

	defer output.Close()
	buffer := make([]client.Schema, len(schema)-3)
	i := 0
	for _, val := range schema {

		if !(val.Id == "collection" || val.Id == "resource" || val.Id == "schema") {
			val.Id = strings.ToUpper(val.Id[:1]) + val.Id[1:]
			buffer[i] = val
			i++
		}
	}

	result := map[string]interface{}{
		"schemas": buffer,
	}

	err = template.Execute(output, result)
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
		if schema.Id == "collection" || schema.Id == "resource" || schema.Id == "schema" {
			continue
		}

		err = generateSchema(schema)
		if err != nil {
			return err
		}
	}
	err = generateClient(schemas.Data)

	return nil
}

func main() {

	err := generateFiles()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
