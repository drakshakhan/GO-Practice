package datatypes

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//to run the code comment all the code in main file except this function call
// type go run main.go < test.txt
func Prog() {
	fmt.Println("New program with map hands on start here")
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")

	type kv struct {
		key   string
		value int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].value > ss[j].value
	})

	for _, s := range ss[:3]{
		fmt.Println(s.key, "appears", s.value, "times")
	}
}
