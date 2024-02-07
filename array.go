package main

import (
	"fmt"
)

func main() {
	var names [2]string
	names[0] = "didi"
	names[1] = "lesmana"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{70, 80, 90, 95}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(len(values))
	values[3] = 100
	fmt.Println(values)

}
