package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rancherio/go-rancher/util"
	"os"
	"strings"

	"github.com/rancherio/go-rancher/client"
)

const (
	POST_METHOD   string = "POST"
	GET_METHOD    string = "GET"
	DELETE_METHOD string = "DELETE"
	PUT_METHOD    string = "PUT"
	containerId   string = "imageId"
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
	allowedActions     []string
}

var helpMessage string

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
		return errors.New("map already set")
	}
	return json.Unmarshal([]byte(value), i)
}

func printFormat(format string, obj interface{}) {
	if format == "json" {
		resp, jsonErr := json.MarshalIndent(obj, "  ", "    ")
		if jsonErr != nil {
			panic(jsonErr.Error())
		}
		fmt.Println(string(resp))
	} else if format == "tabular" {
		PrintTabular(obj)
	} else {
		panic("unknown format " + format)
	}
}

func printTypeLevelHelpMessageAndExit(rancherUrl string, accesKey string) {
	filename := os.Args[0]
	fmt.Println("usage of " + filename + ":")
	helpMessage = "[create delete update list] for the following types; \nuse [TYPE] [action] -h to view settable fields" + helpMessage
	helpMessage = "--accessKey=" + accesKey + "\"the accesKey required to talk to the rancher server\" \n" + helpMessage
	helpMessage = "--url=" + rancherUrl + "\"the url of the rancher server\" \n" + helpMessage
	fmt.Println(helpMessage)
	os.Exit(2)
}

func processSchemaInfos(data *[]client.Schema) map[string]SchemaInfo {
	SchemaInfos := make(map[string]SchemaInfo)
	for _, dataUnit := range *data {
		if dataUnit.Resource.Type != "schema" {
			continue
		}
		id := dataUnit.Resource.Id

		idListable, idCreatable, idUpdatable, idDeletable := setResourceControls(dataUnit.ResourceMethods, dataUnit.CollectionMethods)
		helpMessage = helpMessage + "-" + id + "\n"
		flagSetForIdCreate := flag.NewFlagSet(id+" create", flag.ExitOnError)
		flagSetForIdUpdate := flag.NewFlagSet(id+" update", flag.ExitOnError)
		flagSetForIdDelete := flag.NewFlagSet(id+" delete", flag.ExitOnError)
		flagSetForIdList := flag.NewFlagSet(id+" list", flag.ExitOnError)
		ResourceFieldInfos := make(map[string]ResourceFieldInfo)

		actions := make([]string, 0)

		for key, _ := range dataUnit.Resource.Actions {
			actions = append(actions, key)
			var flagMapVar flagMap
			flagSetForIdList.Var(&flagMapVar, key, "perform the "+key+" action")
		}

		for resourceFieldKey := range dataUnit.ResourceFields {
			resourceFieldValue, ok := dataUnit.ResourceFields[resourceFieldKey]
			//if field cannot be updated or created, dont bother adding it to the flagset
			if !ok || !(resourceFieldValue.Create || resourceFieldValue.Update || resourceFieldValue.Nullable) {
				continue
			}
			ResourceFieldInfos[resourceFieldKey] = ResourceFieldInfo{resourceFieldValue.Create, resourceFieldValue.Update, resourceFieldValue.Required}
			requiredString := ""
			if resourceFieldValue.Required {
				requiredString = "(required)"
			}
			//use closure to get valid flagsets for field
			validFlagSets := func() []*flag.FlagSet {
				returnFlagSets := make([]*flag.FlagSet, 0, 4)
				if resourceFieldValue.Create {
					returnFlagSets = append(returnFlagSets, flagSetForIdCreate)
				}
				if resourceFieldValue.Update {
					returnFlagSets = append(returnFlagSets, flagSetForIdUpdate)
				}
				if resourceFieldValue.Nullable {
					returnFlagSets = append(returnFlagSets, flagSetForIdDelete)
				}
				return append(returnFlagSets, flagSetForIdList)
			}()

			for _, flagSetForId := range validFlagSets {
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
					} else if strings.HasPrefix(resourceFieldValue.Type, "array[") {
						var flagListVar flagList
						flagSetForId.Var(&flagListVar, resourceFieldKey, "set the list (comma seperated) value for "+resourceFieldKey+requiredString)
					} else if strings.HasPrefix(resourceFieldValue.Type, "map[") {
						var flagMapVar flagMap
						flagSetForId.Var(&flagMapVar, resourceFieldKey, "set the map value for "+resourceFieldKey+requiredString)
					} else {
						//default to string
						flagSetForId.String(resourceFieldKey, "", "set the string value for "+resourceFieldKey+requiredString)
					}
				}
			}
		}
		SchemaInfos[id] = SchemaInfo{flagSetForIdList, ResourceFieldInfos, idListable, idCreatable, idUpdatable, idDeletable, actions}
		SchemaInfos[id+"-create"] = SchemaInfo{flagSetForIdCreate, ResourceFieldInfos, idListable, idCreatable, idUpdatable, idDeletable, actions}
		SchemaInfos[id+"-update"] = SchemaInfo{flagSetForIdUpdate, ResourceFieldInfos, idListable, idCreatable, idUpdatable, idDeletable, actions}
		SchemaInfos[id+"-delete"] = SchemaInfo{flagSetForIdDelete, ResourceFieldInfos, idListable, idCreatable, idUpdatable, idDeletable, actions}
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
		//str := recover()
		//if str != nil {
		//	fmt.Print("ERROR: ")
		//	fmt.Println(str)
		//}
	}()

	rancherUrl := flag.String("url", DEFAULT_RANCHER_URL, "the url of the rancher server")
	accessKey := flag.String("access-key", DEFAULT_ACCESS_KEY, "the access key for the rancher server")
	format := flag.String("format", "tabular", "the format in which the output should be printed")

	helpFlag := false

	for index, arg := range os.Args {
		if strings.Contains(arg, "url=") {
			*rancherUrl = strings.Split(arg, "=")[1]
			continue
		}
		if strings.Contains(arg, "accessKey=") {
			*accessKey = strings.Split(arg, "=")[1]
			continue
		}
		//break if we start parsing a schemaType
		if arg == "-h" || arg == "--h" || arg == "--help" || arg == "-help" || arg == "?" {
			//use heuristics to determine if help is being requested for main command or subcommand
			prev := os.Args[index-1]
			if !(prev == "create" || prev == "delete" || prev == "list" || prev == "update") {
				helpFlag = true
			}
		}
	}
	if !helpFlag {
		flag.Parse()
	}

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
			//pivoted onto a valid schema type
			rInfo := info.resourceFieldInfos
			if index+1 >= len(args) {
				info.flagSet.PrintDefaults()
				panic("no valid operation specified [create, update, delete, list]")
			}
			switch args[index+1] {
			case "create":
				info = SchemaInfos[arg+"-create"]
				rInfo = info.resourceFieldInfos
				if info.creatable {
					info.flagSet.Parse(args[index+2:])
					fl := info.flagSet
					reqObj := make(map[string]interface{})
					fl.Visit(func(fx *flag.Flag) {
						if rInfo[fx.Name].creatable {
							reqObj[fx.Name] = fx.Value
						} else {
							fmt.Println("Usage for " + arg + " create:")
							fl.PrintDefaults()
							panic(fx.Name + " not marked as creatable")
						}
					})
					//container id has not been parsed out yet
					if len(fl.Args()) > 0 {
						reqObj[containerId] = fl.Args()[0]
					}
					respObj := make(map[string]interface{})
					err := rancherClient.Create(arg, reqObj, &respObj)
					if err != nil {
						panic(err.Error())
					}
					printFormat(*format, respObj)
				} else {
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
					if len(fl.Args()) > 0 {
						reqObj[containerId] = fl.Args()[0]
					}
					respObj := make(map[string]interface{})
					listOpts := client.NewListOpts()
					listOpts.Filters = reqObj
					err := rancherClient.List(arg, listOpts, &respObj)
					if err != nil {
						panic(err.Error())
					}
					printFormat(*format, respObj)
				} else {
					panic(arg + "not marked as listable")
				}
			case "delete":
				info = SchemaInfos[arg+"-delete"]
				rInfo = info.resourceFieldInfos
				if info.deletable {
					resource := rancherClient.Types[arg].Resource
					err := rancherClient.Delete(arg, &resource)
					if err != nil {
						panic(err.Error())
					}
				} else {
					panic(arg + " not marked as deletable")
				}
			case "update":
				info = SchemaInfos[arg+"-update"]
				rInfo = info.resourceFieldInfos
				if info.updatable {
					resource := rancherClient.Types[arg].Resource
					reqObj := make(map[string]interface{})
					fl := info.flagSet
					fl.Parse(args[index+2:])
					fl.Visit(func(fx *flag.Flag) {
						if rInfo[fx.Name].updatable {
							reqObj[fx.Name] = fx.Value
						} else {
							fmt.Println("Usage for " + arg + " update:")
							fl.PrintDefaults()
							panic(fx.Name + " not marked as updatable")
						}
					})
					if len(fl.Args()) > 0 {
						reqObj[containerId] = fl.Args()[0]
					}
					respObj := make(map[string]interface{})
					err := rancherClient.Update(arg, &resource, reqObj, respObj)
					if err != nil {
						panic(err.Error())
					}
					printFormat(*format, respObj)
				} else {
					panic(arg + " not marked as updatable")
				}
			default:
				//check if it is an action
				if util.Contains(SchemaInfos[arg].allowedActions, args[index+1]) {
					info := SchemaInfos[arg]
					var reqObj map[string]interface{}
					fl := info.flagSet
					fl.Parse(args[index+2:])
					fl.Visit(func(fx *flag.Flag) {
						if fx.Name == "argMap" {
							reqObj = map[string]interface{}(*fx.Value.(*flagMap))
						}
					})
					if len(fl.Args()) > 0 {
						reqObj[containerId] = fl.Args()[0]
					}
					respObj := make(map[string]interface{})
					err := rancherClient.Action(arg, args[index+1], reqObj, respObj)
					if err != nil {
						panic(err.Error())
					}
					printFormat(*format, respObj)
					break
				}
				info.flagSet.PrintDefaults()
				panic("unknown subcommand " + args[index+2])
			}
			parsedFlag = true
			break
		}
	}
	if !parsedFlag {
		printTypeLevelHelpMessageAndExit(DEFAULT_RANCHER_URL, DEFAULT_ACCESS_KEY)
	}
}
