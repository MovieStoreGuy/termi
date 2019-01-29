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

func (i *Integer) IsFlag(name string) bool {
	if i.basic == nil {
		i.basic = new(basic)
	}
	return i.isFlag(name)
}

func (i *Integer) Set(value string) error {
	if i.basic == nil {
		i.basic = new(basic)
	}
	if !i.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	v, err := strconv.Atoi(value)
	i.value.Set(reflect.ValueOf(v))
	return err
}

func (i *Integer) SetValue(value interface{}) Flag {
	if i.basic == nil {
		i.basic = new(basic)
	}
	i.value = reflect.Indirect(reflect.ValueOf(value))
	return i
}

func (i *Integer) SetDescription(description string) Flag {
	if i.basic == nil {
		i.basic = new(basic)
	}
	i.description = description
	return i
}

func (i *Integer) SetName(name string) Flag {
	if i.basic == nil {
		i.basic = new(basic)
	}
	i.setName(name)
	return i
}

func (i *Integer) String() string {
	return i.basic.String("string")
}
