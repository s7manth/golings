// slices2
// Make me compile!

package main

import "fmt"

func main() {
	names := [4]string{"John", "Maria", "Carl", "Peter"}
	lastTwoNames := names[len(names)-2:len(names)]
	fmt.Println(lastTwoNames)
}
