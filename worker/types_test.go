package worker

import (
	"gopkg.in/yaml.v3"
	"log"
	"testing"
)

func TestStringArray_UnmarshalYAML(t *testing.T) {

	t.Run("single", func(t *testing.T) {
		var s = struct {
			Arr StringArray `yaml:"arr"`
		}{}
		err := yaml.Unmarshal(
			[]byte(`arr: a`), &s)
		if err != nil || len(s.Arr) != 1 || s.Arr[0] != "a" {
			log.Fatalf("err:%v arr:%v", err, s.Arr)
		}
	})

	t.Run("multiple", func(t *testing.T) {
		var s = struct {
			Arr StringArray `yaml:"arr"`
		}{}

		err := yaml.Unmarshal(
			[]byte(`
arr: 
  - a
  - b
`), &s)
		if err != nil || len(s.Arr) != 2 || s.Arr[0] != "a" || s.Arr[1] != "b" {
			log.Fatalf("err:%v arr:%v", err, s.Arr)
		}
	})

}
