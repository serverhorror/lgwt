package main

import "fmt"

func HelloI18nWho(who, lang string) (string, error) {
	if who == "" {
		who = "World"
	}

	switch lang {
	case "en":
		return "Hello " + who + "!", nil
	case "es":
		return "Hola " + who + "!", nil
	case "fr":
		return "Bonjour " + who + "!", nil
	default:
		return "Hello " + who + "!", fmt.Errorf("language %v is unknown!", lang)
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
