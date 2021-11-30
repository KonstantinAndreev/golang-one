package main

import (
	"fmt"
)

func main() {

	fmt.Println("Введите первое число")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Введите второе число")
	var b int
	fmt.Scanln(&b)

	fmt.Println("Введите третье число")
	var c int
	fmt.Scanln(&c)

	fmt.Println("Введите четвертое число")
	var d int
	fmt.Scanln(&d)

	sls := []int{a, b, c, d}
	InsertionSort(sls)
	fmt.Println(sls)
}
func InsertionSort(sls []int) {
	for i := 1; i < len(sls); i++ {
		x := sls[i]
		j := i
		for ; j >= 1 && sls[j-1] > x; j-- {
			sls[j] = sls[j-1]
		}
		sls[j] = x
	}
}
