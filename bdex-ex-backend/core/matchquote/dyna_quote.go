package matchquote

import (
	"sync"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/log"
)

/**
撮合引擎的撮合队列
*/
type DynaQuoteUtil struct {
	Quotes map[string]*DynaQuote //行情表 k:matchSymbol key:Code v:行情表
	sync.RWMutex
}

//撮合队列实例模型
type DynaQuote struct {
	MatchSymbol string                              `xorm:"'match_symbol' varchar(8) notnull unique pk"` //匹配symbol  AAA,1    DDD,2
	Code        string                              `xorm:"'Code' varchar(255) "`                        //合约代码
	Name        string                              `xorm:"'name' varchar(255)"`                         //合约名称
	Buy         map[float64][]matchpool.CommonOrder `xorm:"-"`                                           //买档 k:price  v:count
	Sell        map[float64][]matchpool.CommonOrder `xorm:"-"`                                           //卖档 k:price  v:count
	NewPrice    float64                             `xorm:"'new_price' double"`                          //最新价
	MarketCode  int                                 `xorm:"'market_code' Int"`                           //市场代码
	QuoteTime   string                              `xorm:"'quote_time' varchar(255)"`                   //行情时间
}

func NewDynaQuoteUtil() *DynaQuoteUtil {
	return &DynaQuoteUtil{
		Quotes: make(map[string]*DynaQuote),
	}
}

func NewNilDyncQuote(co *matchpool.CommonOrder) *DynaQuote {
	dynaQuote := DynaQuote{
		Code:       co.Code,
		Name:       co.Name,
		Buy:        make(map[float64][]matchpool.CommonOrder, 0),
		Sell:       make(map[float64][]matchpool.CommonOrder, 0),
		NewPrice:   0,
		MarketCode: co.MarketCode,
		QuoteTime:  co.OrderTime,
	}
	return &dynaQuote
}

//func (dqu *DynaQuoteUtil) FlushQuoteMap(quote *DynaQuote) {
//	//刷新内存应该是同步操作
//	dqu.lock.Lock()
//	defer dqu.lock.Unlock()
//	//通过当前完成的订单刷新当前交易行情
//	dqu.Quotes[quote.MarketCode][quote.Code] = quote
//}

func (dqu *DynaQuoteUtil) GetDynaQuote(o *matchpool.CommonOrder) *DynaQuote {
	if dqu.Quotes[o.MatchSymbol] == nil {
		log.Infow("创建一个新的行情", "MatchSymbol:", o.MatchSymbol)
		dqu.Quotes[o.MatchSymbol] = NewNilDyncQuote(o)
	}
	return dqu.Quotes[o.MatchSymbol]
}

func (qu *DynaQuoteUtil) InsertOrderToDynaQuote(cOrder *matchpool.CommonOrder) {
	begin := time.Now()
	if cOrder.BsFlag == configs.BUYInt {
		qu.GetDynaQuote(cOrder).Buy[cOrder.Price] = append(qu.GetDynaQuote(cOrder).Buy[cOrder.Price], *cOrder)
	} else {
		qu.GetDynaQuote(cOrder).Sell[cOrder.Price] = append(qu.GetDynaQuote(cOrder).Sell[cOrder.Price], *cOrder)
	}
	log.Infow("当前行情"+cOrder.Code, "插入耗时", time.Now().Sub(begin), "价格", cOrder.Price)
	log.Infof("当前买队列行情:%s %s:%d", cOrder.Code, "大小", len(qu.GetDynaQuote(cOrder).Buy[cOrder.Price]))
	log.Infof("当前卖队列行情:%s %s:%d", cOrder.Code, "大小", len(qu.GetDynaQuote(cOrder).Sell[cOrder.Price]))
}
