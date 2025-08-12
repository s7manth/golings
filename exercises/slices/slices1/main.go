// slices1
// Make me compile!

package main

import "fmt"

func main() {
	a := make([]int, 4, 10) // play with length and capacity
	b := make(map[string]string)
	_ = b
	fmt.Println("length of 'a':", len(a))
	fmt.Println("capacity of 'a':", cap(a))
}
