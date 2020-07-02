package heap

// Delete - Deletes an element from the heap
func Delete(arr []int) ([]int, int) {
	// Step 1: Get the element
	elem := arr[1]

	// Step 2: Transfer last child to root
	arr[1] = arr[len(arr)-1]

	// Step 3: Check if root is heap
	resArr := arr[:len(arr)-1]
	CheckHeap(resArr, 1)

	// Step 4: Return the new array and the elem
	return resArr, elem
}
