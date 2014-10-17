package main

import (
	"errors"
	"fmt"
	"reflect"
)

func Read(b []byte) (n int, err error) {
	err = errors.New("Should be non-negative numbers!")
	return
}

func main() {
	var j int = 5

	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	fmt.Print(reflect.TypeOf(a))

	//	a()
	//
	//	j *= 2
	//
	//	a()
}
