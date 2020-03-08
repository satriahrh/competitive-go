package k_smallest_element

func countingSort(arr []int) (sortedMap map[int]int) {
	sortedMap = make(map[int]int)
	for _, a := range arr {
		sortedMap[a] += 1
	}
	return
}

func AlgorithmCountingSort(k int, arr []int) int {
	sortedMap := countingSort(arr)
	i := 0
	for ; 0 < k; {
		i += 1
		k -= sortedMap[i]
	}
	return i
}

func heapParent(i int) int {
	return (i - 1) / 2
}

func heapChildLeft(i int) int {
	return 2*i + 1
}

func heapChildRight(i int) int {
	return 2*i + 2
}

func heapShiftDown(arr []int, start, end int) {
	root := start
	for ;heapChildLeft(root) <= end; {
		child := heapChildLeft(root)
		swap := root
		if arr[swap] < arr[child] {
			swap = child
		}
		if child + 1 <= end && arr[swap] < arr[child+1] {
			swap = child + 1
		}
		if swap == root {
			return
		} else {
			temp := arr[root]
			arr[root] = arr[swap]
			arr[swap] = temp
			root = swap
		}
	}
}

func heapShiftUp(arr []int, start, end int) {
	root := start
	for ;heapChildLeft(root) <= end; {
		child := heapChildLeft(root)
		swap := root
		if arr[swap] > arr[child] {
			swap = child
		}
		if child + 1 <= end && arr[swap] > arr[child+1] {
			swap = child + 1
		}
		if swap == root {
			return
		} else {
			temp := arr[root]
			arr[root] = arr[swap]
			arr[swap] = temp
			root = swap
		}
	}
}

func heapSort(arr []int) (sorted []int) {
	// heapify
	start := heapParent(len(arr) - 1)
	for ; 0 <= start; start  -= 1 {
		heapShiftUp(arr, start, len(arr)-1)
	}

	end := len(arr) - 1
	for ; end > 0; {
		temp := arr[end]
		arr[end] = arr[0]
		arr[0] = temp

		end = end - 1
		heapShiftUp(arr, 0, end)
	}
	return arr
}

func AlgorithmHeapSort(k int, arr []int) int {
	sortedArr := heapSort(arr)
	return sortedArr[len(arr) - k]
}
