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

// // RotateSecond - Another method of rotating, needs no aux
// func RotateSecond(arr []int, n int) {
// 	for i := 0; i < n; i++ {
// 		prev := arr[i]
// 		fmt.Println("i", i, prev)
// 		for j := len(arr) - n + i; j >= i; j -= n {
// 			fmt.Println("j", j)
// 			prev, arr[j] = arr[j], prev
// 		}
// 	}
// }
