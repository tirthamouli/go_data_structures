package main

import (
	"fmt"

	library "github.com/Boom0027/data_structure/Array/lib"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	library.RotateFirst(arr, 3)
	fmt.Println(arr)
}
