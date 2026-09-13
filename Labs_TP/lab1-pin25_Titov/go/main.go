package main

import "fmt"

// greet формирует приветственное сообщение для пользователя.
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Print("Enter your name: ")

	var name string
	fmt.Scanln(&name)

	fmt.Println(greet(name))
}
