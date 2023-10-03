package main

import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("2 - 1 = ", 2-1)
	fmt.Println("2 * 2 = ", 2*2)
	fmt.Println("2 / 2 = ", 2/2)
	fmt.Println("10 % 3 = ", 10%3)
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(false || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(!true)
	fmt.Println(!false)
}
