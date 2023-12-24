package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	// getField can be use for both reflect.Value.Field(i) or reflect.Value.Index(i)
	var getField func(int) reflect.Value

	switch val.Kind() {
	// run recursively if meet with Slice
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	// in this case, only 1 string in input
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
