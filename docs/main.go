package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"os"
)

var (
	command  = flag.String("command", "", "generate-docs | generate-description")
	version  = flag.String("version", "v1.0", "docs version")
	language = flag.String("lang", "en", "docs language")
	layout   = flag.String("layout", "rancher-default", "docs layout")
)

func main() {

	flag.Parse()

	if *command != "" {

		switch *command {
		case "generate-description":
			log.Info("Generating the api descriptions file...")
			err := generateDescriptionFile(false)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("Done...")
		case "generate-docs":
			log.Info("Generating the api docs...")
			err := generateFiles()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("Done...")
		case "generate-empty-description":
			log.Info("Generating the api descriptions file with empty descriptions...")
			err := generateDescriptionFile(true)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			log.Info("Done...")
		}
	} else {
		log.Info("Please provide a command [generate-docs | generate-description]")
	}
}
