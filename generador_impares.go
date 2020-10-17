package main

import "fmt"

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var impar = i
		i += 2

		return impar

	}
}

func main() {
	nextPar := generadorImpares()
	for i := 1; i <= 50; i++ {
		fmt.Println(nextPar())
	}

}
