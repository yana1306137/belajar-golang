package main

import "fmt"

func main() {
	var name1 = "Yana"
	var name2 = "Yana"
	var name3 = "Unay"
	var result = name1 == name2
	fmt.Println(result)
	fmt.Println(name1 > name3)

	var no1 = 1
	var no2 = 2
	fmt.Println(no1 >= no2)
	fmt.Println(no1 != no2)
}
