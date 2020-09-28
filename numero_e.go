package main

import (
	"fmt"
)

func fact(n int)(result int) {
	if (n > 0) {
		result = n * fact(n-1)
		return result
	}
	return 1
}

func main(){
	var temp int
	var nume float64

	nume = 1
	fmt.Println("Ingrese el número de iteraciones para calcular e: ")
	fmt.Scan(&temp)

	for i := 1; i<=temp; i++ {
		nume = nume + (1/float64(fact(i)))
	}

	fmt.Println(nume)
}