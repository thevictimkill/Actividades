package main

import (
	"fmt"
	"time"
)

var flag = 0

func main() {
	var op int
	auxid := 1
	flagc := 1

	for flagc == 1{
		fmt.Println("Opciones: ")
		fmt.Scan(&op)

		switch op {
		case 1:
			pro := Proceso{1,auxid}
			pro.Mostrar()
			proce := Procesosall{
				Procesos: []Proceso{
					&pro,
				  },
			  }
			auxid++
		case 2:
			if flag == 1{
				flag = 0
			}
			if flag == 0{
				flag = 1
			}
		case 3:
			fmt.Println("Eliminar proceso")
		case 4:
			fmt.Println("Salir")
		case 4:
			flagc = 0
		default:
			fmt.Println("Opcion no exise")
			
		}
	}
	
}