package termi

import (
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	env "github.com/Netflix/go-env"
)

// FlagSet is my own version of the native 'flag.FlagSet'
// I felt as though it was missing features, so I am improving it as to how I like it.
type FlagSet interface {
	// SetDescription allows you to parse a text/template
	// in order to printed a more meaningful description
	// By default, the variables are:
	// - name         : name of the application
	// - GoVersion    : the version of go being used
	// - environments : a map of the variables that can be set by an env
	// - flags        : All the flags defined in the Set
	SetDescription(description string)

	// SetEnvironment allows you set parse a variable that is used
	// to gather environment values at runtime.
	SetEnvironment(variable interface{}) FlagSet

	// Register allows you to pass a flag object will be stored with this
	// Set
	Register(flag Flag) FlagSet

	// Parse will read all the strings passed and update flags when applied
	// then return the unused args.
	// When "--" is read in the args list, in remaining args will be returned unprocessed
	// after that point.
	Parse(args []string) ([]string, error)

	// ParseEnvironment will gather all environment variables and apply the values
	// to the stored environment values.
	ParseEnvironment() error

	// PrintDescription will print the stored description templated to the given writer.
	// If nill is passed, this function will exit early.
	PrintDescription(w io.Writer) error
}

type set struct {
	description string
	flags       []Flag
	envs        []interface{}
}

func NewFlagSet() FlagSet {
	s := &set{}
	return s
}

func (s *set) PrintDescription(w io.Writer) error {
	if w == nil {
		return nil
	}
	t, err := template.New("termi.usage").Parse(s.description)
	if err != nil {
		return err
	}
	values := map[string]string{}
	for _, e := range s.envs {
		results, err := EnvironmentDescription(e)
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
	for index := 0; index < len(args); index++ {
		if args[index] == "--" {
			return append(remainder, args[index+1:]...), nil
		}
		var (
			name  = args[index]
			value string
		)
		if index+1 < len(args) {
			value = args[index+1]
		}
		count, err := s.processFlag(name, value)
		if err != nil {
			return remainder, err
		}
		if count == 0 {
			remainder = append(remainder, args[index])
		}
		index += count
	}
	return remainder, nil
}

func (s *set) processFlag(name, arg string) (int, error) {
	if !strings.HasPrefix(name, "-") {
		return 0, nil
	}
	count := 0
	for _, flag := range s.flags {
		if !flag.IsFlag(name) {
			continue
		}
		// Boolean is a special case since it most often doesn't require an argument
		// when supplying flags and can be assume to set the variable to true by design.
		switch flag.(type) {
		case *Boolean:
			err := flag.Set(arg)
			switch err {
			case ErrorMissingBoolean:
				flag.Set("true")
				count = 1
			case nil:
				count = 2
			default:
				return 0, err
			}
		default:
			if err := flag.Set(arg); err != nil {
				return 0, err
			}
			count = 2
		}
	}
	return count, nil
}

func (s *set) Register(f Flag) FlagSet {
	s.flags = append(s.flags, f)
	return s
}

func (s *set) SetDescription(description string) {
	s.description = description
}

func (s *set) SetEnvironment(variable interface{}) FlagSet {
	s.envs = append(s.envs, variable)
	return s
}
