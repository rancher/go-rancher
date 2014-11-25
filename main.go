package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/rancherio/go-rancher/client"
)

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()
	rancherUrl := flag.String("url", "http://localhost:8080/v1", "the url of the rancher server")
	opts := new(client.ClientOpts)
	opts.Url = *rancherUrl
	rancherClient, err := client.NewRancherClient(opts)
	if err != nil {
		panic(err.Error())
	}
	schema, err := json.MarshalIndent(rancherClient.Schemas, "  ", "    ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(schema))
}
