package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "goodbye " + name
}

func SayHello(name string) string {
	return "Helo " + name
}

func helper() {
	result := sayGoodBye("lesmana")
	fmt.Println(result)
}
