package main

import "github.com/rancherio/go-rancher/cli"
import "os"

func main() {
	rancherUrl := os.Getenv("RANCHER_URL")
	accessKey := os.Getenv("RANCHER_ACCESS_KEY")

	cli.ParseCli(rancherUrl, accessKey)
}
