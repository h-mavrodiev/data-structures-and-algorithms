package mergesort

func MergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	a := MergeSort(unsorted[len(unsorted)/2:])
	b := MergeSort(unsorted[:len(unsorted)/2])

	return merge(a, b)

}

func merge(a []int, b []int) []int {
	merged := []int{}
	i := 0
	j := 0

	for len(a) > i && len(b) > j {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else if a[i] > b[j] {
			merged = append(merged, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}

	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return merged

}
