package main

import (
	"fmt"
	"go-practice/hello"
	"go-practice/DataTypes"
	"os"
)

func main() {
	// fmt.Println("Hello, World!")

	// the first arg i.e 0 will be the name of program then the args starting with 1 onwards will be the args that we pass while running the program
	fmt.Println(hello.Say(os.Args[1:]))

	a := 2
	b := 2.4

	fmt.Printf("a: %T %v\n", a, a) // T is for type and v is for value // both are verbs
	fmt.Printf("b: %T %v\n", b, b)

	fmt.Printf("a: %8T %[1]v\n", a) // 8 is for formatting and by providing [1] it will reuse the first parameter that was given
	fmt.Printf("b: %8T %[1]v\n", b)

	//program to give average of sum > input can be taken from terminal or file
	// to take input from file > go run main.go < file.txt or cat file.txt | go run main.go
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values provided")
		os.Exit(-1)
	}

	fmt.Println("The average is ", sum/float64(n))

	fmt.Println(datatypes.Str())

	fmt.Println(datatypes.ArrProg())

	datatypes.SliceProg()

	datatypes.MapsProg()

}
