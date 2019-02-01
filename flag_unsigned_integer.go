package termi

import (
	"errors"
	"reflect"
	"strconv"
)

type UnsignedInteger struct {
	*basic
}

func NewUsignedInteger() Flag {
	return &UnsignedInteger{&basic{names: map[string]bool{}}}
}

func (u *UnsignedInteger) IsFlag(name string) bool {
	if u.basic == nil {
		u.basic = new(basic)
	}
	return u.isFlag(name)
}

func (u *UnsignedInteger) Set(value string) error {
	if u.basic == nil {
		u.basic = new(basic)
	}
	if !u.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	v, err := strconv.ParseUint(value, 10, strconv.IntSize)
	u.value.Set(reflect.ValueOf(v))
	return err
}

func (u *UnsignedInteger) SetValue(value interface{}) Flag {
	if u.basic == nil {
		u.basic = new(basic)
	}
	u.value = reflect.Indirect(reflect.ValueOf(value))
	return u
}

func (u *UnsignedInteger) SetDescription(description string) Flag {
	if u.basic == nil {
		u.basic = new(basic)
	}
	u.description = description
	return u
}

func (u *UnsignedInteger) SetName(name string) Flag {
	if u.basic == nil {
		u.basic = new(basic)
	}
	u.setName(name)
	return u
}

func (u *UnsignedInteger) String() string {
	return u.basic.String("unsigned integer")
}
