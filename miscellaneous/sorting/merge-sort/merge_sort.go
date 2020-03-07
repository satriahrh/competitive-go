package merge_sort

func mergeHalve(left, right []int32) (sorted []int32) {
	sorted = make([]int32, len(left)+len(right))
	s, l, r := 0, 0, 0
	for l < len(left) || r < len(right) {
		if l == len(left) {
			sorted[s] = right[r]
			r += 1
		} else if r == len(right) {
			sorted[s] = left[l]
			l += 1
		} else if left[l] < right[r] {
			sorted[s] = left[l]
			l += 1
		} else {
			sorted[s] = right[r]
			r += 1
		}
		s += 1
	}
	return
}

func SortingAlgorithm(arr []int32) (sorted []int32) {
	conquered := make([][]int32, len(arr)+len(arr)%2)
	for i, a := range arr {
		conquered[i] = []int32{a}
	}

	lenConquered := len(conquered)
	for 2 < lenConquered {
		newConquered := make([][]int32, (lenConquered/2)+(lenConquered%2))
		for i := 0; i < lenConquered; i += 2 {
			left := conquered[i]
			right := []int32{}
			if i+1 < lenConquered {
				right = conquered[i+1]
			}
			newConquered[i/2] = mergeHalve(left, right)
		}
		conquered = newConquered
		lenConquered = len(conquered)
	}

	return mergeHalve(conquered[0], conquered[1])
}
