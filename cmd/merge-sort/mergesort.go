package main

import (
	"fmt"

	ms "github.com/h-mavrodiev/data-structures-and-algorithms/pkg/mergesort"
)

func main() {
	a := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}

	sorted := ms.MergeSort(a)
	fmt.Println(sorted)
}
