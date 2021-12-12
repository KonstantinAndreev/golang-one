package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type linkedListNod struct {
	val  int
	next *linkedListNod
}

func main() {
	var numberList int
	fmt.Println("Please type integer values:")
	_, err := fmt.Scanln(&numberList)
	if err != nil {
		defer log.Fatal("Error! Not integer values!\n", err)
		panic(os.Exit)
	}

	var fNod *linkedListNod
	fNod = &linkedListNod{val: numberList}
	pNod := fNod

	for {
		if _, err := fmt.Scanln(&numberList); errors.Is(err, io.EOF) {
			fmt.Println("Waiting!")
			fmt.Println(" ")
			break
		}
		nextNod := &linkedListNod{val: numberList}
		pNod.next = nextNod
		pNod = nextNod
	}

	var rListNod *linkedListNod
	pRevNod := fNod
	curNod := fNod.next
	fNod.next = nil
	for {
		if curNod == nil {
			rListNod = pRevNod
			break
		}
		nextNod := curNod.next
		curNod.next = pRevNod
		pRevNod = curNod
		curNod = nextNod
	}
	fmt.Println("Done, sorted values:")
	for rListNod != nil {
		fmt.Println(rListNod.val)
		rListNod = rListNod.next
	}
}
