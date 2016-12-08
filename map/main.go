package main

import (
	"fmt"
)

func main() {
	map_1 := make(map[int]string)
	map_1[1] = "aaa"
	map_1[2] = "bbb"
	map_1[3] = "ccc"

	fmt.Println(map_1) // map[1:aaa 2:bbb 3:ccc]

	map_2 := map[int]string{
		1: "aaa",
		2: "bbb",
		3: "ccc",
	}

	fmt.Println(map_2) // map[1:aaa 2:bbb 3:ccc]

	map_3 := map[int][]int{
		1: {1, 2, 3},
		2: {4, 5, 6},
		3: {7, 8, 9},
	}
	fmt.Println(map_3) // map[1:[1 2 3] 2:[4 5 6] 3:[7 8 9]]

	// detect exist key
	s, ok := map_1[1]
	fmt.Println(s, ok)
	s, ok = map_1[9]
	fmt.Println(s, ok)
}
