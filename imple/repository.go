package imple

import (
	"github.com/amine-khemissi/lets/errors"
	"github.com/amine-khemissi/lets/logger"
	"os"
)

type Repository struct {
	Origin  string `yaml:"origin"`
	Version string `yaml:"version"`
}

func (r Repository) Sync(name string) error {
	logger.Instance().Info("synchronise repo", name)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		if err := r.Clone(name); err != nil {
			return err
		}
	}
	wd, err := os.Getwd()
	if err != nil {
		return errors.Stack(err, "failed to get current working directory")
	}
	if err := os.Chdir(name); err != nil {
		return errors.Stack(err, "failed to change directory toward (", name, ")")
	}
	defer os.Chdir(wd)
	if err := execAction("git", "pull"); err != nil {
		return errors.Stack(err, "failed to pull")
	}

	if err := execAction("git", "checkout", r.Version); err != nil {
		return errors.Stack(err, "failed to checkout the repository to the version: ", r.Version)
	}
	return nil
}

func (r Repository) Clone(name string) error {
	if err := execAction("git", "clone", r.Origin, name); err != nil {
		return errors.Stack(err, "failed to clone the repository")
	}
	return nil
}
