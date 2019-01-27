package termi

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Flag is a neat little abstraction that allows for simplified
// Command line arguments
type Flag interface {
	SetDescription(description string) Flag

	SetValue(value interface{}) Flag

	SetName(name string) Flag

	Set(value string) error

	IsFlag(name string) bool
}

type Integer struct {
	description string
	names       map[string]bool
	value       reflect.Value
}

func NewInteger() Flag {
	return &Integer{names: map[string]bool{}}
}

func (n *Integer) IsFlag(name string) bool {
	if n.names == nil {
		n.names = make(map[string]bool, 0)
	}
	return n.names[strings.Trim(name, "-")]
}

func (n *Integer) Set(value string) error {
	if !n.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	v, err := strconv.Atoi(value)
	n.value.Set(reflect.ValueOf(v))
	return err
}

func (n *Integer) SetValue(value interface{}) Flag {
	n.value = reflect.Indirect(reflect.ValueOf(value))
	return n
}

func (n *Integer) SetDescription(description string) Flag {
	n.description = description
	return n
}

func (n *Integer) SetName(name string) Flag {
	if n.names == nil {
		n.names = make(map[string]bool, 0)
	}
	n.names[strings.Trim(name, "-")] = true
	return n
}

func (n *Integer) String() string {
	toArray := func(m map[string]bool) []string {
		values := []string{}
		for name := range m {
			values = append(values, name)
		}
		return values
	}
	return fmt.Sprintf("flags: %v (type:%v) -- %s", toArray(n.names), reflect.TypeOf(n).Elem(), n.description)
}

type String struct {
	description string
	names       map[string]bool
	value       reflect.Value
}

func NewString() Flag {
	return &String{names: map[string]bool{}}
}

func (n *String) IsFlag(name string) bool {
	if n.names == nil {
		n.names = make(map[string]bool, 0)
	}
	return n.names[strings.Trim(name, "-")]
}

func (n *String) Set(value string) error {
	if !n.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	n.value.Set(reflect.ValueOf(value))
	return nil
}

func (n *String) SetValue(value interface{}) Flag {
	n.value = reflect.Indirect(reflect.ValueOf(value))
	return n
}

func (n *String) SetDescription(description string) Flag {
	n.description = description
	return n
}

func (n *String) SetName(name string) Flag {
	if n.names == nil {
		n.names = make(map[string]bool, 0)
	}
	n.names[strings.Trim(name, "-")] = true
	return n
}

func (n *String) String() string {
	toArray := func(m map[string]bool) []string {
		values := []string{}
		for name := range m {
			values = append(values, name)
		}
		return values
	}
	return fmt.Sprintf("flags: %v (type:%v) -- %s", toArray(n.names), reflect.TypeOf(n).Elem(), n.description)
}
