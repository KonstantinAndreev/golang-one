package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type lsNod struct {
	val int
	new *lsNod
}

func main() {
	var numberList int
	fmt.Println("Please type ints:")
	if _, err := fmt.Scanln(&numberList); errors.Is(err, io.EOF) {
		os.Exit(1)
	}

	var fNod *lsNod
	fNod = &lsNod{val: numberList}
	pNod := fNod

	for {
		if _, err := fmt.Scanln(&numberList); errors.Is(err, io.EOF) {
			fmt.Println("Done!")
			break
		}
		nextNod := &lsNod{val: numberList}
		pNod.new = nextNod
		pNod = nextNod
	}

	var rListNod *lsNod
	pRevNod := fNod
	cNod := fNod.new
	fNod.new = nil
	for {
		if cNod == nil {
			rListNod = pRevNod
			break
		}
		nNod := cNod.new
		cNod.new = pRevNod
		pRevNod = cNod
		cNod = nNod
	}

	for rListNod != nil {
		fmt.Println(rListNod.val)
		rListNod = rListNod.new
	}
}
