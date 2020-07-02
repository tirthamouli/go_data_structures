package heap

// Heapify - Convert an array to heap
func Heapify(arr []int) {
	for i := (len(arr) / 2) - 1; i >= 1; i-- {
		CheckHeap(arr, i)
	}
}
