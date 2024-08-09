package mediator

import "reflect"

// TypeOf returns the name of the interface in type string
func TypeOf(i interface{}) string {
	if i == nil {
		return ""
	}

	return reflect.TypeOf(i).Elem().Name()
}
