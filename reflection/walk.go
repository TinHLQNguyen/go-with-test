package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	// run recursively if meet with Slice
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	// will iterate until ok is false
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	// in this ease, only 1 string in input
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
