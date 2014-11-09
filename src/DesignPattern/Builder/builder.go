package main

import (
	"fmt"
	"reflect"
	//"unsafe"
)

type Person struct { //目的对象
	head, body, leftarm, rightarm, leftleg, rightleg string
}

type PersonBuilder interface { //抽象处的构建过程
	buildHead()
	buildBody()
	buildArmRight()
	buildArmLeft()
	buildLegRight()
	buildLegLeft()
	display_person() *Person
}

type PersonThinBuilder struct {
	thin_person *Person //最终目标
}

func (this *PersonThinBuilder) buildBody() {
	this.thin_person.body = "draw thin body..."
}

func (this *PersonThinBuilder) buildHead() {
	this.thin_person.head = "draw thin head..."
}

func (this *PersonThinBuilder) buildArmRight() {
	this.thin_person.rightarm = "draw thin arm right..."
}

func (this *PersonThinBuilder) buildArmLeft() {
	this.thin_person.leftarm = "draw thin arm left..."
}

func (this *PersonThinBuilder) buildLegRight() {
	this.thin_person.rightleg = "draw thin leg right..."
}

func (this *PersonThinBuilder) buildLegLeft() {
	this.thin_person.leftleg = "draw thin leg left..."
}

func (p *PersonThinBuilder) display_person() *Person {
	return p.thin_person
}

func NewPersonThinBuilder() *PersonThinBuilder {
	//return new(PersonThinBuilder)
	return &PersonThinBuilder{new(Person)}
}

type PersonFatBuilder struct {
	fat_person *Person
}

func (this *PersonFatBuilder) buildBody() {
	this.fat_person.body = "draw fat body..."
}

func (this *PersonFatBuilder) buildHead() {
	this.fat_person.head = "draw fat head..."
}

func (this *PersonFatBuilder) buildArmRight() {
	this.fat_person.rightarm = "draw fat arm right..."
}

func (this *PersonFatBuilder) buildArmLeft() {
	this.fat_person.leftarm = "draw fat arm left..."
}

func (this *PersonFatBuilder) buildLegRight() {
	this.fat_person.rightleg = "draw fat leg right..."
}

func (this *PersonFatBuilder) buildLegLeft() {
	this.fat_person.leftleg = "draw fat leg left..."
}

func (p *PersonFatBuilder) display_person() *Person {
	return p.fat_person
}

func NewPersonFatBuilder() *PersonFatBuilder {
	return &PersonFatBuilder{new(Person)}
}

type PersonDirector struct {
	pb PersonBuilder
}

func (p *PersonDirector) createPerson() {
	p.pb.buildHead()
	p.pb.buildBody()
	p.pb.buildArmLeft()
	p.pb.buildArmRight()
	p.pb.buildLegLeft()
	p.pb.buildLegRight()
}

func (p *PersonDirector) display_personBuilding() *Person {
	return p.pb.display_person()
}

func main() {

	ptb := NewPersonThinBuilder()
	pdThin := &PersonDirector{ptb}
	pdThin.createPerson()
	person := pdThin.display_personBuilding()
	fmt.Println(person.head)
	fmt.Println()
	value := reflect.ValueOf(*person)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(value.Field(i))
	}
	//x := unsafe.Pointer(building)
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x))))
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(building.body))))
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(building.leftarm))))
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(building.rightarm))))
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(building.leftleg))))
	//fmt.Println(*(*string)(unsafe.Pointer(uintptr(x) + unsafe.Offsetof(building.righleg))))
	//fmt.Println(building)
}
