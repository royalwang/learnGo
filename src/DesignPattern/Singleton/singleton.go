package main

/*
单利模式在GO中以ONCE关键字来做 ,
同样，对于只需执行一次的函数也用ONCE
*/
import (
	"fmt"
	"sync"
)

var once sync.Once
var Mysingle *singleton

type singleton struct {
	s string
}

func NewSigleton() *singleton {
	once.Do(func() {
		Mysingle = new(singleton)
		Mysingle.s = "first"
	})
	return Mysingle
}

func main() {
	//	s1 := NewSigleton()
	//	fmt.Println(s1)
	//	s2 := NewSigleton()
	//	fmt.Println(s2)
	//
	//	s1.s = "abcd"
	//	fmt.Println(s1)
	//	fmt.Println(s2)

	ch := make(chan int)
	end := make(chan int,2)

	f1 := func(i1 chan int,end chan int) {
		s1 := NewSigleton()
		fmt.Println(s1)
		<-i1
		s1.s = "abcd"
		fmt.Println(s1)
		i1 <- 1
		end<-1
	}

	f2 := func(i2 chan int,end chan int) {
		s2 := NewSigleton()
		fmt.Println(s2)
		i2 <- 2
		<-i2
		fmt.Println(s2)
		end<-2
	}

	go f1(ch,end)
	go f2(ch,end)
	<-end
	<-end

	close(ch)
	close(end)
}
