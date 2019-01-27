package termi

import (
	"io"
	"os"
	"path"
	"runtime"
	"text/template"

	env "github.com/Netflix/go-env"
)

type FlagSet interface {
	SetDescription(description string)

	SetEnvironment(variable interface{})

	Register(flag Flag)

	Parse(args []string) ([]string, error)

	ParseEnvironment() error

	PrintDescription(w io.Writer) error
}

type set struct {
	description string
	flags       []Flag
	envs        []interface{}
}

func NewFlagSet() FlagSet {
	return &set{}
}

func (s *set) PrintDescription(w io.Writer) error {
	if w == nil {
		return nil
	}
	t, err := template.New("termi.usage").Parse(s.description)
	if err != nil {
		return err
	}
	var values map[string]string
	for _, env := range s.envs {
		results, err := EnvironmentDescription(env)
		if err != nil {
			return err
		}
		for name, description := range results {
			values[name] = description
		}
	}
	return t.Execute(w, map[string]interface{}{
		"name":         path.Base(os.Args[0]),
		"GoVersion":    runtime.Version(),
		"environments": values,
		"flags":        s.flags,
	})
}

func (s *set) ParseEnvironment() error {
	for _, e := range s.envs {
		if _, err := env.UnmarshalFromEnviron(e); err != nil {
			return err
		}
	}
	return nil
}

func (s *set) Parse(args []string) ([]string, error) {
	var remainder []string
	for i := 0; i < len(args)-1; i++ {
		for _, flag := range s.flags {
			if !flag.IsFlag(args[i]) {
				remainder = append(remainder, args[i])
				continue
			}
			if err := flag.Set(args[i+1]); err != nil {
				return remainder, err
			}
			i++
		}
	}
	return remainder, nil
}

func (s *set) Register(f Flag) {
	s.flags = append(s.flags, f)
}

func (s *set) SetDescription(description string) {
	s.description = description
}

func (s *set) SetEnvironment(variable interface{}) {
	s.envs = append(s.envs, variable)
}
