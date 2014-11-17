package main

//由于不能像其他面向对象语言定义类时那样自持有，所以只能将结构和方法分开抽象
import (
	"fmt"
)

type HandlerAbstractMethod interface { //职责方法抽象
	handleRequest()
}

type handler struct { //职责结构抽象
	next HandlerAbstractMethod
}

func (h *handler) setNext(ha HandlerAbstractMethod) {
	h.next = ha
}

type Concreatehandler1 struct {
	handler
}

func (c *Concreatehandler1) handleRequest() {
	fmt.Println("Concreatehandler1")
	c.next.handleRequest()
}

type Concreatehandler2 struct {
	handler
}

func (c *Concreatehandler2) handleRequest() {
	fmt.Println("Concreatehandler2")
	c.next.handleRequest()
}

type Concreatehandler3 struct {
	handler
}

func (c *Concreatehandler3) handleRequest() {
	fmt.Println("Concreatehandler3")
}

func main() {
	c1 := new(Concreatehandler1)
	c2 := new(Concreatehandler2)
	c3 := new(Concreatehandler3)

	c1.setNext(c2)
	c2.setNext(c3)
	
	c1.handleRequest()
}
