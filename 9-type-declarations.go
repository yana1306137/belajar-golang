package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	var noKTPYana NoKTP = "32147274927"
	var marriedYana Married = true
	fmt.Println(noKTPYana)
	fmt.Println(marriedYana)
}
