package slice

import (
	"reflect"
)

func SliceContains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

func SliceIndexOf[T any](data []T, value T) int {
	for i, v := range data {
		if reflect.DeepEqual(v, value) {
			return i
		}
	}
	return -1
}
