package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type test struct {
	Reply interface{}
	Done  chan *test
}

func main() {
	fmt.Println("Hello, world.  你好，世界！")
	//	ch := make(chan int, 1)
	//	for {
	//		select {
	//		case ch <- 0:
	//		case ch <- 1:
	//		}
	//		i := <-ch
	//		fmt.Println("Value received:", i)
	//	}
	var i int
	mt := new(test)
	mt.Reply = &i
	done := make(chan *test, 10)
	mt.Done = done
	fmt.Println(mt.Reply)
	fmt.Println(&i)

	value := reflect.ValueOf(mt.Reply)
	fmt.Println(value.Type().Kind() == reflect.Ptr)

	v := unsafeAddr(reflect.ValueOf(mt.Reply))

	fmt.Println(v)

	mt.Done <- mt
	//	nt := new(test)
	//	j := 1
	//	nt.Reply = &j
	//	nt.Done = done
	//	nt.Done <- nt

	tt := <-mt.Done
	fmt.Println(reflect.TypeOf(tt))
	fmt.Println(i)
	//	fmt.Println(nt == tt)

}

func unsafeAddr(v reflect.Value) unsafe.Pointer {
	if v.CanAddr() {
		fmt.Println("here")
		return unsafe.Pointer(v.UnsafeAddr())
	}
	x := reflect.New(v.Type()).Elem()
	x.Set(v)
	return unsafe.Pointer(x.UnsafeAddr())
}
