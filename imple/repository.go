package imple

import "fmt"

type Repository struct {
	Path    string `yaml:"path"`
	Version string `yaml:"version"`
}

func (r Repository) Sync() error {
	fmt.Println("git clone", r.Path)
	fmt.Println("git checkout", r.Version)
	return nil
}
