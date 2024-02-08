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

	name = "didilesmana"
	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Panjang")
	case lenght > 5:
		fmt.Println("Sedang")
	default:
		fmt.Println("Pendek")
	}
}
