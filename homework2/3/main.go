package main

import "fmt"

func main() {

	var i int
	fmt.Println("Введите  трехзначное число")
	fmt.Scanln(&i)
	var a = i / 100
	var b = i / 10
	var c = i / 1
	fmt.Println("Количество единиц:", a)
	fmt.Println("Количество десятков:", b)
	fmt.Println("Количество сотен:", c)
}
