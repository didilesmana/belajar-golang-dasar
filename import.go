package main

import (
	"belajar-golang-dasar/helper"
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	result := helper.SayHello("Didi")
	fmt.Println(result)

}
