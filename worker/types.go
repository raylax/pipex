package worker

import (
	"gopkg.in/yaml.v3"
)

type StringArray []string

type StepDefine struct {
	Name    string  `yaml:"name"`
	Timeout Timeout `yaml:"timeout"`

	// 内建参数
	Run StringArray `yaml:"run"`
}

type JobDefine struct {
	Name      string       `yaml:"name"`
	Steps     []StepDefine `yaml:"steps"`
	When      string       `yaml:"when"`      // When expression
	Container string       `yaml:"container"` // Container
	Timeout   Timeout      `yaml:"timeout"`
}

type ConfigDefine struct {
	Name string `yaml:"name"`
}

func (a *StringArray) UnmarshalYAML(value *yaml.Node) error {
	var multi []string
	err := value.Decode(&multi)
	if err != nil {
		var single string
		err := value.Decode(&single)
		if err != nil {
			return err
		}
		*a = []string{single}
	} else {
		*a = multi
	}
	return nil
}
