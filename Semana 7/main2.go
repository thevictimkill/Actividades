package main

import (
	"fmt"
	"time"
)

var flag = 0
var eliminar = 0
//goroutine
func pinger(c chan string) {
	for {
		if eliminar == 0{
			c <- "ping"
		}
		if eliminar == 1{
			return 
		}
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		if flag == 1 {
			msg := <- c
			fmt.Println(msg)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//gorutine
func main()  {
	var c chan string = make(chan string)

	flag = 1
	go pinger(c)
	go ponger(c)
	go printer(c)
	


	var input string
	fmt.Scanln(&input)

	eliminar = 1
	

	var input2 string
	fmt.Scanln(&input2)

}