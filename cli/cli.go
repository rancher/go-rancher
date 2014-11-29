package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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
	required  bool
}

type SchemaInfo struct {
	flagSet            *flag.FlagSet
	resourceFieldInfos map[string]ResourceFieldInfo
	listable           bool
	creatable          bool
	updatable          bool
	deletable          bool
}

var SchemaInfos map[string]SchemaInfo

type flagList []string

func (i *flagList) String() string {
	return fmt.Sprint(*i)
}

func (i *flagList) Set(value string) error {
	for _, arg := range strings.Split(value, ",") {
		*i = append(*i, arg)
	}
	return nil
}

type flagMap map[string]interface{}

func (i *flagMap) String() string {
	return fmt.Sprint(*i)
}

func (i *flagMap) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("interval flag already set")
	}
	return json.Unmarshal([]byte(value), i)
}

func processSchemaInfos(data *[]client.Schema) map[string]SchemaInfo {
	SchemaInfos := make(map[string]SchemaInfo)
	for _, dataUnit := range *data {
		if dataUnit.Resource.Type != "schema" {
			continue
		}
		id := dataUnit.Resource.Id

		idListable, idCreatable, idUpdatable, idDeletable := setResourceControls(dataUnit.ResourceMethods, dataUnit.CollectionMethods)

		flagSetForId := flag.NewFlagSet(id, flag.ExitOnError)
		ResourceFieldInfos := make(map[string]ResourceFieldInfo)

		for resourceFieldKey := range dataUnit.ResourceFields {
			resourceFieldValue, ok := dataUnit.ResourceFields[resourceFieldKey]
			//if field cannot be updated or created, dont bother adding it to the flagset
			if !ok || !(resourceFieldValue.Create || resourceFieldValue.Update) {
				continue
			}
			ResourceFieldInfos[resourceFieldKey] = ResourceFieldInfo{resourceFieldValue.Create, resourceFieldValue.Update, resourceFieldValue.Required}
			requiredString := ""
			if resourceFieldValue.Required {
				requiredString = "(required)"
			}
			switch resourceFieldValue.Type {
			case "int":
				flagSetForId.Int(resourceFieldKey, 0, "set the int value for "+resourceFieldKey+requiredString)
			case "float":
				flagSetForId.Float64(resourceFieldKey, 0.0, "set the flaot value for "+resourceFieldKey+requiredString)
			case "boolean":
				flagSetForId.Bool(resourceFieldKey, false, "set the bool value for "+resourceFieldKey+requiredString)
			case "string", "date", "blob", "password":
				//default to string
				flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey+requiredString)
			default:
				if strings.HasPrefix(resourceFieldValue.Type, "reference[") {
					flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey+requiredString)
				} else if strings.HasPrefix(resourceFieldKey, "array[") {
					var flagListVar flagList
					flagSetForId.Var(&flagListVar, resourceFieldKey, "set the list (comma seperated) value for "+resourceFieldKey+requiredString)
				} else if strings.HasPrefix(resourceFieldKey, "map[") {
					var flagMapVar flagMap
					flagSetForId.Var(&flagMapVar, resourceFieldKey, "set the string value for "+resourceFieldKey+requiredString)
				} else {
					//default to string
					flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey+requiredString)
				}
			}
		}
		SchemaInfos[id] = SchemaInfo{flagSetForId, ResourceFieldInfos, idListable, idCreatable, idUpdatable, idDeletable}
	}
	return SchemaInfos
}

func setResourceControls(resourceMethods []string, collectionMethods []string) (bool, bool, bool, bool) {
	idListable := false
	idCreatable := false
	idUpdatable := false
	idDeletable := false
	for _, method := range resourceMethods {
		switch method {
		case PUT_METHOD:
			idUpdatable = true
		case DELETE_METHOD:
			idDeletable = true
		}
	}

	for _, method := range collectionMethods {
		switch method {
		case GET_METHOD:
			idListable = true
		case POST_METHOD:
			idCreatable = true
		}
	}
	return idListable, idCreatable, idUpdatable, idDeletable
}

func ParseCli(DEFAULT_RANCHER_URL string, DEFAULT_ACCESS_KEY string) {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Print("ERROR: ")
			fmt.Println(str)
		}
	}()

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

	SchemaInfos := processSchemaInfos(&rancherClient.Schemas.Data)

	parsedFlag := false

	args := flag.Args()

	for index, arg := range args {
		if info, ok := SchemaInfos[arg]; ok {
			//pivoted onto a schema type
			rInfo := info.resourceFieldInfos
			if index+1 >= len(args) {
				info.flagSet.PrintDefaults()
				panic("no valid operation specified [create, update, delete, list]")
			}
			switch args[index+1] {
			case "create":
				if info.creatable {
					info.flagSet.Parse(args[index+2:])
					fl := info.flagSet
					reqObj := make(map[string]interface{})
					fl.Visit(func(fx *flag.Flag) {
						if rInfo[fx.Name].creatable {
							reqObj[fx.Name] = fx.Value
						} else {
							fl.PrintDefaults()
							panic(fx.Name + " not marked as creatable")
						}
					})
					respObj := make(map[string]interface{})
					err := rancherClient.Create(arg, reqObj, &respObj)
					if err != nil {
						fl.PrintDefaults()
						panic(err.Error())
					}
					fmt.Println(respObj)
				} else {
					info.flagSet.PrintDefaults()
					panic(arg + " not marked as creatable")
				}
			case "list":
				if info.listable {
					info.flagSet.Parse(args[index+2:])
					fl := info.flagSet
					reqObj := make(map[string]interface{})
					fl.Visit(func(fx *flag.Flag) {
						reqObj[fx.Name] = fx.Value
					})
					respObj := make(map[string]interface{})
					listOpts := client.NewListOpts()
					listOpts.Filters = reqObj
					err := rancherClient.List(arg, listOpts, &respObj)
					if err != nil {
						fl.PrintDefaults()
						panic(err.Error())
					}
					resp, jsonErr := json.MarshalIndent(respObj, "  ", "    ")
					if jsonErr != nil {
						panic(jsonErr.Error())
					}
					fmt.Println(string(resp))
				} else {
					info.flagSet.PrintDefaults()
					panic(arg + "not marked as listable")
				}
			case "delete":
				if info.deletable {
					resource := rancherClient.Types[arg].Resource
					err := rancherClient.Delete(arg, &resource)
					if err != nil {
						info.flagSet.PrintDefaults()
						panic(err.Error())
					}
				} else {
					info.flagSet.PrintDefaults()
					panic(arg + " not marked as deletable")
				}
			case "update":
				if info.updatable {
					resource := rancherClient.Types[arg].Resource
					reqObj := make(map[string]interface{})
					fl := info.flagSet
					fl.Parse(args[index+2:])
					fl.Visit(func(fx *flag.Flag) {
						reqObj[fx.Name] = fx.Value
					})
					respObj := make(map[string]interface{})
					err := rancherClient.Update(arg, &resource, reqObj, respObj)
					if err != nil {
						info.flagSet.PrintDefaults()
						panic(err.Error())
					}
				} else {
					info.flagSet.PrintDefaults()
					panic(arg + " not marked as updatable")
				}
			default:
				info.flagSet.PrintDefaults()
				panic("unknown subcommand " + args[index+2])
			}
			parsedFlag = true
			break
		}
	}
	if !parsedFlag {
		fmt.Println("unknown type")
	}
}
