package main

import (
	"fmt"
	"strings"
)

func main(){
	//fmt.Println(strings.Contains("distribuidos","bui"))
	//fmt.Println(strings.Count("distribuidos","i"))
	//fmt.Println(strings.HasPrefix("distribuidos","di"))
	//fmt.Println(strings.HasSuffix("distribuidos","os"))
	//fmt.Println(strings.Index("distribuidos","os"))
	//fmt.Println(strings.Join([]string{"Sistemas", "Distribuidos","Cucei"},"-"))
	//fmt.Println(strings.Repeat("distribuidos", 2))
	//fmt.Println(strings.Replace("aaaaabbb","a","b",2))
	//fmt.Println(strings.Split("Mi mama me mi mima", " "))
	//fmt.Println(strings.ToLower("Mi mama me mi mima"))
	fmt.Println(strings.ToUpper("Mi mama me mi mima"))

	b := []byte("test")
	fmt.Println(b)

	s := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(s)
}