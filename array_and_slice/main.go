package main

import (
	"fmt"
)

func array_pow(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func slice_pow(a []int) {
	for i, v := range a {
		a[i] = v * v
	}
	return
}

func main() {
	array := [3]int{1, 2, 3}
	slice_1 := make([]int, 3)
	slice_2 := []int{1, 2, 3}

	fmt.Printf("array type is %T\n", array)
	fmt.Printf("slice_1 type is  %T\n", slice_1)
	fmt.Printf("slice_2 type is  %T\n", slice_2)

	fmt.Println(array)   // [1 2 3]
	fmt.Println(slice_1) // [0 0 0]
	slice_1[0] = 1
	slice_1[1] = 2
	slice_1[2] = 3
	fmt.Println(slice_1) // [1 2 3]

	fmt.Printf("array len = %d, cap = %d\n", len(array), cap(array))
	fmt.Printf("slice_1 len = %d, cap = %d\n", len(slice_1), cap(slice_1))

	//append
	slice_1 = append(slice_1, 4)
	fmt.Println(slice_1) // [1 2 3 4]
	fmt.Printf("slice_1 len = %d, cap = %d\n", len(slice_1), cap(slice_1))

	slice_1 = append(slice_1, 5)
	fmt.Println(slice_1)
	fmt.Printf("slice_1 len = %d, cap = %d\n", len(slice_1), cap(slice_1))

	slice_1 = append(slice_1, 6)
	fmt.Println(slice_1)
	fmt.Printf("slice_1 len = %d, cap = %d\n", len(slice_1), cap(slice_1))

	slice_1 = append(slice_1, 7)
	fmt.Println(slice_1)
	fmt.Printf("slice_1 len = %d, cap = %d\n", len(slice_1), cap(slice_1))

	//array = append(array, 4) // first argument to append must be slice

	array_pow(array) // call by value
	fmt.Printf("array_pow = %v\n", array)

	slice_3 := []int{1, 2, 3}
	slice_pow(slice_3) // call by reference
	fmt.Printf("slice_pow = %v\n", slice_3)
}
