package main

import (
	"fmt"
)

type operation interface {
	getResult() float64
	SetNumA(float64)
	SetNumB(float64)
}

type BaseOperation struct {
	numA float64
	numB float64
}

func (operation *BaseOperation) SetNumA(numA float64) {
	operation.numA = numA
}

func (operation *BaseOperation) SetNumB(numB float64) {
	operation.numB = numB
}

type OperationAdd struct {
	BaseOperation
}

func (this *OperationAdd) getResult() float64 {
	return this.numA + this.numB
}

type OperationSub struct {
	BaseOperation
}

func (this *OperationSub) getResult() float64 {
	return this.numA - this.numB
}

type operationFactory struct {
}

func (of *operationFactory) createOP(s string) (op operation) {
	switch s {
	case "+":
		op = new(OperationAdd)
	case "-":
		op = new(OperationSub)
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
	oper.SetNumA(3.0)
	oper.SetNumB(1.0)
	fmt.Println(oper.getResult())
}
