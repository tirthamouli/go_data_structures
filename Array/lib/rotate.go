package library

// RotateFirst - Rotates a slice
func RotateFirst(arr []int, n int) {
	// Step 1: Set up aux
	aux := make([]int, n)
	copy(aux, arr)

	// Step 2: Rotate remaining
	for i := n; i < len(arr); i++ {
		arr[i-n] = arr[i]
	}

	// Step 3: Remaining elements
	for i := 0; i < len(aux); i++ {
		index := len(arr) - n + i
		arr[index] = aux[i]
	}
}
