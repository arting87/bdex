package core

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

	"bdex.in/bdex/bdex-ex-backend/bdexerrors"
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/core/matchquote"
)

type Engine struct {
	Handle *matchpool.MatchPool
	Quote  *matchquote.DynaQuoteUtil
	sync.RWMutex
}

//下单
func (e *Engine) PutOrder(o matchpool.CommonOrder) error {
	//订单池订单备份
	//copyOrder := o

	getMatch := false

	quote := e.Quote.GetDynaQuote(&o)
	fmt.Println("++++++++++++++++++++PutOrder++++++++++++++++++++++")
	//判断行情是否满足成交，如果满足，推送成交信息,否则加入队列
	if o.BsFlag == configs.BUYInt {

		prices := make([]float64, 0)
		for sellPrice, _ := range quote.Sell {
			prices = append(prices, sellPrice)
		}

		sort.Float64s(prices) // 升序

		//buyOrderDealAll := false
		//sellOrderDealAll := false

		for _, sellPrice := range prices {

			if o.Price >= sellPrice {
				log.Infow("买单" + strconv.FormatUint(o.Oid, 10) + "匹配到合适价格:买单价格" + strconv.FormatFloat(o.Price, 'f', -1, 64) + "去买价格" + strconv.FormatFloat(sellPrice, 'f', -1, 64))
				log.Infow("寻找此价格的卖单...")
				quote.NewPrice = o.Price
				quote.QuoteTime = time.Now().Format(time.RFC3339)
				//直接成交
				//for key, order := range cOrders {
				cOrders := quote.Sell[sellPrice]
				cOrdersLen := len(cOrders)
				for key := 0; key < cOrdersLen; {

					getMatch = true

					quote.Sell[sellPrice][key].IsHandle = true
					log.Infow("买单已匹配到" + strconv.FormatFloat(sellPrice, 'f', -1, 64) + "价格的订单")
					log.Infow("买单id:" + strconv.FormatUint(o.Oid, 10) + ",MatchSymbol:" + o.MatchSymbol + ",价格:" + strconv.FormatFloat(o.Price, 'f', -1, 64) + ",buyCount:" + strconv.FormatFloat(o.Count, 'f', -1, 64))
					order := quote.Sell[sellPrice][key]
					log.Infow("卖单id:" + strconv.FormatUint(order.Oid, 10) + ",MatchSymbol:" + order.MatchSymbol + ",价格:" + strconv.FormatFloat(order.Price, 'f', -1, 64) + ",sellCount:" + strconv.FormatFloat(order.Count, 'f', -1, 64))
					//5 1
					oaBaseCount := o.Count * o.Price         //0.6
					obBaseCount := order.Count * order.Price //0.1

					if oaBaseCount < obBaseCount {
						//oa被全部成交
						log.Infow("此买单将被全部成交，对方卖单被部分成交")
						obBaseCount = oaBaseCount / order.Price
						//buyOrderDealAll = true
					} else if oaBaseCount > obBaseCount {
						//ob被全部成交
						log.Infow("此买单将被部分成交，对方卖单被全部成交")
						oaBaseCount = obBaseCount
						obBaseCount = order.Count
						//sellOrderDealAll = true
					} else {
						log.Infow("买/卖单都被全部成交")
						oaBaseCount = obBaseCount
						obBaseCount = order.Count
						//sellOrderDealAll = true
						//buyOrderDealAll = true
					}

					//删除行情中此订单
					if key == len(cOrders)-1 {
						cOrders = cOrders[:key]
					} else {
						cOrders = append(cOrders[:key], cOrders[key+1:]...)
					}
					key--
					quote.Sell[sellPrice] = cOrders

					//开启多线程处理
					go func() {

						log.Infow("提交两单数据上链，买单将被成交" + strconv.FormatFloat(oaBaseCount, 'f', -1, 64) + "个，对方卖单将被成交" + strconv.FormatFloat(obBaseCount, 'f', -1, 64) + "个")
						oaerr, oberr := e.SubmitOrders(&o, &order, oaBaseCount, obBaseCount, order.Price)
						//if oaerr != nil || oberr != nil {
						//	//需要在submit中处理错误的情况（回滚、插入到错误订单数据库，并且标记到deleorder数据库）
						//	return
						//	//return handleError(oaerr, oberr)
						//}
						fmt.Println("买单", o.Oid, "剩余Count", o.Count)
						fmt.Println("卖单", order.Oid, "剩余Count", order.Count)

						//判断买订单是否完全完成,如果未完成，则重新放入等待队列
						if oaerr == nil {
							if o.Count > 0 {
								//放入等待队列
								log.Infow("撮合买单"+strconv.FormatUint(o.Oid, 10)+"未被完全撮合，重新放入等待队列", "剩余Count:", o.Count)

								//放入行情
								o.IsHandle = false
								//e.Quote.InsertOrderToDynaQuote(&o)
								//放入等待队列
								e.Handle.GetInnerSktList(o.MatchSymbol).InsertOrder(&o)
							}
						}

						//判断卖订单是否完全完成,如果未完成，则重新放入等待队列
						if oberr == nil {
							if order.Count > 0 {
								//放入等待队列
								log.Infow("行情队列卖单"+strconv.FormatUint(order.Oid, 10)+"未被完全撮合，重新放入等待队列", "剩余Count:", order.Count)

								//放入行情
								order.IsHandle = false
								//放入等待队列
								e.Handle.GetInnerSktList(order.MatchSymbol).InsertOrder(&order)
							}
						}
					}()
					break
				}
			}
			//跳出撮合
			if getMatch {
				break
			}
		}
		//如果没有撮合到则放入行情队列
		if !getMatch {
			log.Infow("此买单"+strconv.FormatUint(o.Oid, 10)+"没有匹配，放入行情队列", "剩余Count:", o.Count)

			//放入行情
			o.IsHandle = false
			e.Quote.InsertOrderToDynaQuote(&o)
		}

	} else {

		prices := make([]float64, 0)
		for buyPrice, _ := range quote.Buy {
			prices = append(prices, buyPrice)
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(prices))) //降序

		//sellOrderDealAll := false
		//buyOrderDealAll := false

		for _, buyPrice := range prices {

			if o.Price <= buyPrice {
				log.Infow("卖单" + strconv.FormatUint(o.Oid, 10) + "匹配到合适价格:卖单价格" + strconv.FormatFloat(o.Price, 'f', -1, 64) + "被价格" + strconv.FormatFloat(buyPrice, 'f', -1, 64) + "买")
				log.Infow("寻找此价格的买单...")
				quote.NewPrice = o.Price
				quote.QuoteTime = time.Now().Format(time.RFC3339)
				//直接成交
				cOrders := quote.Buy[buyPrice]
				cOrdersLen := len(cOrders)
				for key := 0; key < cOrdersLen; {

					getMatch = true

					log.Infow("卖单已匹配到" + strconv.FormatFloat(buyPrice, 'f', -1, 64) + "价格的订单")
					log.Infow("卖单id:" + strconv.FormatUint(o.Oid, 10) + ",MatchSymbol:" + o.MatchSymbol + ",价格:" + strconv.FormatFloat(o.Price, 'f', -1, 64) + ",sellCount:" + strconv.FormatFloat(o.Count, 'f', -1, 64))
					order := cOrders[key]
					log.Infow("买单id:" + strconv.FormatUint(order.Oid, 10) + ",MatchSymbol:" + order.MatchSymbol + ",价格:" + strconv.FormatFloat(order.Price, 'f', -1, 64) + ",buyCount:" + strconv.FormatFloat(order.Count, 'f', -1, 64))
					quote.Buy[buyPrice][key].IsHandle = true

					oaBaseCount := order.Count * order.Price //买单的计价代币的数量 15 EOS
					obBaseCount := o.Count * o.Price         //卖单的计价代币的数量 7.5 EOS

					if oaBaseCount < obBaseCount {
						//oa被全部成交
						log.Infow("此卖单将被部分成交，对方买单被全部成交")
						obBaseCount = oaBaseCount / o.Price
						//buyOrderDealAll = true
					} else if oaBaseCount > obBaseCount {
						//ob被全部成交
						log.Infow("此卖单将被全部成交，对方买单被部分成交")
						oaBaseCount = obBaseCount
						obBaseCount = o.Count
						//sellOrderDealAll = true
					} else {
						log.Infow("卖/买单都被全部成交")
						oaBaseCount = obBaseCount
						obBaseCount = o.Count
						//sellOrderDealAll = true
						//buyOrderDealAll = true
					}

					//删除撮合队列中此订单
					//删除行情内存中此订单并更新
					if key == len(cOrders)-1 {
						cOrders = cOrders[:key]
					} else {
						cOrders = append(cOrders[:key], cOrders[key+1:]...)
					}
					key--
					quote.Buy[buyPrice] = cOrders

					go func() {
						//提交订单数据上链
						log.Infow("提交两单数据上链，卖单将被成交" + strconv.FormatFloat(obBaseCount, 'f', -1, 64) + "个自身代币，对方买单将被成交" + strconv.FormatFloat(oaBaseCount, 'f', -1, 64) + "个自身代币")
						oaerr, oberr := e.SubmitOrders(&order, &o, oaBaseCount, obBaseCount, o.Price)
						fmt.Println("卖单", o.Oid, "剩余Count", o.Count)
						fmt.Println("买单", order.Oid, "剩余Count", order.Count)

						if oberr == nil {
							//判断买订单是否完全完成,如果未完成，则重新放入等待队列
							if o.Count > 0 {
								//放入等待队列
								log.Infow("撮合卖单"+strconv.FormatUint(o.Oid, 10)+"未被完全撮合，重新放入等待队列", "剩余Count:", o.Count)

								//放入行情
								o.IsHandle = false
								//e.Quote.InsertOrderToDynaQuote(&o)
								//放入等待队列
								e.Handle.GetInnerSktList(o.MatchSymbol).InsertOrder(&o)
							}
						}

						if oaerr == nil {
							//判断卖订单是否完全完成,如果未完成，则重新放入等待队列
							if order.Count > 0 {
								//放入等待队列
								log.Infow("行情队列买单"+strconv.FormatUint(order.Oid, 10)+"未被完全撮合，重新放入等待队列", "剩余Count:", order.Count)

								//放入行情
								order.IsHandle = false
								//放入等待队列
								e.Handle.GetInnerSktList(order.MatchSymbol).InsertOrder(&order)
							}
						}

					}()

					cOrdersLen = len(cOrders)
					key++
					break
				}
			}
			if getMatch {
				break
			}
		}
		//如果没有撮合到则放入行情队列
		if !getMatch {
			log.Infow("此卖单"+strconv.FormatUint(o.Oid, 10)+"没有匹配，放入行情队列", "剩余Count:", o.Count)

			//放入行情
			o.IsHandle = false
			e.Quote.InsertOrderToDynaQuote(&o)
		}
	}
	return nil
}

//ordera:买单
//orderb:卖单
//counta:买单撮合的数量，需要使用order中的count*price后传入
//countb:卖单撮合的数量
func (e *Engine) SubmitOrders(ordera, orderb *matchpool.CommonOrder, counta, countb, dealPrice float64) (error, error) {
	orderach := make(chan error, 3)
	orderbch := make(chan error, 3)

	var wait sync.WaitGroup
	oaerr := bdexerrors.ErrNil
	oberr := bdexerrors.ErrNil
	wait.Add(2)
	//处理订单a
	go func() {
		defer wait.Done()

		oaerr = UpstateOneOrder(ordera, orderb.Oid, configs.CannotWithdraw)
		if oaerr != nil {
			log.Errorw("UpstateOneOrder", "error", oaerr)
			orderach <- oaerr
			return
		}
		orderach <- bdexerrors.ErrNil

		select {
		//判断订单b更新状态是否报错
		case upstateBErr := <-orderbch:
			if upstateBErr != bdexerrors.ErrNil {
				//将订单a链上状态修改为true
				oaerr = UpstateOneOrder(ordera, orderb.Oid, configs.AllowWithdraw)
				if oaerr != nil {
					log.Errorw("BackUpstateOneOrder", "error", oaerr)
				}
				fmt.Println("orderberr != bdexerrors.ErrNil")
				return
			}

			oaerr = SubmitOneOrder(ordera, counta, dealPrice, orderb.Oid)
			if oaerr != nil {
				log.Errorw("订单a submit失败", "error", oaerr)
				orderach <- oaerr
				return
			}
			orderach <- bdexerrors.ErrNil

			select {
			//判断订单b提交是否报错
			case submitBErr := <-orderbch:
				if submitBErr != bdexerrors.ErrNil {
					//回滚此a订单
					backorderErr := BackOneOrder(ordera)
					if backorderErr != nil {
					}
					return
				}
				//如果未报错，则进行调用transneed
				transNeedErr := TransNeedOneOrder(ordera, counta, dealPrice)
				if transNeedErr != nil {
					return
				}
			}
		}

	}()
	//处理订单b
	go func() {
		defer wait.Done()

		oberr = UpstateOneOrder(orderb, ordera.Oid, configs.CannotWithdraw)
		if oberr != nil {
			log.Errorw("UpstateOneOrder", "error", oberr)
			orderbch <- oberr
			return
		}
		orderbch <- bdexerrors.ErrNil

		select {
		//判断订单b更新状态是否报错
		case upstateAErr := <-orderach:
			if upstateAErr != bdexerrors.ErrNil {
				//将订单b链上状态修改为true
				oaerr = UpstateOneOrder(orderb, ordera.Oid, configs.AllowWithdraw)
				if oaerr != nil {
					log.Errorw("BackUpstateOneOrder", "error", oaerr)
				}
				fmt.Println("orderaerr != bdexerrors.ErrNil")
				return
			}

			oberr = SubmitOneOrder(orderb, countb, dealPrice, ordera.Oid)
			if oberr != nil {
				log.Errorw("订单b submit失败", "error", oberr)
				orderbch <- oberr
				return
			}
			orderbch <- bdexerrors.ErrNil

			select {
			//判断订单b提交是否报错
			case submitAErr := <-orderach:
				if submitAErr != bdexerrors.ErrNil {
					//回滚此a订单
					backorderErr := BackOneOrder(orderb)
					if backorderErr != nil {
					}
					return
				}
				//如果未报错，则进行调用transneed
				transNeedErr := TransNeedOneOrder(orderb, countb, dealPrice)
				if transNeedErr != nil {
					return
				}
			}
		}
	}()
	wait.Wait()
	//}

	return oaerr, oberr
}
