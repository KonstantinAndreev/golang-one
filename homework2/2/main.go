package main

import "fmt"
import "math"

func main() {
	var a float64
	fmt.Println("Калькулятор расчета диаметра и длины окружности")
	fmt.Println("Введите площадь окружности")
	fmt.Scanln(&a)
	var b = math.Sqrt(a * 4 * math.Pi)
	fmt.Printf("Длина окружности равна: %f\n", b)
	fmt.Printf("Диаметр равен: %f\n", b/math.Pi)
}
