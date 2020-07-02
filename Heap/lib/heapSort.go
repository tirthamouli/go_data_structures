package heap

// Sort - Sort an array
func Sort(arr []int) []int {
	// Step 1: Create heap
	Heapify(arr)

	// Step 2: Create res arr
	resArr := make([]int, len(arr)-1)
	for i := 0; i < len(resArr); i++ {
		arr, resArr[i] = Delete(arr)
	}

	// Step 3: Return the result
	return resArr
}
