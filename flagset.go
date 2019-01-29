package termi

import (
	"errors"
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
	var (
		index     int
		remainder []string
	)
	for {
	start:
		if index >= len(args) {
			break
		}
		if args[index] == "--" {
			return append(remainder, args[index+1:]...), nil
		}
		for _, flag := range s.flags {
			// Needed to ensure I don't accidentally consume a non flag value
			if strings.HasPrefix(args[index], "-") && flag.IsFlag(args[index]) {
				consumedArgs := 1
				switch flag.(type) {
				case *Boolean:
					switch {
					case index+1 >= len(args):
						if err := flag.Set("true"); err != nil {
							panic(err)
						}
					default:
						if err := flag.Set(args[index+1]); err != nil {
							switch err {
							case ErrorMissingBoolean:
								if err := flag.Set("true"); err != nil {
									panic(err)
								}
							}
						} else {
							consumedArgs++
						}
					}
				default:
					if index+1 >= len(args) {
						return remainder, errors.New("missing required argument on end of list")
					}
					if err := flag.Set(args[index+1]); err != nil {
						return remainder, err
					}
					consumedArgs++
				}
				index += consumedArgs
				goto start
			}
		}
		remainder = append(remainder, args[index])
		index++
	}
	return remainder, nil
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
