package main

import (
	"fmt"
	"math"
	"math/big"
	_ "math/big"
)

func main() {
	var fact = new(big.Int)

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
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка деления на ноль")
		} else {
			fmt.Println(a / b)
		}
	case "^":
		fmt.Println(math.Pow(a, b))
	case "!":
		if b != 0 {
			fmt.Println("Ошибка,для вычисления факториала" +
				" введите только первое число!")
		} else {

			fact.MulRange(1, int64(a))
			fmt.Println(fact.String())
		}
	}

}
