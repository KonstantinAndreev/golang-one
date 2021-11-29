package main

import "fmt"

func main() {
	var i int
	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&i)
	var a = i / 100
	var b = i / 10 % 10
	var c = i % 10
	fmt.Println("Количество сотен:", a)
	fmt.Println("Количество десятков:", b)
	fmt.Println("Количество единиц:", c)
}
