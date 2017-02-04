package event

import (
	"reflect"
)

func simpleName(v interface{}) string {

	switch item := v.(type) {
	case reflect.Type:
		{
			return item.String()

		}
	default:
		return reflect.TypeOf(v).String()
	}
}
