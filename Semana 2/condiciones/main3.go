package main

import "fmt"

func main() {
	var op int

	fmt.Println("Opciones: ")
	fmt.Scan(&op)

	switch op {
	case 1:
		fmt.Println("Opcion 1")
	case 2:
		fmt.Println("opcion 2")
	case 3:
		fmt.Println("Opcion 3")
	default:
		fmt.Println("Opcion no exise")
		
	}
	
}
