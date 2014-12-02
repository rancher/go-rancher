package cli

import (
	"encoding/json"
	"flag"
	"github.com/rancherio/go-rancher/client"
	"github.com/rancherio/go-rancher/util"
	"io/ioutil"
	"testing"
)

var (
	executeBeforeTest = false
	schemaTypeMap     map[string]client.Schema
	outStr            string
)

func beforeTest(t *testing.T) {
	if !executeBeforeTest {
		executeBeforeTest = true
		body, err := ioutil.ReadFile("cli_test_schema.json")
		if err != nil {
			t.Fatal("failed to load file")
		}
		var schema client.Schemas
		jsonErr := json.Unmarshal(body, &schema)
		if jsonErr != nil {
			t.Fatal("could not read json from loaded file cli_test_schema.json")
		}
		schemaTypeMap = make(map[string]client.Schema)
		for _, dataUnit := range schema.Data {
			schemaTypeMap[dataUnit.Resource.Id] = dataUnit
		}
	}
}

//pick an arbitrary schema type and check that the resource controls are set properly
func TestSetResourceControls(t *testing.T) {
	beforeTest(t)
	list, create, update, deleteFlag := setResourceControls(schemaTypeMap["authorized"].ResourceMethods, schemaTypeMap["authorized"].CollectionMethods)
	if !(list && create && update && deleteFlag) {
		t.Error("error in setting the resource control flags")
	}
}

//pick an arbitrary schema type and check that the flagSets are being set correctly
func TestProcessSchemaInfos(t *testing.T) {
	beforeTest(t)
	schemaInfos := processSchemaInfos(&[]client.Schema{schemaTypeMap["container"]})
	expectedFieldNames := []string{
		"agentId",
		"allocationState",
		"command",
		"commandArgs",
		"compute",
		"count",
		"created",
		"credentialIds",
		"data",
		"description",
		"directory",
		"dockerVolumes",
		"domain",
		"environment",
		"firstRunning",
		"hostname",
		"id",
		"imageId",
		"imageUuid",
		"instanceLinks",
		"instanceTriggeredStop",
		"kind",
		"memoryMb",
		"name",
		"networkIds",
		"ports",
		"privileged",
		"publishAllPorts",
		"removeTime",
		"removed",
		"requestedHostId",
		"startOnCreate",
		"subnetIds",
		"token",
		"transitioningMessage",
		"transitioningProgress",
		"user",
		"userdata",
		"uuid",
		"validHostIds",
		"vnetIds"}
	realCount := 0
	expectedCount := len(expectedFieldNames)
	schemaInfos["container"].flagSet.VisitAll(func(fx *flag.Flag) {
		if util.Contains(expectedFieldNames, fx.Name) {
			realCount++
		}
	})
	if realCount != expectedCount {
		t.Error("error in adding all fields as flag arguments")
	}
}
