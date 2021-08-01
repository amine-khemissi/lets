package main

import (
	"flag"
	"github.com/amine-khemissi/lets/imple"
	"github.com/amine-khemissi/lets/logger"
	"log"
)

var (
	command      string
	target       string
	commandsFile string
	only         bool
	sync         bool
	debugEnabled bool
)

func init() {
	flag.StringVar(&commandsFile, "f", "lets.yaml", "commands file path")
	flag.StringVar(&command, "c", "", "Command to execute")
	flag.StringVar(&target, "t", "", "target repo")
	flag.BoolVar(&debugEnabled, "d", false, "enable Debug mode")
	flag.BoolVar(&only, "o", false, "execute Only command without befores")
	flag.BoolVar(&sync, "s", false, "Sync with remote repositories")
	flag.Parse()

	logLevel := logger.INFO
	if debugEnabled {
		logLevel = logger.DEBUG
	}
	logger.Init(logLevel)

}

func main() {

	logger.Instance().Debug("parsing", commandsFile)
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
		if err = k.Exec(command, target, only); err != nil {
			log.Fatal(err)
		}
	}
}
