package main

import (
	"fmt"
)

func main() {
	p := &[3]int{1, 2, 3}
	fmt.Printf("%d\n", (*p)[0]) // c style
	//In go ranguage, runtime judges pointer value
	fmt.Printf("%d\n", p[0]) // go style
}
