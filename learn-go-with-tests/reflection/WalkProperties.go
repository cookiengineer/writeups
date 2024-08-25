package reflection

import "reflect"

func toReflectValue(raw any) reflect.Value {

	value := reflect.ValueOf(raw)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	return value

}

func WalkProperties(raw any, callback func(input string)) {

	value := toReflectValue(raw)

	switch value.Kind() {
	case reflect.String:

		callback(value.String())

	case reflect.Struct:

		for i := 0; i < value.NumField(); i++ {
			WalkValue(value.Field(i), callback)
		}

	case reflect.Slice, reflect.Array:

		for i := 0; i < value.Len(); i++ {
			WalkValue(value.Index(i), callback)
		}

	case reflect.Map:

		for _, key := range value.MapKeys() {
			WalkValue(value.MapIndex(key), callback)
		}

	case reflect.Chan:

		// Same loop as below, but written differently
		// for val, ok := value.Recv(); ok; val, ok = value.Recv() {
		// 	WalkValue(val, callback)
		// }

		val, ok := value.Recv()

		for ok {

			WalkValue(val, callback)
			val, ok = value.Recv()

		}

	case reflect.Func:

		result := value.Call(nil)

		for _, val := range result {
			WalkValue(val, callback)
		}

	}

}

