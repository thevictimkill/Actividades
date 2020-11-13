package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)
	fmt.Scan(&mes)

	//fmt.Println(dia)
	//fmt.Println(mes)

	switch {
	case (mes == 12 && (dia >= 22 && dia < 32)) || (mes == 1 && (dia <= 20 && dia > 0)):
		fmt.Println("capricornio")
	case (mes == 1 && (dia >= 21 && dia < 32)) || (mes == 2 && (dia <= 19 && dia > 0)):
		fmt.Println("acuario")
	case (mes == 2 && (dia >= 20 && dia < 30)) || (mes == 3 && (dia <= 20 && dia > 0)):
		fmt.Println("pisis")
	case (mes == 3 && (dia >= 21 && dia < 32)) || (mes == 4 && (dia <= 20 && dia > 0)):
		fmt.Println("aries")
	case (mes == 4 && (dia >= 21 && dia < 31)) || (mes == 5 && (dia <= 21 && dia > 0)):
		fmt.Println("tauro")
	case (mes == 5 && (dia >= 22 && dia < 32)) || (mes == 6 && (dia <= 21 && dia > 0)):
		fmt.Println("geminis")
	case (mes == 6 && (dia >= 22 && dia < 31)) || (mes == 7 && (dia <= 23 && dia > 0)):
		fmt.Println("cancer")
	case (mes == 7 && (dia >= 24 && dia < 32)) || (mes == 8 && (dia <= 23 && dia > 0)):
		fmt.Println("leo")
	case (mes == 8 && (dia >= 24 && dia < 32)) || (mes == 9 && (dia <= 23 && dia > 0)):
		fmt.Println("virgo")
	case (mes == 9 && (dia >= 24 && dia < 31)) || (mes == 10 && (dia <= 23 && dia > 0)):
		fmt.Println("libra")
	case (mes == 10 && (dia >= 24 && dia < 32)) || (mes == 11 && (dia <= 22 && dia > 0)):
		fmt.Println("escorpio")
	case (mes == 11 && (dia >= 23 && dia < 31)) || (mes == 12 && (dia <= 21 && dia > 0)):
		fmt.Println("sagitario")
	default:
		fmt.Println("fecha no valida")
	}

}
