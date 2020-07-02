package heap

// CheckHeap - Checks if heap, if not convert to heap
func CheckHeap(arr []int, n int) {
	// Step 1: When there is no right child - there should always be a left child
	if (n*2)+1 >= len(arr) {
		// Step 1.1 Check if left child is greater than the current node
		if arr[n*2] > arr[n] {
			arr[n], arr[n*2] = arr[n*2], arr[n]
		}
		return
	}

	// Step 2: Has both left and right child
	cur := arr[n]
	left := arr[n*2]
	right := arr[(n*2)+1]

	// Step 3 Check if left is the greatest
	if left >= right && left > cur {
		// Step 3.1: Swap
		arr[n], arr[n*2] = arr[n*2], arr[n]

		// Step 3.2: Check if the child with which the swap was made with is a heap if required
		if n*2 <= (len(arr)/2)-1 {
			CheckHeap(arr, n*2)
		}
		return
	}

	// Step 4: Check if right is the greatest
	if right > left && right > cur {
		// Step 4.1: Swap
		arr[n], arr[(n*2)+1] = arr[(n*2)+1], arr[n]

		// Step 3.2: Check if the child with which the swap was made with is a heap if required
		if (n*2)+1 <= (len(arr)/2)-1 {
			CheckHeap(arr, (n*2)+1)
		}
		return
	}
}
