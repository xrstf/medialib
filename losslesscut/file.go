package losslesscut

import (
	"os"

	"go.xrstf.de/medialib/losslesscut/types"

	"gopkg.in/yaml.v3"
)

func ReadFile(filename string) (*types.Project, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var project types.Project

	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&project); err != nil {
		return nil, err
	}

	return &project, err
}
