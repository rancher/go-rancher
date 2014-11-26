package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rancherio/go-rancher/client"
)

const (
	POST_METHOD   string = "POST"
	GET_METHOD    string = "GET"
	DELETE_METHOD string = "DELETE"
	PUT_METHOD    string = "PUT"
)

type ResourceFieldInfo struct {
	creatable bool
	updatable bool
}

type SchemaInfo struct {
	flagSet            *flag.FlagSet
	resourceFieldInfos map[string]ResourceFieldInfo
}

var SchemaInfos map[string]SchemaInfo

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Print("ERROR: ")
			fmt.Println(str)
		}
	}()

	DEFAULT_RANCHER_URL := os.Getenv("RANCHER_URL")
	DEFAULT_ACCESS_KEY := os.Getenv("RANCHER_ACCESS_KEY")

	//rancherUrl := flag.String("url", DEFAULT_RANCHER_URL, "the url of the rancher server")
	//accessKey := flag.String("access-key", DEFAULT_ACCESS_KEY, "the access key for the rancher server")

	rancherUrl := DEFAULT_RANCHER_URL
	accessKey := DEFAULT_ACCESS_KEY

	//flag.Parse()

	if len(rancherUrl) == 0 {
		panic("RANCHER_URL cannot be empty, please set environment variable RANCHER_URL or use opt --url")
	}

	opts := new(client.ClientOpts)
	opts.Url = rancherUrl
	if len(accessKey) > 0 {
		opts.AccessKey = accessKey
	}
	rancherClient, err := client.NewRancherClient(opts)
	if err != nil {
		panic(err.Error())
	}

	SchemaInfos = make(map[string]SchemaInfo)
	args := os.Args

	for _, dataUnit := range rancherClient.Schemas.Data {
		if dataUnit.Resource.Type != "schema" {
			continue
		}
		id := dataUnit.Resource.Id
		flagSetForId := flag.NewFlagSet(id, flag.ExitOnError)
		for _, method := range dataUnit.CollectionMethods {
			switch method {
			case GET_METHOD:
				break
			case PUT_METHOD:
				flagSetForId.Bool("update", false, "update a "+id)
				break
			case POST_METHOD:
				flagSetForId.Bool("create", false, "create a "+id)
				break
			case DELETE_METHOD:
				flagSetForId.Bool("delete", false, "delete a "+id)
				break
			}
			ResourceFieldInfos := make(map[string]ResourceFieldInfo)
			for resourceFieldKey := range dataUnit.ResourceFields {
				resourceFieldValue, ok := dataUnit.ResourceFields[resourceFieldKey]
				if !ok {
					continue
				}
				ResourceFieldInfos[resourceFieldKey] = ResourceFieldInfo{resourceFieldValue.Create, resourceFieldValue.Update}
				if id == "agentInstanceProvider" {
					fmt.Println(resourceFieldKey + ":" + resourceFieldValue.Type)
				}
				switch resourceFieldValue.Type {
				case "string", "date", "blob", "password":
					flagSetForId.String(resourceFieldKey, "", "set the "+resourceFieldKey)
					break
				case "int":
					flagSetForId.Int(resourceFieldKey, 0, "set the "+resourceFieldKey)
					break
				case "float":
					flagSetForId.Float64(resourceFieldKey, 0.0, "set the "+resourceFieldKey)
					break
				case "boolean":
					flagSetForId.Bool(resourceFieldKey, false, "set the "+resourceFieldKey)
					break
				default:
					if strings.HasPrefix(resourceFieldValue.Type, "reference[") {
						flagSetForId.String(resourceFieldKey, "", "set the "+resourceFieldKey)
					}
					//flagSetForId.Var(resourceFieldValue.Default, resourceFieldKey, "set the "+resourceFieldKey)
				}
			}
			SchemaInfos[id] = SchemaInfo{flagSetForId, ResourceFieldInfos}
		}
	}
	// program doesn't reach this block
	SchemaInfos["agentInstanceProvider"].flagSet.Parse(args[1:])
	fl := SchemaInfos["agentInstanceProvider"].flagSet.Lookup("accountId")
	fmt.Println(fl.Value.String())
}
