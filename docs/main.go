package main

import (
	"flag"
	log "github.com/Sirupsen/logrus"
	"os"
)

var (
	command = flag.String("command", "", "generate-docs | generate-description")
)

func main() {

	flag.Parse()

	if *command != "" {

		switch *command {
		case "generate-description":
			log.Info("Generating the api descriptions file...")
			err := generateDescriptionFile()
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
		}
	} else {
		log.Info("Please provide a command [generate-docs | generate-description]")
	}
}
