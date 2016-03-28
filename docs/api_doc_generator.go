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

const (
	API_OUTPUT_DIR = "./output/api-resources"
	API_INPUT_DIR  = "./input"
)

var (
	blackListTypes          map[string]bool
	commonFields            []string
	spaceRegexp             *regexp.Regexp = regexp.MustCompile(`([a-z])([A-Z])`)
	refeRegexp              *regexp.Regexp = regexp.MustCompile(`reference\[([a-zA-Z]+)\]`)
	descriptionsMap         map[string]ResourceDescription
	schemaMap               map[string]client.Schema
	descriptionsOverrideMap map[string]ResourceDescription
)

func init() {
	blackListTypes = make(map[string]bool)
	blackListTypes["schema"] = true
	blackListTypes["resource"] = true
	blackListTypes["collection"] = true
	schemaMap = make(map[string]client.Schema)
}

//struct to hold resourceFields

type APIField struct {
	client.Field
	Description   string `json:"description"`
	TypeReference string
}

type ActionInput struct {
	Name      string
	FieldMap  map[string]APIField
	InputJson string
}

type APIAction struct {
	Input       ActionInput
	Output      string
	Description string `json:"description"`
	Verb        string
	ActionURL   string
}

func addSpace(input string) string {
	return strings.ToLower(spaceRegexp.ReplaceAllString(input, `${1} ${2}`))
}

func getTypeRef(input string) string {
	return refeRegexp.ReplaceAllString(input, "[$1]({{site.baseurl}}/rancher/api/api-resources/$1/)")
}

func capitalize(s string) string {
	if len(s) <= 1 {
		return strings.ToUpper(s)
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func isCommonField(fieldName string) bool {
	for _, commonFieldName := range commonFields {
		if fieldName == commonFieldName {
			return true
		}
	}
	return false
}

func getFieldMap(schema client.Schema) map[string]APIField {
	fieldMap := make(map[string]APIField)

	for fieldName, field := range schema.ResourceFields {

		if isCommonField(fieldName) {
			continue
		}

		apiField := APIField{}
		apiField.Field = field
		if refeRegexp.MatchString(apiField.Field.Type) {
			//put the link to the referenced field in the form
			//[type]({{site.baseurl}}/rancher/api/type/)
			apiField.TypeReference = getTypeRef(apiField.Field.Type)
		} else if _, ok := schemaMap[apiField.Field.Type]; ok {
			apiField.TypeReference = "[" + apiField.Field.Type + "]({{site.baseurl}}/rancher/api/api-resources/" + apiField.Field.Type + "/)"
		}
		apiField.Description = getFieldDescription(schema.Id, fieldName)
		fieldMap[fieldName] = apiField
	}

	return fieldMap
}

func getActionInput(schemaID string) ActionInput {
	actionInput := ActionInput{}
	actionInput.Name = schemaID
	actionInput.FieldMap = getFieldMap(schemaMap[schemaID])
	actionInput.InputJson = generateJsonFromFields(schemaMap[schemaID].ResourceFields)

	return actionInput
}

func generateJsonFromFields(resourceFields map[string]client.Field) string {
	j, err := json.MarshalIndent(generateFieldTypeMap(resourceFields), "\n", "\t")

	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func generateFieldTypeMap(resourceFields map[string]client.Field) map[string]interface{} {
	fieldTypeJsonMap := make(map[string]interface{})
	for fieldName, field := range resourceFields {
		fieldTypeJsonMap[fieldName] = generateTypeValue(field)

	}
	return fieldTypeJsonMap
}

func generateTypeValue(field client.Field) interface{} {
	//get default value if available
	if field.Default != nil {
		return field.Default
	}

	//basic types
	switch field.Type {
	case "string":
		return "string"
	case "int":
		return 0
	case "boolean":
		return true
	case "array[string]":
		return [...]string{"string1", "string2", "...stringN"}
	case "map[string]":
		return map[string]string{"key1": "value1", "key2": "value2", "keyN": "valueN"}
	case "password":
		return field.Type
	}

	//another resourceType
	subSchema, ok := schemaMap[field.Type]
	if ok {
		return generateFieldTypeMap(subSchema.ResourceFields)
	}

	return field.Type
}

func getActionMap(schema client.Schema, isCUD bool) map[string]APIAction {
	actionMap := make(map[string]APIAction)

	if isCUD {
		//get create
		for _, method := range schema.CollectionMethods {
			if method == "POST" {
				//add create
				apiAction := APIAction{}
				apiAction.Description = getActionDescription(schema.Id, "create")
				apiAction.Verb = "POST"
				apiAction.ActionURL = "/v1/" + schema.Id
				resourceFields := make(map[string]client.Field)

				for fieldName, field := range schema.ResourceFields {
					if field.Create {
						resourceFields[fieldName] = field
					}
				}

				apiAction.Input.InputJson = generateJsonFromFields(resourceFields)
				actionMap["Create"] = apiAction
			}
		}

		for _, method := range schema.ResourceMethods {
			if method == "PUT" {
				//add update
				apiAction := APIAction{}
				apiAction.Description = getActionDescription(schema.Id, "update")
				apiAction.Verb = "PUT"
				apiAction.ActionURL = "links.self"
				resourceFields := make(map[string]client.Field)

				for fieldName, field := range schema.ResourceFields {
					if field.Update {
						resourceFields[fieldName] = field
					}
				}

				apiAction.Input.InputJson = generateJsonFromFields(resourceFields)
				actionMap["Update"] = apiAction
			} else if method == "DELETE" {
				//add delete
				apiAction := APIAction{}
				apiAction.Description = getActionDescription(schema.Id, "delete")
				apiAction.Verb = "DELETE"
				apiAction.ActionURL = "links.self"
				actionMap["Delete"] = apiAction
			}
		}

	} else {
		for actionName, action := range schema.ResourceActions {
			if strings.EqualFold("create", actionName) || strings.EqualFold("update", actionName) || strings.EqualFold("delete", actionName) {
				continue
			}
			apiAction := APIAction{}
			apiAction.Description = getActionDescription(schema.Id, actionName)
			apiAction.Input = getActionInput(action.Input)
			apiAction.Output = action.Output
			apiAction.Verb = "POST"
			apiAction.ActionURL = "actions." + actionName
			actionMap[actionName] = apiAction
		}
	}

	return actionMap
}

func generateDoc(schema client.Schema, isResource bool) error {
	err := setupDirectory(API_OUTPUT_DIR + "/" + schema.Id)
	if err != nil {
		return err
	}

	output, err := os.Create(path.Join(API_OUTPUT_DIR, schema.Id, "index.md"))

	if err != nil {
		return err
	}

	defer output.Close()

	data := map[string]interface{}{
		"schemaId":            schema.Id,
		"resourceDescription": getResourceDescription(schema.Id),
		"fieldMap":            getFieldMap(schema),
		"CUDActions":          getActionMap(schema, true),
		"actionMap":           getActionMap(schema, false),
		"pluralName":          schema.PluralName,
	}

	var templateName string
	if isResource {
		templateName = "apiResource.template"
	} else {
		templateName = "apiActionInput.template"
	}
	typeTemplate, err := template.New(templateName).ParseFiles("./templates/" + templateName)
	if err != nil {
		return err
	}

	return typeTemplate.Execute(output, data)
}

func generateApiHomePage() error {
	output, err := os.Create(path.Join(API_OUTPUT_DIR, "index.md"))

	if err != nil {
		return err
	}

	defer output.Close()

	data := map[string]interface{}{
		"schemaMap": schemaMap,
	}

	funcMap := template.FuncMap{
		"getResourceDescription": getResourceDescription,
	}

	typeTemplate, err := template.New("apiHomePage.template").Funcs(funcMap).ParseFiles("./templates/apiHomePage.template")
	if err != nil {
		return err
	}

	return typeTemplate.Execute(output, data)
}

func setupDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}

	return nil
}

func generateFiles() error {
	err := setupDirectory(API_OUTPUT_DIR)
	if err != nil {
		return err
	}

	err = readDescFile()
	if err != nil {
		return err
	}

	err = readCommonFieldsFile()
	if err != nil {
		return err
	}

	schemaBytes, err := ioutil.ReadFile(API_INPUT_DIR + "/schemas.json")
	if err != nil {
		return err
	}

	var schemas client.Schemas

	err = json.Unmarshal(schemaBytes, &schemas)
	if err != nil {
		return err
	}

	for _, schema := range schemas.Data {
		schemaMap[schema.Id] = schema
	}

	for _, schema := range schemaMap {
		if _, ok := blackListTypes[schema.Id]; ok {
			continue
		}

		var isResource bool
		collectionLink := schema.Resource.Links["collection"]
		if collectionLink != "" {
			//a resource has a collection link present. We should only document resources
			isResource = true
		} else {
			isResource = false
		}

		err = generateDoc(schema, isResource)
		if err != nil {
			return err
		}
	}

	//generate the home page
	generateApiHomePage()

	return nil
}

func readDescFile() error {
	composeBytes, err := ioutil.ReadFile(API_INPUT_DIR + "/api_description.yml")
	if err != nil {
		return err
	}
	descriptionsMap = make(map[string]ResourceDescription)
	err = yaml.Unmarshal(composeBytes, &descriptionsMap)
	if err != nil {
		return err
	}

	composeBytes, err = ioutil.ReadFile(API_INPUT_DIR + "/api_description_override.yml")
	if err != nil {
		return err
	}

	descriptionsOverrideMap = make(map[string]ResourceDescription)

	return yaml.Unmarshal(composeBytes, &descriptionsOverrideMap)
}

func readCommonFieldsFile() error {
	//read yaml file to load the common fields
	composeBytes, err := ioutil.ReadFile(API_INPUT_DIR + "/common_fields.yml")
	if err != nil {
		return err
	}
	return yaml.Unmarshal(composeBytes, &commonFields)
}

func getResourceDescription(resourceID string) string {
	var desc string

	descStruct, ok := descriptionsOverrideMap[resourceID]

	if ok && descStruct.Description != "" {
		desc = descStruct.Description
	} else {
		descStruct, ok = descriptionsMap[resourceID]

		if ok {
			desc = descStruct.Description
		} else {
			desc = "This is the " + resourceID + " resource"
		}
	}

	return desc
}

func getFieldDescription(resourceID string, fieldID string) string {
	desc := "This is the " + fieldID + " field"

	resDescStruct, ok := descriptionsOverrideMap[resourceID]

	if ok {

		fieldDesc, ok := resDescStruct.ResourceFields[fieldID]
		if ok && fieldDesc != "" {

			return fieldDesc
		}
	}

	resDescStruct, ok = descriptionsMap[resourceID]

	if ok {
		fieldDesc, ok := resDescStruct.ResourceFields[fieldID]
		if ok {
			return fieldDesc
		}
	}

	return desc
}

func getActionDescription(resourceID string, actionID string) string {
	desc := "This is the " + actionID + " action"

	resDescStruct, ok := descriptionsOverrideMap[resourceID]

	if ok {
		actionDesc, ok := resDescStruct.ResourceFields[actionID]
		if ok && actionDesc != "" {
			return actionDesc
		}
	}

	resDescStruct, ok = descriptionsMap[resourceID]

	if ok {
		actionDesc, ok := resDescStruct.ResourceActions[actionID]
		if ok {
			return actionDesc
		}
	}

	return desc
}
