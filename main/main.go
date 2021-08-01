package main

import (
	"flag"
	"github.com/amine-khemissi/kubtan/imple"
	"log"
)

var (
	command      string
	commandsFile string
	only         bool
	sync         bool
)

func init() {
	flag.StringVar(&commandsFile, "f", "commands.yaml", "commands.yaml File path")
	flag.StringVar(&command, "c", "", "Command to execute")
	flag.BoolVar(&only, "o", false, "execute Only command without befores")
	flag.BoolVar(&sync, "s", false, "sync with remote repositories")
	flag.Parse()
}

func main() {

	k, err := imple.New(commandsFile)
	if err != nil {
		log.Fatal(err)
	}
	if sync {
		if err = k.Sync(); err != nil {
			log.Fatal(err)
		}
	}
	if command != "" {
		if err = k.Exec(command, only); err != nil {
			log.Fatal(err)
		}
	}
}
