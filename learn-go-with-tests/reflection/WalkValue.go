package reflection

import "reflect"

func WalkValue(value reflect.Value, callback func(input string)) {
	WalkProperties(value.Interface(), callback)
}
