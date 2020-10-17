package main

import "fmt"

func theBiggestNum(args ...int) int {
	bigNum := args[0]

	for _, v := range args {
		if v > bigNum {
			bigNum = v
		}
	}
	return bigNum
}

func main() {
	a := []int{1, 4, 2, 6, 1, 10, 8, 111, 21, 32, 434, 232, 1222}

	fmt.Println("Elementos: ", a)
	fmt.Println("El numero mas grande es: ", theBiggestNum(a...))
}
