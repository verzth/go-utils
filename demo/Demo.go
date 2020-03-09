package main

import (
	"fmt"
	"github.com/verzth/go-utils/utils"
)

func main() {
	fmt.Printf("RandString: %v\n", utils.RandString(20))
	fmt.Printf("RandStringWithCharset: %v\n", utils.RandStringWithCharset(20,"abc123efg456"))
	fmt.Printf("RandStringLower: %v\n", utils.RandStringLower(20))
	fmt.Printf("RandomHex: %v\n",utils.RandomHex(16))
	fmt.Printf("RandomHexUpper: %v\n",utils.RandomHexUpper(16))
	ranString := utils.RandStringUpper(10)
	fmt.Printf("RandStringUpper: %v\n",ranString) // 'ea10a9c919'

	fmt.Printf("Substring: %v\n", utils.Substring(ranString,2)) // '10a9c919'
	fmt.Printf("Substring: %v\n", utils.Substring(ranString,1,8)) // 'a10a9c9'
	fmt.Printf("Substring: %v\n", utils.Substring(ranString,6,2)) // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString,10,15)) // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString,-1,5)) // ''
	fmt.Printf("Substring: %v\n", utils.Substring(ranString,5,10)) // '9c919'

	arr := []int{ 1,2,3,4,5 }
	fmt.Printf("Array Before: %v\n", arr)
	utils.RemoveAt(&arr, 3) // Remove slice at index
	fmt.Printf("Array After 1: %v\n", arr)
	utils.AddTo(&arr, 4, 3)  // Add data to index 3
	fmt.Printf("Array After 2: %v\n", arr)
	utils.AddTo(&arr, 25, 1) // Add data to index 1
	fmt.Printf("Array After 3: %v\n", arr)
	utils.AddTo(&arr, 100) // Add data without specify index, will be added to last index
	fmt.Printf("Array After 4: %v\n", arr)
	utils.AddTo(&arr, -5) // Add data with minus index, will be added to first index which is treated as 0
	fmt.Printf("Array After 5: %v\n", arr)

	utils.FileMove("/root/project/filename","/root/project/newname") // Move file from path to path
	utils.FileMove("/root/project/oldfolder/","/root/project/newfolder/", "filename") // Move file from to new location with same name
	utils.FileMove("/root/project/oldfolder/","/root/project/newfolder/", "filename", "newname") // Move file from to new location with new name
}
