package main

import "fmt"

func main() {
	name := "lesmana"

	switch name {
	case "didi":
		fmt.Println("helo didi")
	case "lesmana":
		fmt.Println("helo lesmana")
	default:
		fmt.Println("helo world!")
	}

	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("benar")
	case false:
		fmt.Println("salah")
	}
}
