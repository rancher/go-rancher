package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/rancherio/go-rancher/client"
)

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()

	DEFAULT_RANCHER_URL := os.Getenv("RANCHER_URL")
	DEFAULT_ACCESS_KEY := os.Getenv("DEFAULT_ACCESS_KEY")

	rancherUrl := flag.String("url", DEFAULT_RANCHER_URL, "the url of the rancher server")
	accessKey := flag.String("access-key", DEFAULT_ACCESS_KEY, "the access key for the rancher server")
	flag.Parse()

	if !(rancherUrl != nil && len(*rancherUrl) > 0) {
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
	schema, err := json.MarshalIndent(rancherClient.Schemas, "  ", "    ")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(schema))
}
