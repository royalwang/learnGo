package main

import ("fmt")

type People interface {
	wash()
	sweep()
	say()
}

//抽象产品类,基类
type Leifeng struct {
	People
}

func (l *Leifeng) sweep() {
    fmt.Println("無償扫地")
}
 
func (l *Leifeng) wash() {
    fmt.Println("無償洗衣")
}

type Payman struct {
    People
}

func (l *Payman) sweep() {
    fmt.Println("有償扫地")
}
 
func (l *Payman) wash() {
    fmt.Println("有償洗衣")
}

//具体产品类
type Volunteer struct {
   Leifeng
}

func (l *Volunteer) say() {
    fmt.Println("我是Volunteer")
}

type GoodStudent struct {
   Leifeng
}

func (l *GoodStudent) say() {
    fmt.Println("我是GoodStudent")
}

type Cleaner struct{
   Payman
}

func (l *Cleaner) say() {
    fmt.Println("我是Cleaner")
}
//工場抽象接口
type abstractfactory interface{
  sendSomeoneDoJob(string) People
}

type WCFactory struct{
    abstractfactory
}

func (w *WCFactory)sendSomeoneDoJob(k string) (p People){
   switch k{
     case "v":
        p = new(Volunteer)
     case "s":
        p = new(GoodStudent)
     default :
       panic("无相关类型")
   }
   return
}


type YCFactory struct{
    abstractfactory
}

func (y *YCFactory)sendSomeoneDoJob(k string) (p People){
   switch k{
     case "c":
        p = new(Cleaner)
     default :
       panic("无相关类型")
   }
   return
}

type God struct{
   s1 abstractfactory
   s2 abstractfactory
}

func main(){
   g:=&God{s1:new(WCFactory),s2:new(YCFactory)}
   //工场创建产品
   v:=g.s1.sendSomeoneDoJob("v")
   s:=g.s1.sendSomeoneDoJob("s")
   c:=g.s2.sendSomeoneDoJob("c")
   
   //打印产品信息
   v.say()
   s.say()
   c.say()
}

