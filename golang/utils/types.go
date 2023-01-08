package utils

import "reflect"

func IsSlice(x interface{}) bool {
	objectType := reflect.TypeOf(x).Kind()
	return objectType == reflect.Slice
}

func InterfaceToInt(x interface{}) int {
	return x.(int)
}

func InterfaceToSlice(x interface{}) []interface{} {
	return x.([]interface{})
}
