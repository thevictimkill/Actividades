package main

import "fmt"

func main() {
	var temp int

	fmt.Println("Temp: ")
	fmt.Scan(&temp)

	switch {
	case temp < 0:
		fmt.Println("Esta helado")
	case temp >=0 && temp < 12:
		fmt.Println("Esta haciendo frio")
	case temp >=12 && temp < 18:
		fmt.Println("Esta agusto")
	default:
		fmt.Println("Esta caluroso")
	}
	
}
