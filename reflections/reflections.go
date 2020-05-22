package reflections

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	//numberOfValues := 0
	//var getField func(int) reflect.Value
	walkValud := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		//numberOfValues = val.NumField()
		//getField = val.Field
		for i := 0; i < val.NumField(); i++ {
			walkValud(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		//numberOfValues = val.Len()
		//getField = val.Index
		for i := 0; i < val.Len(); i++ {
			walkValud(val.Index(i))
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walkValud(val.MapIndex(k))
		}
	}
	/*
		for i := 0; i < numberOfValues; i++ {
			walk(getField(i).Interface(), fn)
		}
	*/
}
