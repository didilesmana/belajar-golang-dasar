package main

import "fmt"

func sayHelloFilter(name string, filter func(string) string) {
	filterName := filter(name)
	fmt.Println("hello", filterName)
}

func spamFilter(name string) string {
	if name == "gogog" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("didi", spamFilter)
	sayHelloFilter("gogog", spamFilter)

	filter := spamFilter
	sayHelloFilter("gogog", filter)

}
