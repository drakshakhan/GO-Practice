package st

import "fmt"

func Str() string {
	s := "élite"
	a := []byte(s)
	b := []rune(s)
	fmt.Printf("%8T, %[1]v, %d\n", s, len(s))
	fmt.Printf("%8T, %[1]v, %d\n", a, len(a))
	fmt.Printf("%8T, %[1]v, %d\n", b, len(b))
	return s
}