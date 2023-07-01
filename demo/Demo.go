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
	utils.RemoveAt[int](&arr, 3) // Remove slice at index
	fmt.Printf("Array After 1: %v\n", arr)
	utils.AddTo[int](&arr, 4, 3) // Add data to index 3
	fmt.Printf("Array After 2: %v\n", arr)
	utils.AddTo[int](&arr, 25, 1) // Add data to index 1
	fmt.Printf("Array After 3: %v\n", arr)
	arr = utils.Prepend[int](arr, 100) // Prepend data to first index
	fmt.Printf("Array After 4: %v\n", arr)
	arr = append(arr, -5) // Append data to last index available in golang native function
	fmt.Printf("Array After 5: %v\n", arr)
	i := utils.Exist[int](arr, func(v int) bool {
		return v == 9
	})
	fmt.Println("Exist", i)
	j := utils.Exist[int](arr, func(v int) bool {
		return v == 100
	})
	fmt.Println("Exist", j)

	utils.FileMove("/root/project/filename", "/root/project/newname", os.ModePerm)                             // Move file from path to path
	utils.FileMove("/root/project/oldfolder/", "/root/project/newfolder/", os.ModePerm, "filename")            // Move file from to new location with same name
	utils.FileMove("/root/project/oldfolder/", "/root/project/newfolder/", os.ModePerm, "filename", "newname") // Move file from to new location with new name

	arrDuplicate := []int{1, 0, 0, 2, 3, 2, 4, 5, 6, 7, 4, 4, 7, 4, 7, 7, 7, 15}
	fmt.Println(arrDuplicate)
	utils.Distinct[int](&arrDuplicate)
	fmt.Println(arrDuplicate)

	objArr := []map[string]int{
		{"id": 4, "num": 2}, {"id": 3, "num": 9}, {"id": 0, "num": 1}, {"id": 8, "num": 2},
	}
	utils.Where[map[string]int](&objArr, func(obj map[string]int) bool {
		return obj["num"] == 2
	})
	fmt.Printf("Where: %v\n", objArr) // Get all value in slice with condition

	fmt.Printf("First: %v\n", utils.First[map[string]int](objArr)) // Get first value in slice
	fmt.Printf("FirstWhere: %v\n", utils.FirstWhere[map[string]int](objArr, func(obj map[string]int) bool {
		return obj["num"] == 9
	})) // Get first value in slice with given condition
	fmt.Printf("IndexWhere: %v\n", utils.IndexWhere[map[string]int](objArr, func(obj map[string]int) bool {
		return obj["id"] == 0
	})) // Get first index in slice with given condition

	fmt.Printf("IndexesWhere: %v\n", utils.IndexesWhere[map[string]int](objArr, func(obj map[string]int) bool {
		return obj["id"]%2 == 0
	})) // Get indexes in slice with given condition

	fmt.Printf("Last: %v\n", utils.Last(objArr)) // Get first value in slice
	fmt.Printf("LastWhere: %v\n", utils.LastWhere[map[string]int](objArr, func(obj map[string]int) bool {
		return obj["id"] == 4
	})) // Get last value in slice with given condition
	fmt.Printf("LastIndexWhere: %v\n", utils.LastIndexWhere[map[string]int](objArr, func(obj map[string]int) bool {
		return obj["num"] == 2
	})) // Get last index in slice with given condition

	utils.RemoveWhere[map[string]int](&objArr, func(obj map[string]int) bool {
		return obj["id"]%2 == 0
	})
	fmt.Printf("RemoveWhere: %v\n", objArr) // Remove in slice with given condition
}
