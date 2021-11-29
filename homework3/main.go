package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println("Введите первое число")
	var a float64

	fmt.Scanln(&a)

	fmt.Println("Введите второе число")
	var b float64

	fmt.Scanln(&b)

	fmt.Println("Введите операцию,(+, -, *, /, ^, !): ")

	var operation string
	fmt.Scanln(&operation)

	switch operation {

	case "+":
		fmt.Printf("%.2f\n", a+b)
	case "-":
		fmt.Printf("%.2f\n", a-b)
	case "*":
		fmt.Printf("%.2f\n", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка деления на ноль")
		} else {
			fmt.Printf("%.2f\n", a/b)
		}
	case "^":
		fmt.Printf("%.2f\n", math.Pow(a, b))
	case "!":
		var fact = new(big.Int)
		if b != 0 {
			fmt.Println("Ошибка,для вычисления факториала" +
				" введите только первое число!")
		} else {

			fact.MulRange(1, int64(a))
			fmt.Println(fact.String())
		}
	}

}
