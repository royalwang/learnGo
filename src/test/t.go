package main
import (
	"fmt"
)

type ii interface {
	aa() string
}

type bb struct {
}

func (b *bb) aa() (c string) {
	c = "bb"
	return
}

type tt struct {
	a interface{}
}

func main() {
	b := &tt{a: new(bb)}
	d := b.a.(bb).aa()
	fmt.Println(d)
}
