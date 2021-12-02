package main

import (
	fmt "fmt"
)

func main() {
	ar1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		var a int
		fmt.Println("Введите число:")
		fmt.Scanln(&a)
		ar1[i] = a
	}
	for i := 1; i < len(ar1); i++ {
		x := ar1[i]
		j := i
		for ; j >= 1 && ar1[j-1] > x; j-- {
			ar1[j] = ar1[j-1]
		}
		ar1[j] = x
	}
	fmt.Println("Результат сортировки:\n", ar1)
}
