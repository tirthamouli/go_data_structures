package main

import (
	"fmt"

	heap "github.com/Boom0027/go_data_structures/Heap/lib"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap.Heapify(arr)
	fmt.Println(arr)

	arr, elem := heap.Delete(arr)
	fmt.Println(arr, elem)
	arr, elem = heap.Delete(arr)
	fmt.Println(arr, elem)
}
