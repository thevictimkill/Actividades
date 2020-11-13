package main

import "fmt"

func main()  {
	var size int
  fmt.Scan(size)
  var slice [...]int{}
  var aux int
  var i = 0

  for  i < size {
    fmt.Scan(aux)
    slice = append(slice, aux)
    fmt.Println(slice)
    i++
  }
}