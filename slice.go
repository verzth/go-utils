package utils

import (
	"reflect"
)

func Prepend[T any](collections []T, cols ...T) (res []T) {
	res = append(cols, collections...)
	return
}

func AddTo[T any](collections *[]T, collection T, index int) {
	temp := make([]T, 0)
	if index < len(*collections) && index >= 0 {
		tail := (*collections)[index:]
		temp = (*collections)[:index]
		temp = append(temp, collection)
		temp = append(temp, tail...)
	} else if index < 0 {
		temp = Prepend(*collections, collection)
	} else {
		temp = append(*collections, collection)
	}
	collections = &temp
}

func Where[T any](collections *[]T, fn func(T) bool) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			} else {
				return
			}
		} else if indirect.Kind() == reflect.Slice {
			el = indirect
		} else {
			return
		}

		cols := reflect.MakeSlice(el.Type(), 0, 0)
		for i := 0; i < el.Len(); i++ {
			if fn((*collections)[i]) {
				cols = reflect.Append(cols, reflect.ValueOf(el.Index(i).Interface()))
			}
		}

		indirect.Elem().Set(cols)
	}
}

func IndexWhere[T any](collections []T, fn func(T) bool) (index int) {
	index = -1
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			} else {
				return
			}
		} else if indirect.Kind() == reflect.Slice {
			el = indirect
		} else {
			return
		}

		for i := 0; i < el.Len(); i++ {
			if fn(collections[i]) {
				index = i
				return
			}
		}
	}
	return
}

func IndexesWhere[T any](collections []T, fn func(T) bool) (indexes []int) {
	for i := range collections {
		if fn(collections[i]) {
			indexes = append(indexes, i)
			return
		}
	}
	return
}

func LastIndexWhere[T any](collections []T, fn func(T) bool) (index int) {
	index = -1
	for i := range collections {
		if fn(collections[i]) {
			index = i
		}
	}
	return
}

func RemoveAt[T any](collections *[]T, index int) {
	*collections = append((*collections)[:index], (*collections)[index+1:]...)
}

func RemoveIn[T any](collections *[]T, indexes []int) {
	for _, index := range indexes {
		RemoveAt[T](collections, index)
	}
}

func RemoveWhere[T any](collections *[]T, fn func(T) bool) {
	ixs := IndexesWhere[T](*collections, fn)
	RemoveIn[T](collections, ixs)
}

// Distinct (collections): Distinct slices value
func Distinct[T any](collections *[]T) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i := 0; i < indirect.Elem().Len(); i++ {
			v1 := indirect.Elem().Index(i)
			for j := i + 1; j < indirect.Elem().Len(); j++ {
				v2 := indirect.Elem().Index(j)
				if v1.Interface() == v2.Interface() {
					RemoveAt[T](collections, j)
					j--
				}
			}
		}
	}
}

func IndexesOf[T any](collections []T, val T) (indexes []int) {
	indexes = []int{}
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			} else {
				return
			}
		} else if indirect.Kind() == reflect.Slice {
			el = indirect
		} else {
			return
		}

		v := reflect.ValueOf(val)
		for i := 0; i < el.Len(); i++ {
			if el.Index(i).Interface() == v.Interface() {
				indexes = append(indexes, i)
			}
		}
	}
	return
}

func First[T any](collections []T) (col T) {
	if len(collections) > 0 {
		col = collections[0]
	}
	return
}

func Last[T any](collections []T) (col T) {
	if len(collections) > 0 {
		col = collections[len(collections)-1]
	}
	return
}

func FirstWhere[T any](collections []T, fn func(T) bool) (col T) {
	for _, c := range collections {
		if fn(c) {
			return c
		}
	}
	return
}

func LastWhere[T any](collections []T, fn func(T) bool) (col T) {
	for _, c := range collections {
		if fn(c) {
			col = c
		}
	}
	return
}

func Exist[T any](collections []T, fn func(T) bool) (state bool) {
	i := FirstWhere[T](collections, fn)
	state = i != nil
	return
}
