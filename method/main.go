package main

import (
	"fmt"
)

type Point struct{ X, Y int }

// *Point型のメソッドRender
func (p *Point) Render() {
	fmt.Printf("X:%d, Y:%d\n", p.X, p.Y)
}

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

type User struct {
	Id   int
	Name string
}

func NewUser(id int, name string) *User {
	u := new(User)
	u.Id = id
	u.Name = name
	return u
}

func main() {
	p := Point{X: 10, Y: 20}
	p.Render() // X:10, Y:20
	// p.Renderの別の呼び出し方
	((*Point).Render)(&Point{X: 100, Y: 200}) // X:100, Y:200
	// さらに別のp.Renderの呼び出し方
	p2 := Point{X: 5, Y: 10}
	f := p2.Render
	f() // X:5, Y:10

	// int型のaliasを貼ってメソッドを追加する
	fmt.Println(MyInt(4).Plus(2))

	// コンストラクタのパターン
	fmt.Println(NewUser(1, "Taro")) // &{1 Taro}
}
