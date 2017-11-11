package take

import "reflect"

func String(items interface{}, fn func(i int) string) []string {
	rv := reflect.ValueOf(items)
	size := rv.Len()
	result := make([]string, size)

	for i := 0; i < size; i++ {
		result[i] = fn(i)
	}
	return result
}

func Int64(items interface{}, fn func(i int) int64) []int64 {
	rv := reflect.ValueOf(items)
	size := rv.Len()
	result := make([]int64, size)

	for i := 0; i < size; i++ {
		result[i] = fn(i)
	}
	return result
}

func Int(items interface{}, fn func(i int) int) []int {
	rv := reflect.ValueOf(items)
	size := rv.Len()
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = fn(i)
	}
	return result
}
