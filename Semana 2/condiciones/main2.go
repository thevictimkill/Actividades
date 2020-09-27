package main

import "fmt"

func main() {
	var temp int
	fmt.Println("Temp: ")
	fmt.Scan(&temp)
	if temp < 0 {
		fmt.Println("Esta helado")
	} else if temp >= 0 && temp < 12 {
		fmt.Println("Esta haciendo frio")
	} else if temp >= 12 && temp < 18 {
		fmt.Println("Esta agusto")
	} else {
		fmt.Println("Esta caluroso")
	}
}
