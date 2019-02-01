package termi

import (
	"errors"
	"reflect"
	"time"
)

type Duration struct {
	*basic
}

func NewDuration() Flag {
	return &Duration{&basic{names: map[string]bool{}}}
}

func (d *Duration) IsFlag(name string) bool {
	if d.basic == nil {
		d.basic = new(basic)
	}
	return d.isFlag(name)
}

func (d *Duration) Set(value string) error {
	if d.basic == nil {
		d.basic = new(basic)
	}
	if !d.value.IsValid() {
		return errors.New("unable to update value if value is not valid")
	}
	v, err := time.ParseDuration(value)
	d.value.Set(reflect.ValueOf(v))
	return err
}

func (d *Duration) SetValue(value interface{}) Flag {
	if d.basic == nil {
		d.basic = new(basic)
	}
	d.value = reflect.Indirect(reflect.ValueOf(value))
	return d
}

func (d *Duration) SetDescription(description string) Flag {
	if d.basic == nil {
		d.basic = new(basic)
	}
	d.description = description
	return d
}

func (d *Duration) SetName(name string) Flag {
	if d.basic == nil {
		d.basic = new(basic)
	}
	d.setName(name)
	return d
}

func (d *Duration) String() string {
	return d.basic.String("duration")
}
