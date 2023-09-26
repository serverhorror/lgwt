package main

import "fmt"

func HelloWho(who string) string {
	return "Hello " + who + "!"
}

func HelloWorld() string {
	return "Hello World!"
}

func main() {
	fmt.Println(HelloWorld())
}
