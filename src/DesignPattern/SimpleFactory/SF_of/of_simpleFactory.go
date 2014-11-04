package main
/*
传统的策略模式写法
*/
import (
	"fmt"
)

type operation interface {
	getResult(float64, float64) (c float64)
}

type OP func(float64, float64) float64

func (op OP) getResult(a float64, b float64) (c float64) {
	c = op(a, b)
	return
}

func Add(a float64, b float64) (c float64) {
	c = a + b
	return
}

func Sub(a float64, b float64) (c float64) {
	c = a - b
	return
}

type operationFactory struct {
}

func (of *operationFactory) createOP(s string) (op operation) {
	switch s {
	case "+":
		op = OP(Add)
	case "-":
		op = OP(Sub)
	default:
		panic("运算符号错误！")
	}
	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var fac operationFactory
	oper := fac.createOP("-")
	fmt.Println(oper.getResult(3.0, 2.0))
}
