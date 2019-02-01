package termi

import (
	"fmt"
	"reflect"
	"strings"
)

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
	var (
		defaultValue interface{} = "unset"
	)
	if b.value.IsValid() {
		defaultValue = b.value.Interface()
	}
	return fmt.Sprintf("%v\t(type:%v)\n└──\t%s (default: %v)", toArray(b.names), t, b.description, defaultValue)
}

func (b *basic) setName(name string) {
	if b.names == nil {
		b.names = make(map[string]bool, 0)
	}
	b.names[strings.Trim(name, "-")] = true
}
