package main

import (
	"fmt"
	"golangOne/golang-one/homework10/sort"
)

func main() {
	var length int
	fmt.Println("Введите длинну слайса:")
	fmt.Scanln(&length)

	arr := make([]int, 0, length)
	for i := 0; i < length; i++ {
		var number int
		fmt.Println("Введите число:")
		fmt.Scanln(&number)
		arr = append(arr, number)
	}

	arr = sort.InsertionSort(arr)

	fmt.Println("Результат сортировки:\n", arr)
}
