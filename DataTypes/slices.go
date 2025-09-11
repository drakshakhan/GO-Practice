package datatypes

import "fmt"

func SliceProg() []int{
	// slices are underlying arrays that have variable length
	// slices have length and capacity
	// slices are passed by reference; no copying but updating is ok
    // [2:8] > 2 will be included but 8 won't
	fmt.Println("Slices Program starts here")
	var a []int
	var b = []int{5,4,7,8,2}

	fmt.Printf("a: %v, b: %v\n", a, b)

	a = append(a, 1)
	b = append(b, 6)

	fmt.Printf("a: %v, b: %v\n", a, b)

	c := make([]int, 7)
	d := b // b will overwrite d so now d and b will point to same memory
	// that means if we do d[0] == b[0] it will return true

	fmt.Printf("c: %v, d: %v\n", c, d)

	c = a


	fmt.Printf("a: %v\n, b: %v\n, c:%v\n, d: %v\n", a, b, c, d)
	return nil
}