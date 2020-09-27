package main

import "fmt"

func main() {
	var lado float64

	fmt.Println("Ingresa el lado: ")
	fmt.Scanf("%f", &lado)
	area := lado * lado
	fmt.Println("El area es: ", area)
}