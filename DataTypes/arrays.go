package datatypes

import "fmt"

func ArrProg() [4]int { // uppercase func names will be exported and small case will be imported
	var a [3]int
	b := [3]int{2,8,7}
	var c = [...]int{1,3,1}

	var d [3]int
	d = b // assignable as arrays are passed by value, thus element can be copied

	var m = [...]int{1,2,3,4}
	// m = c // type mismatch

	fmt.Printf("a: %v\n , b: %v\n , c : %v\n , d: %v\n , m: %v\n", a, b, c, d, m)

	return m

}