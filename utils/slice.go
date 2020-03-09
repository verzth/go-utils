package utils

import "fmt"

func addInterfaceAt(collections *[]interface{},collection interface{}, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addMapAt(collections *[]map[interface{}]interface{}, collection map[interface{}]interface{}, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addIntAt(collections *[]int, collection int, index []int)  {
	temp := *collections
	if len(index) == 0 {
		temp = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			head := append([]int{},temp[:index[0]]...)
			tail := append([]int{},temp[index[0]:]...)
			temp = append(head, collection)
			temp = append(temp, tail...)
		}else if index[0] >= len(temp){
			temp = append(temp, collection)
		}else if index[0] < 0 {
			temp = append([]int{collection}, temp...)
		}
	}
	*collections = temp
}

func addInt8At(collections *[]int8, collection int8, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addInt16At(collections *[]int16, collection int16, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addInt32At(collections *[]int32, collection int32, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addInt64At(collections *[]int64, collection int64, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addUintAt(collections *[]uint, collection uint, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addUint8At(collections *[]uint8, collection uint8, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addUint16At(collections *[]uint16, collection uint16, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addUint32At(collections *[]uint32, collection uint32, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addUint64At(collections *[]uint64, collection uint64, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addFloat32At(collections *[]float32, collection float32, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func addFloat64At(collections *[]float64, collection float64, index []int)  {
	temp := *collections
	if len(index) == 0 {
		*collections = append(temp, collection)
	}else{
		if index[0] < len(temp) && index[0] >= 0 {
			*collections = append(temp[:index[0]], collection)
			*collections = append(*collections, temp[index[0]:]...)
		}
	}
}

func removeMapAt(collections *[]map[interface{}]interface{}, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeIntAt(collections *[]int, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeInt8At(collections *[]int8, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeInt16At(collections *[]int16, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeInt32At(collections *[]int32, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeInt64At(collections *[]int64, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeUintAt(collections *[]uint, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeUint8At(collections *[]uint8, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeUint16At(collections *[]uint16, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeUint32At(collections *[]uint32, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeUint64At(collections *[]uint64, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeFloat32At(collections *[]float32, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeFloat64At(collections *[]float64, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeStringAt(collections *[]string, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeInterfaceAt(collections *[]interface{}, index int){
	temp := *collections
	if index < len(temp) && index >= 0 {
		*collections = append(temp[:index], temp[index+1:]...)
	}
}

func removeIntDuplicate(collections *[]int)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeInt8Duplicate(collections *[]int8)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeInt16Duplicate(collections *[]int16)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeInt32Duplicate(collections *[]int32)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeInt64Duplicate(collections *[]int64)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeUintDuplicate(collections *[]uint)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeUint8Duplicate(collections *[]uint8)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeUint16Duplicate(collections *[]uint16)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeUint32Duplicate(collections *[]uint32)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeUint64Duplicate(collections *[]uint64)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeFloat32Duplicate(collections *[]float32)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeFloat64Duplicate(collections *[]float64)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeStringDuplicate(collections *[]string)  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}

func removeInterfaceDuplicate(collections *[]interface{})  {
	temp := *collections
	if len(temp)>1 {
		for i:=0; i<len(temp); i++ {
			for j:=i+1; j<len(temp); j++ {
				fmt.Println(temp,i,temp[i],j,temp[j])
				if temp[i] == temp[j] {
					RemoveAt(&temp, j)
					j--
				}
			}
		}
		*collections = temp
	}
}