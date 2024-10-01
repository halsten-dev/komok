package slice

import (
	"reflect"
)

func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func IndexOf[T any](data []T, value T) int {
	for i, v := range data {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}
	return -1
}
