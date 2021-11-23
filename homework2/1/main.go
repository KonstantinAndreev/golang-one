package main

import "fmt"

func main() {
	var a, b float32
	fmt.Printf("Калькулятор расчета площади прямоугольника \n")
	fmt.Println("Введите первое число")
	fmt.Scanln(&a)

	fmt.Println("Введите второе число")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника равна: %f\n", a*b)
}
