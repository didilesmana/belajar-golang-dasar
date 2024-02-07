package main

import "fmt"

func main() {

	type NoKTP string

	var ktpDidi NoKTP = "11111111111"

	var contoh string = "22222222222"
	var contohLtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpDidi)
	fmt.Println(contohLtp)
}
