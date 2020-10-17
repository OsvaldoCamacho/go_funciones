package main

import "fmt"

func intercambio(a, b *int) { // se recibe un puntero

	var temp int = *a
	*a = *b
	*b = temp

}

func main() {
	x := 10
	y := 20

	fmt.Println(" a: ", x, "   b:", y)
	intercambio(&x, &y)
	fmt.Println("intercambio...")
	fmt.Println(" a: ", x, "   b:", y)

}
