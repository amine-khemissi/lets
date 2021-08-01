package imple

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Kubtan struct {
	Repositories map[string]Repository `yaml:"repositories"`
	Commands     map[string]Command    `yaml:"commands"`
}

func New(filePath string) (Kubtan, error) {
	kubtan := Kubtan{}
	yfile, err := ioutil.ReadFile(filePath)
	if err != nil {

		return kubtan, err
	}
	err2 := yaml.Unmarshal(yfile, &kubtan)

	if err2 != nil {
		return kubtan, err2
	}

	if kubtan.detectLoop() {
		return kubtan, errors.New("detected a loop")
	}

	return kubtan, nil
}

func (k Kubtan) Exec(cmdName string, only bool) error {
	cmd, found := k.Commands[cmdName]
	if !found {
		return errors.New(fmt.Sprintf("%s not found", cmdName))
	}

	if err := cmd.Exec(k.Commands, only); err != nil {
		return errors.New(fmt.Sprintf("failed to execute %s, reason: %s", cmdName, err.Error()))
	}

	return nil
}

func (k Kubtan) detectLoop() bool {
	for cmd, v := range k.Commands {
		if k.detectLoopForOneCommand(cmd, v, k.Commands, nil) {
			return true
		}
	}
	return false

}

func (k Kubtan) detectLoopForOneCommand(cmd string, v Command, commands map[string]Command, registeredCommands map[string]struct{}) bool {
	if registeredCommands == nil {
		registeredCommands = map[string]struct{}{cmd: {}}
	}
	for _, subCmd := range v.Befores {
		if _, found := registeredCommands[subCmd]; found {
			return true
		}
		registeredCommands[subCmd] = struct{}{}
		if k.detectLoopForOneCommand(subCmd, commands[subCmd], commands, registeredCommands) {
			return true
		}
	}
	return false
}

func (k Kubtan) Sync() error {
	for name, repo := range k.Repositories {
		fmt.Println("cd", name)
		if err := repo.Sync(); err != nil {
			return errors.New(fmt.Sprintf("failed to sync %s, reason: %s", name, err.Error()))
		}
		fmt.Println("cd ..")
	}
	return nil
}
