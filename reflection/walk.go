package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	// run recursively if meet with Slice
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	// in this case, only 1 string in input
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
