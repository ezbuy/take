package take

import "reflect"

func StringUnique(items interface{}, fn func(i int) string) []string {
	rv := reflect.ValueOf(items)
	size := rv.Len()

	if size < 512 {
		return stringUniqueFor(size, fn)
	}

	return stringUniqueMap(size, fn)
}

func stringUniqueMap(size int, fn func(i int) string) []string {
	result := make([]string, 0, size)
	checkMap := make(map[string]bool)

	for i := 0; i < size; i++ {
		item := fn(i)
		if _, ok := checkMap[item]; !ok {
			result = append(result, item)
			checkMap[item] = true
		}
	}
	return result
}

func stringUniqueFor(size int, fn func(i int) string) []string {
	result := make([]string, 0, size)

	for i := 0; i < size; i++ {
		item := fn(i)
		found := false
		for _, existingItem := range result {
			if item == existingItem {
				found = true
				break
			}
		}
		if !found {
			result = append(result, item)
		}
	}
	return result
}

func Int64Unique(items interface{}, fn func(i int) int64) []int64 {
	rv := reflect.ValueOf(items)
	size := rv.Len()

	if size < 512 {
		return int64UniqueFor(size, fn)
	}

	return int64UniqueMap(size, fn)
}

func int64UniqueMap(size int, fn func(i int) int64) []int64 {
	result := make([]int64, 0, size)
	checkMap := make(map[int64]bool)

	for i := 0; i < size; i++ {
		item := fn(i)
		if _, ok := checkMap[item]; !ok {
			result = append(result, item)
			checkMap[item] = true
		}
	}
	return result
}

func int64UniqueFor(size int, fn func(i int) int64) []int64 {
	result := make([]int64, 0, size)

	for i := 0; i < size; i++ {
		item := fn(i)
		found := false
		for _, existingItem := range result {
			if item == existingItem {
				found = true
				break
			}
		}
		if !found {
			result = append(result, item)
		}
	}
	return result
}

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
