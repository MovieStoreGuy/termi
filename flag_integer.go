package termi

import (
	"errors"
	"reflect"
	"strconv"
)

type Integer struct {
	*basic
}

func NewInteger() Flag {
	return &Integer{&basic{names: map[string]bool{}}}
}

func (n *Integer) IsFlag(name string) bool {
	if n.basic == nil {
		n.basic = new(basic)
	}
	return n.isFlag(name)
}

func (s *Integer) Set(value string) error {
	if s.basic == nil {
		s.basic = new(basic)
	}
	if !s.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	v, err := strconv.Atoi(value)
	s.value.Set(reflect.ValueOf(v))
	return err
}

func (n *Integer) SetValue(value interface{}) Flag {
	if n.basic == nil {
		n.basic = new(basic)
	}
	n.value = reflect.Indirect(reflect.ValueOf(value))
	return n
}

func (n *Integer) SetDescription(description string) Flag {
	if n.basic == nil {
		n.basic = new(basic)
	}
	n.description = description
	return n
}

func (n *Integer) SetName(name string) Flag {
	if n.basic == nil {
		n.basic = new(basic)
	}
	n.setName(name)
	return n
}

func (n *Integer) String() string {
	return n.basic.String("string")
}
