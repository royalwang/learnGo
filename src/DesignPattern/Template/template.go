package main

/*
  模板方法也可用反射实现 ， 但是并不好，尽量不用反射
*/
import (
	"fmt"
)

type Abclass struct {
	abmethod interface{} //抽象方法集合接口
}

type AbMethods interface { //定义出抽象方法集合
	run()
	jump()
}

func (a *Abclass) Dosports() { //抽象模板中的具体方法
	fmt.Println("doing sports")
	a.abmethod.(AbMethods).run()
	a.abmethod.(AbMethods).jump()
}

type Conclass1 struct {
	//Abclass
}

func (c *Conclass1) run() { //具体实现方法
	fmt.Println("run1....")
}

func (c *Conclass1) jump() { //具体实现方法
	fmt.Println("jump1...")
}

type Conclass2 struct {
	//Abclass
}

func (c *Conclass2) run() { //具体实现方法
	fmt.Println("run2....")
}

func (c *Conclass2) jump() { //具体实现方法
	fmt.Println("jump2...")
}

func ConcFactory(i int) (a *Abclass) {
	a = new(Abclass)
	switch i {
	case 1:
		a.abmethod = new(Conclass1)
	case 2:
		a.abmethod = new(Conclass2)
	}
	return
}

func main() {
	c1 := ConcFactory(1)
	c2 := ConcFactory(2)

	c1.Dosports()
	c2.Dosports()
}
