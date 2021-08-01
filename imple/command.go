package imple

import (
	"fmt"
	"github.com/amine-khemissi/lets/errors"
	"github.com/amine-khemissi/lets/logger"
	"os"
	"strings"
)

type Command struct {
	Actions []string `yaml:"actions"`
	Befores []string `yaml:"befores"`
}

func (c Command) Exec(repoName, cmdName string, commands map[string]Command, only bool) error {

	os.Setenv("REPOSITORY_NAME", strings.ReplaceAll(repoName, "/", "-"))
	if !only {
		for _, before := range c.Befores {
			cmd, found := commands[before]
			if !found {
				return errors.New(before, "not found")
			}
			if err := cmd.Exec(repoName, before, commands, false); err != nil {
				return errors.Stack(err, "failed to execute", before)
			}
		}
	}
	logger.Instance().Info(fmt.Sprintf("* %s[%s]:", repoName, cmdName))
	wd, err := os.Getwd()
	if err != nil {
		return errors.Stack(err, "failed to get current working directory")
	}
	if err := os.Chdir(repoName); err != nil {
		return errors.Stack(err, "failed to change directory toward (", repoName, ")")
	}
	defer os.Chdir(wd)
	for _, action := range c.Actions {
		logger.Instance().Debug("->", action)
		if err := execAction(strings.Split(action, " ")[0], strings.Split(action, " ")[1:]...); err != nil {
			return errors.Stack(err, "failed to execAction action (", action, ")")
		}
	}
	return nil
}
