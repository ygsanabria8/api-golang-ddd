package mediator

import "reflect"

// TypeOf returns the interface's name in string type
func TypeOf(i interface{}) string {
	if i == nil {
		return ""
	}

	return reflect.TypeOf(i).Elem().Name()
}
