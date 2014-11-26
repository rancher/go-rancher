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
	creatable          bool
	updatable          bool
	deletable          bool
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

	rancherUrl := flag.String("url", DEFAULT_RANCHER_URL, "the url of the rancher server")
	accessKey := flag.String("access-key", DEFAULT_ACCESS_KEY, "the access key for the rancher server")

	flag.Parse()

	if len(*rancherUrl) == 0 {
		panic("RANCHER_URL cannot be empty, please set environment variable RANCHER_URL or use opt --url")
	}

	opts := new(client.ClientOpts)
	opts.Url = *rancherUrl
	if len(*accessKey) > 0 {
		opts.AccessKey = *accessKey
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
		idCreatable := false
		idUpdatable := false
		idDeletable := false
		flagSetForId := flag.NewFlagSet(id, flag.ExitOnError)
		for _, method := range dataUnit.CollectionMethods {
			switch method {
			case PUT_METHOD:
				idCreatable = true
			case POST_METHOD:
				idUpdatable = true
			case DELETE_METHOD:
				idDeletable = true
			default:
			}
		}
		ResourceFieldInfos := make(map[string]ResourceFieldInfo)
		for resourceFieldKey := range dataUnit.ResourceFields {
			resourceFieldValue, ok := dataUnit.ResourceFields[resourceFieldKey]
			if !ok {
				continue
			}
			ResourceFieldInfos[resourceFieldKey] = ResourceFieldInfo{resourceFieldValue.Create, resourceFieldValue.Update}
			switch resourceFieldValue.Type {
			case "string", "date", "blob", "password":
				flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey)
			case "int":
				flagSetForId.Int(resourceFieldKey, 0, "set the int value for "+resourceFieldKey)
			case "float":
				flagSetForId.Float64(resourceFieldKey, 0.0, "set the flaot value for "+resourceFieldKey)
			case "boolean":
				flagSetForId.Bool(resourceFieldKey, false, "set the bool value for "+resourceFieldKey)
			default:
				if strings.HasPrefix(resourceFieldValue.Type, "reference[") {
					flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey)
				}
				//flagSetForId.Var(resourceFieldValue.Default, resourceFieldKey, "set the "+resourceFieldKey)
			}
		}
		SchemaInfos[id] = SchemaInfo{flagSetForId, ResourceFieldInfos, idCreatable, idUpdatable, idDeletable}
	}

	for index, arg := range args[1:] {
		if info, ok := SchemaInfos[arg]; ok {
			//pivoted onto an id
			rInfo := info.resourceFieldInfos
			fmt.Println(args[index+2:])
			switch args[index+2] {
			case "create":
				if info.creatable {
					// do something
					info.flagSet.Parse(args[index+2:])
					fl := info.flagSet
					fl.Visit(func(fx *flag.Flag) {
						if rInfo[fx.Name].creatable {
							fmt.Println(fx.Value)
						} else {
							fl.PrintDefaults()
						}
					})
				} else {
					info.flagSet.PrintDefaults()
					break
				}
			default:
				info.flagSet.PrintDefaults()
				break
			}
			break
		}
	}
}
