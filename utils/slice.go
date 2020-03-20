package utils

import "reflect"

func AddTo(collections interface{}, collection interface{}, indexs... int) {
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

func RemoveAt(collections interface{}, index int) {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		if index < indirect.Elem().Len() && index >= 0 {
			temp := reflect.AppendSlice(indirect.Elem().Slice(0,index), indirect.Elem().Slice(index+1, indirect.Elem().Len()))
			indirect.Elem().Set(temp)
		}
	}
}

// Uniquify(collections): Uniquify slices value
func Uniquify(collections interface{})  {
	indirect := reflect.ValueOf(collections)
	if indirect.IsValid() && indirect.Elem().Kind() == reflect.Slice {
		for i:=0; i<indirect.Elem().Len(); i++{
			v1 := indirect.Elem().Index(i)
			for j:=i+1; j<indirect.Elem().Len(); j++{
				v2 := indirect.Elem().Index(j)
				if v1.Interface() == v2.Interface() {
					RemoveAt(collections, j)
					j--
				}
			}
		}
	}
}