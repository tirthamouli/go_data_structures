package library

// Reverse - Reverses an array
func Reverse(arr []int, s int, l int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
