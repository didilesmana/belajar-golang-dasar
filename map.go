package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "didi",
		"address": "depok",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku golang"
	book["author"] = "didi"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
