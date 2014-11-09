package main

import (
	"fmt"
)

type ii interface {
	aa() string
}

type bb struct {
	tt
}

func (b *bb) aa() (c string) {
	c = "bb"
	return
}

type tt struct {
	a interface{}
}

func main() {
	b := new(tt)
	var e ii = new(bb)
	b.a = e
	//b := &tt{a: new(bb)}
	//	b := &tt{a: &bb{}}
	
	d := b.a.(ii)//d := b.a.(*bb)
	fmt.Println(d.aa())
}
