package core

import (
	"github.com/go-yaml/yaml"
	"os"
)

type Character struct {
	base     Base `yaml:"base"`
	weapon   *Weapon
	artifact *Artifact
	final    Base
}

type skill struct {
	sType    string    `yaml:"SType"`
	eType    string    `yaml:"EType"`
	value    []int     `yaml:"Value"`
	reaction []float32 `yaml:"reaction"` //reaction rate
}

func (c *Character) Azone() float32 {
	return 0
}

func ParseCharacter(path string) *Character {
	file, err := os.ReadFile(path)
	if err != nil {

	}
	character := &Character{}
	err = yaml.Unmarshall(file, character)
	return nil
}
