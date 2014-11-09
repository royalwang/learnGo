package main

import (
	"DesignPattern/ObserverPattern"
)

func main() {
	boss := Observer.NewBoss()
	sec := Observer.NewSecretary()

	ts1 := new(Observer.NBAObserver)
	ts1.Name = "孙星"
	ts2 := new(Observer.StockObserver)
	ts2.Name = "杨2"

	var ecb1 Observer.EventCallback = ts1.GoOnWatchNBADirectSeeding
	var ecb2 Observer.EventCallback = ts1.OpenNBADirectSeeding
	var ecb3 Observer.EventCallback = ts1.CloseNBADirectSeeding

	var ecb4 Observer.EventCallback = ts2.GoOnWatchStockMarket
	var ecb5 Observer.EventCallback = ts2.OpenStockMarket
	var ecb6 Observer.EventCallback = ts2.CloseStockMarket

	boss.AddEventListener(Observer.Boss_in, &ecb3)
	boss.AddEventListener(Observer.Boss_in, &ecb6)

	boss.AddEventListener(Observer.Boss_out, &ecb1)
	boss.AddEventListener(Observer.Boss_out, &ecb4)

	sec.AddEventListener(Observer.Secretary_in_out, &ecb5)
	sec.AddEventListener(Observer.Boss_in, &ecb2)

	boss.Gointo()
	boss.Goout()
	sec.Gointo()
	sec.Goout()
}
