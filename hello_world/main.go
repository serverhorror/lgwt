package main

import "fmt"

func HelloWho(who string) string {
	if who == "" {
		who = "World"
	}
	return "Hello " + who + "!"
}

func HelloWorld() string {
	return "Hello World!"
}

func main() {
	fmt.Println(HelloWorld())
}
