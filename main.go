package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rancherio/go-rancher/client"
)

const (
	POST_METHOD   string = "POST"
	GET_METHOD    string = "GET"
	DELETE_METHOD string = "DELETE"
	PUT_METHOD    string = "PUT"
)

type schemaInfo struct {
	creatable bool
	updatable bool
	deletable bool
	listable  bool
}

var schemaInfos map[string]schemaInfo

func main() {
	defer func() {
		str := recover()
		if str != nil {
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

	//args := os.Args

	schemaInfos = make(map[string]schemaInfo)

	for _, dataUnit := range rancherClient.Schemas.Data {
		if dataUnit.Resource.Type != "schema" {
			continue
		}
		id := dataUnit.Resource.Id
		var typeInfoObj schemaInfo = schemaInfo{false, false, false, false}
		for _, method := range dataUnit.CollectionMethods {
			switch method {
			case GET_METHOD:
				typeInfoObj.listable = true
				break
			case PUT_METHOD:
				typeInfoObj.updatable = true
				break
			case POST_METHOD:
				typeInfoObj.creatable = true
				break
			case DELETE_METHOD:
				typeInfoObj.deletable = true
				break
			}
			schemaInfos[id] = typeInfoObj
		}
	}
	fmt.Println(schemaInfos["container"].listable)
}
