package main

import "fmt"

func getGoodBye(name string) string {
	return "Good " + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("didi"))
	fmt.Println(misal("lesmana"))
}
