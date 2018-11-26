// Boiling prints the boiling point of water.
package main

import (
	"fmt"
)

const boilingF = 212.0

func main() {
	f := boilingF
	c := (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
	// Output:
	// boiling point = 212F or 100C
}
