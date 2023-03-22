package worker

import (
	"gopkg.in/yaml.v3"
)

type StringArray []string

type StepDefine struct {
	Name    string  `yaml:"name"`
	Timeout Timeout `yaml:"timeout"`

	Checkout string `yaml:"checkout" builtin:"checkout"`

	// 内建参数 - run
	Sh         StringArray `yaml:"sh" builtin:":run"`
	Bash       StringArray `yaml:"bash" builtin:":run"`
	Cmd        StringArray `yaml:"cmd" builtin:":run"`
	Powershell StringArray `yaml:"powershell" builtin:":run"`

	// 插件
	Use  string            `yaml:"use"`
	With map[string]string `yaml:"with"`
}

type TaskDefine struct {
	Name      string       `yaml:"name"`
	Steps     []StepDefine `yaml:"steps"`
	When      string       `yaml:"when"`      // When expression
	Container string       `yaml:"container"` // Container
	Timeout   Timeout      `yaml:"timeout"`
}

type PipelineDefine struct {
	Name string `yaml:"name"`
}

func (a *StringArray) UnmarshalYAML(value *yaml.Node) error {
	arr, err := unmarshalYAMLArray[string](value)
	if err != nil {
		return err
	}
	*a = arr
	return nil
}

func unmarshalYAMLArray[T any](value *yaml.Node) ([]T, error) {
	var multi []T
	err := value.Decode(&multi)
	if err != nil {
		var single T
		err := value.Decode(&single)
		if err != nil {
			return nil, err
		}
		return []T{single}, nil
	} else {
		return multi, nil
	}
}
