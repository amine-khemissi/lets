package imple

import (
	"errors"
	"fmt"
)

type Command struct {
	Actions []string `yaml:"actions"`
	Befores []string `yaml:"befores"`
}

func (c Command) Exec(commands map[string]Command, only bool) error {

	if !only {
		for _, before := range c.Befores {
			cmd, found := commands[before]
			if !found {
				return errors.New(fmt.Sprintf("%s not found", before))
			}
			if err := cmd.Exec(commands, false); err != nil {
				return errors.New(fmt.Sprintf("failed to execute %s, reason: %s", before, err.Error()))
			}
		}
	}
	for _, action := range c.Actions {
		fmt.Println("->", action)
	}
	return nil
}
