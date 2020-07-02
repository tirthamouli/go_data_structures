package library

// Merge - Merge left and right into res
func merge(req []int, res []int, start int, mid int, end int) {
	// Step 1: Initialize
	i, j, k := start, mid+1, start

	// Step 2: Calculate for level 1
	for i <= mid && j <= end {
		if req[i] <= req[j] {
			res[k] = req[i]
			i++
		} else {
			res[k] = req[j]
			j++
		}
		k++
	}

	// Step 2: Calculate for remaining
	for i <= mid {
		res[k] = req[i]
		i++
		k++
	}
	for j <= end {
		res[k] = req[j]
		j++
		k++
	}

	// Step 3: Copy to req
	for i = start; i <= end; i++ {
		req[i] = res[i]
	}
}

// MergeSort - Merge sort
func MergeSort(req []int, res []int, start int, end int) {
	// Step 0: Base case
	if start == end {
		return
	}

	// Step 1: Get the middle point
	mid := (start + end) / 2

	// Step 2: Merge sort from start to mid
	MergeSort(req, res, start, mid)

	// Step 3: Merge sort from the mid+1 to the end
	MergeSort(req, res, mid+1, end)

	// Step 4: Merge into the res
	merge(req, res, start, mid, end)
}
