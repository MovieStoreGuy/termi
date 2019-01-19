package termi

import (
	"fmt"
	"reflect"
)

// EnvironmentDescription fetches struct tags being used in order to print out to
// flag.Usage or other meaningful parts of the application
func EnvironmentDescription(obj interface{}) (map[string]string, error) {
	v := reflect.ValueOf(obj)
	if !(v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct) || v.Kind() != reflect.Struct {
		return nil, ErrInvalidType
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	entries := make(map[string]string, 0)
	for i := 0; i < v.NumField(); i++ {
		name, exist := v.Type().Field(i).Tag.Lookup(EnvironmentTag)
		if !exist {
			continue
		}
		if description, exist := v.Type().Field(i).Tag.Lookup(DescriptionTag); exist {
			entries[name] = description
		} else {
			entries[name] = fmt.Sprintf(DefaultDescription, v.Field(i).Kind())
		}
	}
	return entries, nil
}