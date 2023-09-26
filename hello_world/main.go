package main

import "fmt"

func HelloI18nWho(who, lang string) (string, error) {
	if who == "" {
		who = "World"
	}
	if lang == "en" {
		return "Hello " + who + "!", nil
	} else if lang == "es" {
		return "Hola " + who + "!", nil
	} else if lang == "fr" {
		return "Bonjour " + who + "!", nil
	} else {
		return "Hello " + who + "!", fmt.Errorf("Language %v is unknown!", lang)
	}
}

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
