package main

import (
	"fmt"

	library "github.com/Boom0027/go_data_structures/Array/lib"
)

func main() {
	// Rotate
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	library.RotateFirst(arr, 3)
	fmt.Println(arr)

	// Merge sort
	arr = []int{3, 2, 76, 2, 45, 7, 78, 47, 23}
	res := make([]int, len(arr))
	library.MergeSort(arr, res, 0, len(arr)-1)
	fmt.Println(res)
}
