package main

import (
	"fmt"
)

// type Point struct { X, Y int }
type Point struct {
	X int
	Y int
}

func value_swap(p Point) {
	x, y := p.X, p.Y
	p.X = y
	p.Y = x
	return
}

func reference_swap(p *Point) {
	x, y := p.X, p.Y
	p.X = y
	p.Y = x
	return
}

func main() {
	var pt_1 Point
	pt_1.X = 10
	pt_1.Y = 20
	fmt.Println(pt_1)

	pt_2 := Point{1, 2}
	fmt.Println(pt_2)

	pt_3 := Point{X: 100, Y: 200}
	fmt.Println(pt_3)

	// Omission of type name
	type Hoge struct {
		int
		float64
		string
	}
	h_1 := Hoge{1, 1.5, "hoge"}
	fmt.Println(h_1)

	// struct inside struct
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed // Omission fieald name
	}

	a := Animal{}
	a.Name = "dog"
	a.Amount = 100000 // automatic detect field
	a.Feed.Name = "springdog"

	fmt.Println(a)

	//swap
	pt_4 := Point{X: 1, Y: 2}
	fmt.Println(pt_4) // {1 2}

	// call by value
	value_swap(pt_4)
	fmt.Println(pt_4) // {1 2}

	// referene by value
	reference_swap(&pt_4)
	fmt.Println(pt_4) // {2 1}

	// another method of generate struct's pointer
	pt_5 := new(Point)
	pt_5.X = 10
	pt_5.Y = 20
	fmt.Println(pt_5) // {2 1}

	reference_swap(pt_5)
	fmt.Println(pt_5) // &{20 10}
}
