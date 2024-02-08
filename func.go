package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func sayHelloTo(firstName string, lastName string, no int) {
	fmt.Println("hello", firstName, lastName, no)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullname() (string, string) {
	return "Didi", "lesmana"
}

func getCompleteName() (firstName, lastName string) {
	firstName = "didi"
	lastName = "lesmana"

	return firstName, lastName
}

func main() {
	sayHello()
	sayHello()
	sayHello()

	sayHelloTo("Didi", "Lesmana", 10)
	sayHelloTo("Didi", "Lesmana", 5)
	sayHelloTo("Lesmana", "didi", 1)

	result := getHello("didi")
	fmt.Println(result)
	fmt.Println(getHello("lesmana"))

	fmt.Println(getFullname())

	firstName, _ := getFullname()
	fmt.Println(firstName)

	a, b := getCompleteName()
	fmt.Println(a, b)

	fmt.Println(getCompleteName())

}
