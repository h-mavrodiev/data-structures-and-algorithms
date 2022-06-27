package bisearch

import (
	"errors"
	"sort"
)

func BiSearch(a []int, target int) (int, error) {

	// Binary search demands a sorted array
	if isSorted := sort.IntsAreSorted(a); !isSorted {
		err := errors.New("the provided arrays is not sorted. Binary search works only on sorted array")
		return -999, err

	}

	// h is the high value of for our "half elimination"
	// -1 as len counts from 1 and we need h is an index
	h := len(a) - 1

	// l is the low value index for our "half elimination"
	l := 0

	// for loop will run as long as l <= h
	for l <= h {
		mid := (h + l) / 2
		if target == a[mid] {
			return mid, nil
		} else if target > a[mid] {
			l = mid + 1
		} else if target < a[mid] {
			h = mid - 1
		}
	}

	return -999, nil

}
