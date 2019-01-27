package termi

import (
	"errors"
	"reflect"
)

type String struct {
	*basic
}

func NewString() Flag {
	return &String{new(basic)}
}

func (s *String) IsFlag(name string) bool {
	if s.basic == nil {
		s.basic = new(basic)
	}
	return s.isFlag(name)
}

func (s *String) Set(value string) error {
	if s.basic == nil {
		s.basic = new(basic)
	}
	if !s.value.IsValid() {
		return errors.New("unable to set value due to no value being previously defined")
	}
	s.value.Set(reflect.ValueOf(value))
	return nil
}

func (s *String) SetValue(value interface{}) Flag {
	if s.basic == nil {
		s.basic = new(basic)
	}
	s.value = reflect.Indirect(reflect.ValueOf(value))
	return s
}

func (s *String) SetDescription(description string) Flag {
	if s.basic == nil {
		s.basic = new(basic)
	}
	s.description = description
	return s
}

func (s *String) SetName(name string) Flag {
	if s.basic == nil {
		s.basic = new(basic)
	}
	s.setName(name)
	return s
}

func (s *String) String() string {
	return s.basic.String("string")
}
