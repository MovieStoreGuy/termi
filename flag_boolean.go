package termi

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	// ErrorMissingBoolean allows for checking if we need to stumble when parsing flag arguments
	ErrorMissingBoolean = errors.New("non boolean value parsed")
)

type Boolean struct {
	*basic
}

func NewBoolean() Flag {
	return &Boolean{new(basic)}
}

func (b *Boolean) IsFlag(name string) bool {
	if b.basic == nil {
		b.basic = new(basic)
	}
	return b.isFlag(name)
}

func (b *Boolean) Set(value string) error {
	if b.basic == nil {
		b.basic = new(basic)
	}
	if !b.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	bool, err := strconv.ParseBool(value)
	if err != nil {
		return ErrorMissingBoolean
	}
	b.value.Set(reflect.ValueOf(bool))
	return err
}

func (b *Boolean) SetValue(value interface{}) Flag {
	if b.basic == nil {
		b.basic = new(basic)
	}
	b.value = reflect.Indirect(reflect.ValueOf(value))
	return b
}

func (b *Boolean) SetDescription(description string) Flag {
	if b.basic == nil {
		b.basic = new(basic)
	}
	b.description = description
	return b
}

func (b *Boolean) SetName(name string) Flag {
	if b.basic == nil {
		b.basic = new(basic)
	}
	b.setName(name)
	return b
}

func (b *Boolean) String() string {
	return b.basic.String("boolean")
}
