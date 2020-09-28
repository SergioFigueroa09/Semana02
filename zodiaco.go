package main

import "fmt"

func main2() {
	var mes int
	var dia int

	fmt.Println("Ingrese el mes: ")
	fmt.Scan(&mes)
	fmt.Println("Ingrese el dia: ")
	fmt.Scan(&dia)
	
	if mes == 12 && dia >=22 && dia <= 31 || mes == 1 && dia >= 1 && dia <= 20  {
		fmt.Println("Su signo es Capricornio")
	} else if mes == 1 && dia >=21 && dia <= 31 || mes == 2 && dia >= 1 && dia <= 19  {
		fmt.Println("Su signo es Acuario")
	} else if mes == 2 && dia >=20 && dia <= 29 || mes == 3 && dia >= 1 && dia <= 20  {
		fmt.Println("Su signo es Piscis")
	} else if mes == 3 && dia >=21 && dia <= 31 || mes == 4 && dia >= 1 && dia <= 20  {
		fmt.Println("Su signo es Aries")
	} else if mes == 4 && dia >=21 && dia <= 30 || mes == 5 && dia >= 1 && dia <= 21  {
		fmt.Println("Su signo es Tauro")
	} else if mes == 5 && dia >=22 && dia <= 31 || mes == 6 && dia >= 1 && dia <= 21  {
		fmt.Println("Su signo es Géminis")
	} else if mes == 6 && dia >=22 && dia <= 30 || mes == 7 && dia >= 1 && dia <= 23  {
		fmt.Println("Su signo es Cáncer")
	} else if mes == 7 && dia >=24 && dia <= 31 || mes == 8 && dia >= 1 && dia <= 23  {
		fmt.Println("Su signo es Leo")
	} else if mes == 8 && dia >=24 && dia <= 31 || mes == 9 && dia >= 1 && dia <= 23  {
		fmt.Println("Su signo es Virgo")
	} else if mes == 9 && dia >=24 && dia <= 30 || mes == 10 && dia >= 1 && dia <= 23  {
		fmt.Println("Su signo es Libra")
	} else if mes == 10 && dia >=24 && dia <= 31 || mes == 11 && dia >= 1 && dia <= 22  {
		fmt.Println("Su signo es Escorpio")
	} else if mes == 11 && dia >=23 && dia <= 30 || mes == 12 && dia >= 1 && dia <= 21  {
		fmt.Println("Su signo es Sagitario")
	} else {
		fmt.Println("Ingrese una fecha correcta.")
	}

}