package main

import "fmt"

func fibonacci(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}

func main() {
	var x uint
	fmt.Println("Ingresa un numero:")
	fmt.Scan(&x)
	fmt.Println("Resultado: ", fibonacci(x))
}
