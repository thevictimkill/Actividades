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
		fmt.Println("Capricornio")
	case (mes == 1 && (dia >= 21 && dia < 32)) || (mes == 2 && (dia <= 19 && dia > 0)):
		fmt.Println("Acuario")
	case (mes == 2 && (dia >= 20 && dia < 30)) || (mes == 3 && (dia <= 20 && dia > 0)):
		fmt.Println("Pisis")
	case (mes == 3 && (dia >= 21 && dia < 32)) || (mes == 4 && (dia <= 20 && dia > 0)):
		fmt.Println("Aries")
	case (mes == 4 && (dia >= 21 && dia < 31)) || (mes == 5 && (dia <= 21 && dia > 0)):
		fmt.Println("Tauro")
	case (mes == 5 && (dia >= 22 && dia < 32)) || (mes == 6 && (dia <= 21 && dia > 0)):
		fmt.Println("Geminis")
	case (mes == 6 && (dia >= 22 && dia < 31)) || (mes == 7 && (dia <= 23 && dia > 0)):
		fmt.Println("Cancer")
	case (mes == 7 && (dia >= 24 && dia < 32)) || (mes == 8 && (dia <= 23 && dia > 0)):
		fmt.Println("Leo")
	case (mes == 8 && (dia >= 24 && dia < 32)) || (mes == 9 && (dia <= 23 && dia > 0)):
		fmt.Println("Virgo")
	case (mes == 9 && (dia >= 24 && dia < 31)) || (mes == 10 && (dia <= 23 && dia > 0)):
		fmt.Println("Libra")
	case (mes == 10 && (dia >= 24 && dia < 32)) || (mes == 11 && (dia <= 22 && dia > 0)):
		fmt.Println("Escorpio")
	case (mes == 11 && (dia >= 23 && dia < 31)) || (mes == 12 && (dia <= 21 && dia > 0)):
		fmt.Println("Sagitario")
	default:
		fmt.Println("fecha no valida")
	}

}
