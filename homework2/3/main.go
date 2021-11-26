package main

import "fmt"

func main() {

	var i int
	fmt.Println("Введите  трехзначное число")
	fmt.Scanln(&i)
	var a = (i - (i%100))/1
	var b = (i - (i%10))/10
	var c = (i - (i%1))/100
	fmt.Println("Количество единиц:", a)
	fmt.Println("Количество десятков:", b)
	fmt.Println("Количество сотен:", c)
}
