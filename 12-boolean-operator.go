package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80
	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusUjian && lulusAbsensi)
}
