package main

import (
	"fmt"
)

//这种方式貌似更好
type HandlerAbstractMethod interface { //职责方法抽象
	handleRequest()
}

type handler struct { //职责基本元素的抽象，定义了职责规则
	next       *handler
	handleFunc interface{}
}

func (c *handler) handleRequest() {
	c.handleFunc.(HandlerAbstractMethod).handleRequest()
	if c.next != nil {
		c.next.handleRequest()
	}
}

func (h *handler) setNext(ha *handler) {
	h.next = ha
}

type Concreatehandler1 struct {
}

func (c *Concreatehandler1) handleRequest() {
	fmt.Println("Concreatehandler1")
}

type Concreatehandler2 struct {
}

func (c *Concreatehandler2) handleRequest() {
	fmt.Println("Concreatehandler2")
}

type Concreatehandler3 struct {
}

func (c *Concreatehandler3) handleRequest() {
	fmt.Println("Concreatehandler3")
}

func main() {
	h1 := &handler{handleFunc: new(Concreatehandler1)}
	h2 := &handler{handleFunc: new(Concreatehandler2)}
	h3 := &handler{handleFunc: new(Concreatehandler3)}

	h1.setNext(h2)
	h2.setNext(h3)

	h1.handleRequest()
}
