package termi

import (
	"fmt"
	"reflect"
	"strings"
)

// Flag is a neat little abstraction that allows for simplified
// Command line arguments
type Flag interface {
	// SetDescription allows you to set the usage of the flag
	SetDescription(description string) Flag

	// SetValue allows you to pass the reference to the value you wish to
	// update with when parse is called.
	SetValue(value interface{}) Flag

	// SetName allows you to define what argument is associated with this flag.
	// Can be called multiple times
	SetName(name string) Flag

	// Set is called when parsing the argument within the FlagSet
	Set(value string) error

	// IsFlag allows the flag set to parse the command line arg
	IsFlag(name string) bool
}

type basic struct {
	description string
	names       map[string]bool
	value       reflect.Value
}

func (b *basic) isFlag(name string) bool {
	if b.names == nil {
		b.names = make(map[string]bool, 0)
	}
	return b.names[strings.Trim(name, "-")]
}

func (b *basic) String(t interface{}) string {
	toArray := func(m map[string]bool) []string {
		values := []string{}
		for name := range m {
			values = append(values, name)
		}
		return values
	}
	return fmt.Sprintf("%v (type:%v) -- %s", toArray(b.names), t, b.description)
}

func (b *basic) setName(name string) {
	if b.names == nil {
		b.names = make(map[string]bool, 0)
	}
	b.names[strings.Trim(name, "-")] = true
}
