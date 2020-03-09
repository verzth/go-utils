package utils

func Substring(text string, indexs... int) string {
	runes := []rune(text)
	if len(indexs) > 0 {
		if indexs[0] < 0 || indexs[0] > len(runes) {
			return ""
		}
		if len(indexs) > 1 {
			if indexs[0] >= indexs[1] {
				return ""
			}
			if indexs[1] <= len(runes) {
				return string(runes[indexs[0]:indexs[1]])
			}
		}
		return string(runes[indexs[0]:])
	}else{
		return text
	}
}

func AddTo(collections interface{}, collection interface{}, indexs... int) {
	switch collections.(type) {
	case *[]interface{}: addInterfaceAt(collections.(*[]interface{}), collection, indexs)
	case *[]map[interface{}]interface{}: addMapAt(collections.(*[]map[interface{}]interface{}), collection.(map[interface{}]interface{}), indexs)
	case *[]int: addIntAt(collections.(*[]int), collection.(int), indexs)
	case *[]int8: addInt8At(collections.(*[]int8), collection.(int8), indexs)
	case *[]int16: addInt16At(collections.(*[]int16), collection.(int16), indexs)
	case *[]int32: addInt32At(collections.(*[]int32), collection.(int32), indexs)
	case *[]int64: addInt64At(collections.(*[]int64), collection.(int64), indexs)
	case *[]uint: addUintAt(collections.(*[]uint), collection.(uint), indexs)
	case *[]uint8: addUint8At(collections.(*[]uint8), collection.(uint8), indexs)
	case *[]uint16: addUint16At(collections.(*[]uint16), collection.(uint16), indexs)
	case *[]uint32: addUint32At(collections.(*[]uint32), collection.(uint32), indexs)
	case *[]uint64: addUint64At(collections.(*[]uint64), collection.(uint64), indexs)
	case *[]float32: addFloat32At(collections.(*[]float32), collection.(float32), indexs)
	case *[]float64: addFloat64At(collections.(*[]float64), collection.(float64), indexs)
	}
}

func RemoveAt(collections interface{}, index int) {
	switch collections.(type) {
	case *[]map[interface{}]interface{}: removeMapAt(collections.(*[]map[interface{}]interface{}), index)
	case *[]int: removeIntAt(collections.(*[]int), index)
	case *[]int8: removeInt8At(collections.(*[]int8), index)
	case *[]int16: removeInt16At(collections.(*[]int16), index)
	case *[]int32: removeInt32At(collections.(*[]int32), index)
	case *[]int64: removeInt64At(collections.(*[]int64), index)
	case *[]uint: removeUintAt(collections.(*[]uint), index)
	case *[]uint8: removeUint8At(collections.(*[]uint8), index)
	case *[]uint16: removeUint16At(collections.(*[]uint16), index)
	case *[]uint32: removeUint32At(collections.(*[]uint32), index)
	case *[]uint64: removeUint64At(collections.(*[]uint64), index)
	case *[]float32: removeFloat32At(collections.(*[]float32), index)
	case *[]float64: removeFloat64At(collections.(*[]float64), index)
	case *[]string: removeStringAt(collections.(*[]string), index)
	case *[]interface{}: removeInterfaceAt(collections.(*[]interface{}), index)
	}
}

func FileMove(src string, dst string, file... string)  {
	pCount := len(file)
	if pCount == 0 {
		movePath(src,dst)
	}else if pCount == 1 {
		moveFile(src,dst,file[0])
	}else{
		moveFileRename(src,dst,file[0],file[1])
	}
}

func Uniquify(collections interface{})  {
	switch collections.(type) {
	case *[]int: removeIntDuplicate(collections.(*[]int))
	case *[]int8: removeInt8Duplicate(collections.(*[]int8))
	case *[]int16: removeInt16Duplicate(collections.(*[]int16))
	case *[]int32: removeInt32Duplicate(collections.(*[]int32))
	case *[]int64: removeInt64Duplicate(collections.(*[]int64))
	case *[]uint: removeUintDuplicate(collections.(*[]uint))
	case *[]uint8: removeUint8Duplicate(collections.(*[]uint8))
	case *[]uint16: removeUint16Duplicate(collections.(*[]uint16))
	case *[]uint32: removeUint32Duplicate(collections.(*[]uint32))
	case *[]uint64: removeUint64Duplicate(collections.(*[]uint64))
	case *[]float32: removeFloat32Duplicate(collections.(*[]float32))
	case *[]float64: removeFloat64Duplicate(collections.(*[]float64))
	case *[]interface{}: removeInterfaceDuplicate(collections.(*[]interface{}))
	case *[]string: removeStringDuplicate(collections.(*[]string))
	}
}