package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/rancher/go-rancher/client"
	"gopkg.in/yaml.v2"
)

var (
	descRegexp    *regexp.Regexp = regexp.MustCompile(`<resource>`)
	optionsRegexp *regexp.Regexp = regexp.MustCompile(`<list of options>`)
)

type ResourceDescription struct {
	Description     string            `yaml:"description"`
	ResourceFields  map[string]string `yaml:"resourceFields,omitempty"`
	ResourceActions map[string]string `yaml:"resourceActions,omitempty"`
}

type FieldDescription struct {
	Description string `yaml:"description"`
}

type ActionDescription struct {
	Description string `yaml:"description"`
}

func getOptions(options []string) string {
	return strings.Join(options, ", ")
}

func generateDescriptionFile() error {

	schemaBytes, err := ioutil.ReadFile("./input/schemas.json")
	if err != nil {
		return err
	}

	var schemas client.Schemas

	err = json.Unmarshal(schemaBytes, &schemas)
	if err != nil {
		return err
	}

	genDescMap := make(map[string]string)
	err = readGenDescFile(genDescMap)
	if err != nil {
		return err
	}

	descriptionsMap = make(map[string]ResourceDescription)

	for _, schema := range schemas.Data {
		if _, ok := blackListTypes[schema.Id]; ok {
			continue
		}

		//add to descriptionsMap
		rd := ResourceDescription{}
		rd.ResourceActions = make(map[string]string)
		rd.ResourceFields = make(map[string]string)

		for actionName, _ := range schema.ResourceActions {
			rd.ResourceActions[actionName] = "To " + actionName + " the " + schema.Id
		}

		for fieldName, field := range schema.ResourceFields {
			//check if a generic desc exists
			genDesc, ok := genDescMap[fieldName]
			var description string
			if ok {
				description = descRegexp.ReplaceAllString(genDesc, schema.Id)
				description = optionsRegexp.ReplaceAllString(description, "["+getOptions(field.Options)+"]")
			} /*else {
				//description = "The " + fieldName + " for the " + schema.Id
			}*/
			rd.ResourceFields[fieldName] = description
		}

		descriptionsMap[schema.Id] = rd
	}

	err = setupDirectory(API_OUTPUT_DIR)
	if err != nil {
		return err
	}

	output, err := os.Create(path.Join(API_INPUT_DIR, "api_description.yml"))

	if err != nil {
		return err
	}

	defer output.Close()

	data := map[string]interface{}{
		"descriptionMap": descriptionsMap,
	}

	templateName := "apiDescription.template"

	typeTemplate, err := template.New(templateName).ParseFiles("./templates/" + templateName)
	if err != nil {
		return err
	}

	return typeTemplate.Execute(output, data)
}

func readGenDescFile(structure map[string]string) error {
	//read yaml file to load the common desc
	composeBytes, err := ioutil.ReadFile("./input/generic_descriptions.yml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(composeBytes, &structure)
}
