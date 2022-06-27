package main

import (
	"fmt"

	bs "github.com/h-mavrodiev/data-structures-and-algorithms/pkg/bisearch"
)

func main() {
	a := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Input array : %v \n", a)
	for _, v := range a {
		res, err := bs.BiSearch(a, v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v is at position %v.\n", v, res)
	}
}
