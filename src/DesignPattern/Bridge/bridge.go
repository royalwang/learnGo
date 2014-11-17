package main

//将抽象部分与实现部分分离，使它们都可以独立的变化

//实现具体实现类的解耦
import (
	"fmt"
)

type AbstractCar interface {
	run()
}

type Ford struct {
	AbstractCar
}

func (f *Ford) run() {
	fmt.Println("Ford is running")
}


type AbstractRoad struct {
	car AbstractCar
}

type SuperWay struct {
	AbstractRoad
}

func (s *SuperWay) runOn() {
	fmt.Println("On the SuperWay")
	s.car.run()
}

type StreatWay struct {
	AbstractRoad
}

func (s *StreatWay) runOn() {
	fmt.Println("On the StreatWay")
	s.car.run()
}

func NewStreatWay(car AbstractCar) *StreatWay {
	return &StreatWay{AbstractRoad: AbstractRoad{car: car}}
}

func main() {
	car := new(Ford)
	way := NewStreatWay(car)
	way.runOn()
}
