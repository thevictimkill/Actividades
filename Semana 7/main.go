package main

import (
	"fmt"
	"time"
)

//goroutine
func f(n int) {
	for i := 0; i <10; i++{
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

//gorutine
func main()  {
	go f(0)
	go f(1)

	var input string
	fmt.Scanln(&input)
}