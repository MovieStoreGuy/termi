package termi

import (
	"fmt"
	"reflect"
	"time"
)

// NewFlag is a generic function that will return the correct concrete flag type
// so that the value will be correctly updated.
func NewFlag(value interface{}) (Flag, error) {
	var f Flag
	switch value.(type) {
	case *int:
		f = NewInteger().SetValue(value)
	case *string:
		f = NewString().SetValue(value)
	case *uint:
		f = NewUsignedInteger().SetValue(value)
	case *bool:
		f = NewBoolean().SetValue(value)
	case *time.Duration:
		f = NewDuration().SetValue(value)
	default:
		return f, fmt.Errorf("unable to convert value (%v) to concrete flag type", reflect.ValueOf(value).Type())
	}
	return f, nil
}

// Must allows for wrapping the generic flag parser to provide simple usage for the lazy developer
// if the err is not nil, it will panic.
func Must(f Flag, err error) Flag {
	if err != nil {
		panic(err)
	}
	return f
}
