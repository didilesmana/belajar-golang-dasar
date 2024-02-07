package main

import "fmt"

func main() {
	name := "lesmana"

	if name == "didi" {
		fmt.Println("helo world")
	} else if name == "joko" {
		fmt.Println("Hello dodo")
	} else {
		fmt.Println("hello eko")
	}

	if lenght := len(name); lenght > 5 {
		fmt.Println("nama panjang")
	} else {
		fmt.Println("nama pendek")
	}

}
