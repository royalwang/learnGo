package Observer

import ()

//业务相关主题动作抽象
type officeOP interface {
	Gointo()
	Goout()
}

//具体的主题类
type Secretary struct {
	AbstractSubject
}

func (this *Secretary) Notify(event int) {
	params := make(map[string]interface{})
	params["id"] = 1001
	this.DispatchEvent(event, params)
}

func (this *Secretary) Gointo() {
	this.Notify(Secretary_in_out)
}

func (this *Secretary) Goout() {
	this.Notify(Secretary_in_out)
}

type Boss struct {
	AbstractSubject
}

func (this *Boss) Notify(event int) {
	params := make(map[string]interface{})
	params["id"] = 1001
	this.DispatchEvent(event, params)
}

func (this *Boss) Gointo() {
	this.Notify(Boss_in)
}

func (this *Boss) Goout() {
	this.Notify(Boss_out)
}

func NewBoss() (b *Boss) {
	b = new(Boss)
	b.Init_listeners()
	return
}

func NewSecretary() (s *Secretary) {
	s = new(Secretary)
	s.Init_listeners()
	return
}
