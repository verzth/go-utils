package utils

import (
	"reflect"
)

var Slice slice

type slice struct {}

func (s slice) AddTo(collections interface{}, collection interface{}, indexs... int) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		if len(indexs) == 0 {
			temp := reflect.Append(indirect.Elem().Slice(0, indirect.Elem().Len()), reflect.ValueOf(collection))
			indirect.Elem().Set(temp)
		}else{
			if indexs[0] < indirect.Elem().Len() && indexs[0] >= 0 {
				tail := reflect.MakeSlice(indirect.Elem().Type(), indirect.Elem().Slice(indexs[0], indirect.Elem().Len()).Len(), indirect.Elem().Cap())
				reflect.Copy(tail, indirect.Elem().Slice(indexs[0], indirect.Elem().Len()))
				temp := reflect.Append(indirect.Elem().Slice(0, indexs[0]), reflect.ValueOf(collection))
				temp = reflect.AppendSlice(temp, tail)
				indirect.Elem().Set(temp)
			}
		}
	}
}

func (s slice) IndexOf(collections interface{}, val interface{}) (index int) {
	index = -1
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		v := reflect.ValueOf(val)
		for i:=0; i<el.Len(); i++{
			if el.Index(i).Interface() == v.Interface() {
				index = i
				return
			}
		}
	}
	return
}

func (s slice) Where(collections interface{}, fn func(int) bool) (cols interface{}) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		newCols := reflect.MakeSlice(indirect.Type(), 0, 0)
		for i:=0; i<el.Len(); i++{
			if fn(i) {
				newCols = reflect.Append(newCols, reflect.ValueOf(el.Index(i).Interface()))
			}
		}
		cols = newCols
	}
	return
}

func (s slice) IndexWhere(collections interface{}, fn func(int) bool) (index int) {
	index = -1
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		for i:=0; i<el.Len(); i++{
			if fn(i) {
				index = i
				return
			}
		}
	}
	return
}

func (s slice) LastIndexOf(collections interface{}, val interface{}) (index int) {
	index = -1
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice {
			el = indirect
		}else{
			return
		}

		v := reflect.ValueOf(val)
		for i:=0; i<el.Len(); i++{
			if el.Index(i).Interface() == v.Interface() {
				index = i
			}
		}
	}
	return
}

func (s slice) LastIndexWhere(collections interface{}, fn func(int) bool) (index int) {
	index = -1
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		for i:=0; i<el.Len(); i++{
			if fn(i) {
				index = i
			}
		}
	}
	return
}

func (s slice) RemoveAt(collections interface{}, index int) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		if index < indirect.Elem().Len() && index >= 0 {
			temp := reflect.AppendSlice(indirect.Elem().Slice(0,index), indirect.Elem().Slice(index+1, indirect.Elem().Len()))
			indirect.Elem().Set(temp)
		}
	}
}

func (s slice) RemoveIn(collections interface{}, indexes []int) {
	for _, index := range indexes {
		s.RemoveAt(collections, index)
	}
}

// Uniquify(collections): Uniquify slices value
func (s slice) Uniquify(collections interface{})  {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i:=0; i<indirect.Elem().Len(); i++{
			v1 := indirect.Elem().Index(i)
			for j:=i+1; j<indirect.Elem().Len(); j++{
				v2 := indirect.Elem().Index(j)
				if v1.Interface() == v2.Interface() {
					s.RemoveAt(collections, j)
					j--
				}
			}
		}
	}
}

func (s slice) IndexesOf(collections interface{}, val interface{}) (indexes []int) {
	indexes = []int{}
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice {
			el = indirect
		}else{
			return
		}

		v := reflect.ValueOf(val)
		for i:=0; i<el.Len(); i++{
			if el.Index(i).Interface() == v.Interface() {
				indexes = append(indexes,i)
			}
		}
	}
	return
}

func (s slice) First(collections interface{}) (col interface{}) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		if el.Len() > 0 {
			return el.Index(0)
		}
	}
	return
}

func (s slice) Last(collections interface{}) (col interface{}) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		if el.Len() > 0 {
			return el.Index(el.Len()-1)
		}
	}
	return
}

func (s slice) FirstWhere(collections interface{}, fn func(int) bool) (col interface{}) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		for i:=0; i<el.Len(); i++{
			if fn(i) {
				return el.Index(i)
			}
		}
	}
	return
}

func (s slice) LastWhere(collections interface{}, fn func(int) bool) (col interface{}) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() {
		var el reflect.Value
		if indirect.Kind() == reflect.Ptr {
			if indirect.Elem().Kind() == reflect.Slice {
				el = indirect.Elem()
			}else{
				return
			}
		}else if indirect.Kind() == reflect.Slice{
			el = indirect
		}else{
			return
		}

		for i:=0; i<el.Len(); i++{
			if fn(i) {
				col = el.Index(i)
			}
		}
	}
	return
}

func (s slice) Exist(collections interface{}, val interface{}) (state bool) {
	i := s.IndexOf(collections, val)
	state = i!=-1
	return
}