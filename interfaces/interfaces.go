package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	fmt.Println("i = &T{\"Hello\"}에 대해")
	i = &T{"Hello"}
	describe(i)
	i.M()

	fmt.Println("i = F(math.Pi)에 대해")
	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("인터페이스의 (값, 타입) : (%v, %T)\n", i, i)
}
