package datatypes

import "fmt"

func MapsProg() {
	fmt.Println("Maps program starts from here..................")
	var m map[string]int
	n := make(map[string]int)

	fmt.Println("Printing map m: ", m) // nil no storage 
	fmt.Println("printing map n: ", n) // non nil, but empty

	p := m["hi"] // return 0 as m is nil, we can read from a nil map but write operation can't be done it will panic
	q := n["hi"] // return 0

	// m["hello"] = 1 // panic as we cant perform write operation in nil map > panic: assignment to entry in nil map

	// fmt.Println("printing m again to see if by doing insertion operation the code will panic or not")
	m = n // now m will not be nil, as it will point to same hash table as n

	fmt.Printf("printing p %v and q: %v", p , q)

	m["hey"]++

	fmt.Println("checking if inserting in m will reflect or not as it is not nil now: ", m)

	r := n["hey"]
	
	fmt.Printf("printing r to see if n will reflect the change done in m as both the maps i.e, m and n points to same storage. r: %v, n: %v", r, n)

	var o = map[string]int{
		"why": 1,
		"who": 1,
		"where": 2,
	}

	fmt.Println("printing o: ", o)

	var s map[string]int // nil

	fmt.Println("printing q", q)

	//operations in map

	// a := o == s // syntax error as we cant compare maps to one other > map can only be compared to nil
	b := s == nil // true as maps can be compared only to nil as special case
	c  := len(o) // 3
	// d := cap(o) // type mismatch, there is nothing such as capacity in maps > invalid argument: o (variable of type map[string]int) for built-in cap
	
	fmt.Printf("printing o: %v , b: %v, c: %d: ", o, b, c)

	// maps have a special two result lookup function
	// the variable will tell if the key was there or not
	
	k := map[string]int{} // non nil but empty
	fmt.Printf("k is: %v", k)

	l := k["what"] // return 0
	t, ok := k["what"] // 0, false
	fmt.Println("ok: ", ok)
	k["hoho"]++

	u, ok := k["hoho"] // 1, true

	fmt.Printf("printing k: %v, l: %v, t: %v, ok: %v, u: %v",k, l, t, ok, u)

}