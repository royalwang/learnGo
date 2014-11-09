package Observer

import (
	"fmt"
)

//观察者
type AbstractObserver struct {
	Name string
}

type NBAObserver struct {
	AbstractObserver
}

func (this *NBAObserver) GoOnWatchNBADirectSeeding(params Params) {
	fmt.Println(this.Name, "继续看NBA直播！")
}

func (this *NBAObserver) OpenNBADirectSeeding(params Params) {
	fmt.Println(this.Name, "打开NBA直播，停止工作！")
}

func (this *NBAObserver) CloseNBADirectSeeding(params Params) {
	fmt.Println(this.Name, "关闭NBA直播，继续工作！")
}

type StockObserver struct {
	AbstractObserver
}

func (this *StockObserver) GoOnWatchStockMarket(params Params) {
	fmt.Println(this.Name, "继续看股票行情！")
}

func (this *StockObserver) OpenStockMarket(params Params) {
	fmt.Println(this.Name, "打开股票行情，停止工作！")
}

func (this *StockObserver) CloseStockMarket(params Params) {
	fmt.Println(this.Name, "关闭股票行情，继续工作！")
}
