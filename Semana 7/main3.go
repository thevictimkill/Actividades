package main

import (
	"fmt"
	"time"
	"strconv"
)



//gorutine
func main()  {
	c1 := make(chan string)
	c2 := make(chan string)

	go func (){
		aux :=0
		for {
			c1 <- strconv.Itoa(aux)
			time.Sleep(time.Second * 2)
			aux ++
		}
	}()

	go func () {
		aux :=0
		for {
			c1 <- strconv.Itoa(aux)
			time.Sleep(time.Second * 2)
			aux ++
		}
	}()

	go func() {
		for {
			select {
			case msg := <- c1:
				fmt.Println(msg)
			case msg := <- c2:
				fmt.Println(msg)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

	

	var input2 string
	fmt.Scanln(&input2)
}