package utils

import "reflect"

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

func (s slice) Find(collections interface{}, val interface{}) (indexes []int, state bool) {
	indexes = []int{}; state = false
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i:=0; i<indirect.Elem().Len(); i++{
			if indirect.Elem().Index(i).Interface() == val {
				indexes = append(indexes,i)
			}
		}
		if len(indexes) > 0 {
			state = true
		}
	}
	return
}

func (s slice) First(collections interface{}, val interface{}) (index int, state bool) {
	index = -1; state = false
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i:=0; i<indirect.Elem().Len(); i++{
			if indirect.Elem().Index(i).Interface() == val {
				index = i; state=true
				return
			}
		}
	}
	return
}

func (s slice) Last(collections interface{}, val interface{}) (index int, state bool) {
	index = -1; state = false
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i:=0; i<indirect.Elem().Len(); i++{
			if indirect.Elem().Index(i).Interface() == val {
				index = i; state=true
			}
		}
	}
	return
}

func (s slice) Exist(collections interface{}, val interface{}) (state bool) {
	_, state = s.First(collections, val)
	return
}