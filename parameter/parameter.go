package main

import "fmt"

func plus(x int, y int) int {
	return x + y
}

func minus(x, y int) int {
	return x - y
}

func main() {
	fmt.Println("3 더하기 5 =", plus(3, 5))
	fmt.Println("10 빼기 7 =", minus(10, 7))
}
