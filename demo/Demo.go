package main

import (
	"fmt"
	"git.verzth.work/go/utils"
	"os"
)

func main() {
	fmt.Printf("RandString: %v\n", utils.RandString(20))
	fmt.Printf("RandStringWithCharset: %v\n", utils.RandStringWithCharset(20, "abc123efg456"))
	fmt.Printf("RandStringLower: %v\n", utils.RandStringLower(20))
	fmt.Printf("RandomHex: %v\n", utils.RandomHex(16))
	fmt.Printf("RandomHexUpper: %v\n", utils.RandomHexUpper(16))
	ranString := utils.RandStringUpper(10)
	fmt.Printf("RandStringUpper: %v\n", ranString) // 'ea10a9c919'

	fmt.Printf("Substring: %v\n", utils.Substring(ranString, 2))      // '10a9c919'
	fmt.Printf("Substring: %v\n", utils.Substring(ranString, 1, 8))   // 'a10a9c9'
	fmt.Printf("Substring: %v\n", utils.Substring(ranString, 6, 2))   // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString, 10, 15)) // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString, -1, 5))  // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString, 5, 10))  // '9c919'

	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array Before: %v\n", arr)
	utils.Slice.RemoveAt(&arr, 3) // Remove slice at index
	fmt.Printf("Array After 1: %v\n", arr)
	utils.Slice.AddTo(&arr, 4, 3) // Add data to index 3
	fmt.Printf("Array After 2: %v\n", arr)
	utils.Slice.AddTo(&arr, 25, 1) // Add data to index 1
	fmt.Printf("Array After 3: %v\n", arr)
	utils.Slice.AddTo(&arr, 100) // Add data without specify index, will be added to last index
	fmt.Printf("Array After 4: %v\n", arr)
	utils.Slice.AddTo(&arr, -5) // Add data with minus index, will be added to first index which is treated as 0
	fmt.Printf("Array After 5: %v\n", arr)
	i := utils.Slice.Exist(&arr, 9)
	fmt.Println("Exist", i)
	j := utils.Slice.Exist(&arr, 100)
	fmt.Println("Exist", j)

	utils.FileMove("/root/project/filename", "/root/project/newname", os.ModePerm)                             // Move file from path to path
	utils.FileMove("/root/project/oldfolder/", "/root/project/newfolder/", os.ModePerm, "filename")            // Move file from to new location with same name
	utils.FileMove("/root/project/oldfolder/", "/root/project/newfolder/", os.ModePerm, "filename", "newname") // Move file from to new location with new name

	arrDuplicate := []int{1, 0, 0, 2, 3, 2, 4, 5, 6, 7, 4, 4, 7, 4, 7, 7, 7, 15}
	fmt.Println(arrDuplicate)
	utils.Slice.Uniquify(&arrDuplicate)
	fmt.Println(arrDuplicate)

	objArr := []map[string]int{
		{"id": 4, "num": 2}, {"id": 3, "num": 9}, {"id": 0, "num": 1}, {"id": 8, "num": 2},
	}
	utils.Slice.Where(&objArr, func(i int) bool {
		return objArr[i]["num"] == 2
	})
	fmt.Printf("Where: %v\n", objArr) // Get all value in slice with condition

	fmt.Printf("First: %v\n", utils.Slice.First(objArr))               // Get first value in slice
	fmt.Printf("IndexOf: %v\n", utils.Slice.IndexOf(arrDuplicate, 15)) // Find first index of value in slice
	fmt.Printf("FirstWhere: %v\n", utils.Slice.FirstWhere(objArr, func(i int) bool {
		return objArr[i]["num"] == 9
	})) // Get first value in slice with given condition
	fmt.Printf("IndexWhere: %v\n", utils.Slice.IndexWhere(objArr, func(i int) bool {
		return objArr[i]["id"] == 0
	})) // Get first index in slice with given condition

	fmt.Printf("Last: %v\n", utils.Slice.Last(objArr))                        // Get first value in slice
	fmt.Printf("LastIndexOf: %v\n", utils.Slice.LastIndexOf(arrDuplicate, 0)) // Find last index of value in slice
	fmt.Printf("LastWhere: %v\n", utils.Slice.LastWhere(objArr, func(i int) bool {
		return objArr[i]["id"] == 4
	})) // Get last value in slice with given condition
	fmt.Printf("LastIndexWhere: %v\n", utils.Slice.LastIndexWhere(objArr, func(i int) bool {
		return objArr[i]["num"] == 2
	})) // Get last index in slice with given condition
}
