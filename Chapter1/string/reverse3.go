package main

import "fmt"

func main() {
	var str string
	fmt.Println("enter a string  to  reverse :")
	fmt.Scan(&str)

	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	fmt.Println("Reversed string:", reversed)
}
