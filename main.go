package main

import "fmt"

const (
	c1 = iota
	_
	c3
	c4
)

func joinSlice(sl1, sl2 []int) []int {
	return append(sl1, sl2...)
}

func main() {
	sl1 := make([]int, 10, 15)

	sl2 := []int{10, 12, 18, 5}

	fmt.Println("Len sl1 = ", len(sl1), "Cap sl1 = ", cap(sl1))
	fmt.Println("Len sl2 = ", len(sl2), "Cap sl2 = ", cap(sl2))

	result := joinSlice(sl1, sl2)

	fmt.Println("Result for working func join: ", result)
}
