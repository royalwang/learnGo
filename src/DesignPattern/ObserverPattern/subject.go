package Observer

import ()

type Subject interface { //主题的抽象
	AddEventListener(int, *EventCallback)
	RemoveEventListener(int, *EventCallback)
	Notify()
	setState(string)
	getState() string
}

type Params map[string]interface{} //回调参数

type EventCallback func(Params) //回调函数

type EventChain struct {
	chs       []chan Params
	callbacks []*EventCallback
}

type AbstractSubject struct {
	listeners map[int]*EventChain
	state     string
}

func (this *AbstractSubject) Init_listeners() {
	this.listeners = make(map[int]*EventChain)
}

func (this *AbstractSubject) AddEventListener(event int, callback *EventCallback) {
	eventChain, ok := this.listeners[event]
	if !ok {
		eventChain = &EventChain{chs: []chan Params{}, callbacks: []*EventCallback{}}
		this.listeners[event] = eventChain
	}

	for _, item := range eventChain.callbacks { //已存在对应的回调函数
		if item == callback {
			return
		}
	}

	ch := make(chan Params)
	eventChain.chs = append(eventChain.chs[:], ch)
	eventChain.callbacks = append(eventChain.callbacks[:], callback)

	go func() { //监控每一个channel ，一旦有数据，调用回调函数
		for {
			params := <-ch
			if params == nil {
				break
			}
			(*callback)(params)
		}
	}()
}

func (this *AbstractSubject) RemoveEventListener(event int, callback *EventCallback) {
	eventChain, ok := this.listeners[event]
	if !ok {
		return
	}

	var ch chan Params

	for k, item := range eventChain.callbacks {
		if item == callback {
			ch = eventChain.chs[k]
			ch <- nil
			eventChain.chs = append(eventChain.chs[:k], eventChain.chs[k+1:]...)
			eventChain.callbacks = append(eventChain.callbacks[:k], eventChain.callbacks[k+1:]...)
			break
		}
	}

}

func (this *AbstractSubject) setState(state string) {
	this.state = state
}

func (this *AbstractSubject) getState() (state string) {
	state = this.state
	return
}

func (this *AbstractSubject) DispatchEvent(event int, params Params) { //根据事件，调度回调函数
	eventChain, ok := this.listeners[event]
	if ok {
		for _, chEvent := range eventChain.chs {
			chEvent <- params
		}
	}
}
