package datatypes

import (
	"fmt"
	"strings"
)

func Str() string {
	s := "élite"
	a := []byte(s)
	b := []rune(s)
	fmt.Printf("%8T, %[1]v, %d\n", s, len(s))
	fmt.Printf("%8T, %[1]v, %d\n", a, len(a))
	fmt.Printf("%8T, %[1]v, %d\n", b, len(b))


	newStr := "The quick brown fox"
	fmt.Printf("newStr: %v\n", newStr)
	x := len(newStr)
	y := newStr[:3]
	z := newStr[4:9]
	w := newStr[:3] + " slow" + newStr[9:]

	fmt.Printf("x: %v, y: %v, z: %v, w: %v\n", x,y,z,w)
	// newStr[5] = 'a' // syntax error as strings are immutable
	newStr += "es" // can be done as it will point to a new reference

	fmt.Printf("newStr after adding es: %v\n", newStr) // newStr points to new memory as we can't change a string

	newStr = strings.ToUpper(newStr) // same it will point to new memory

	return newStr
}