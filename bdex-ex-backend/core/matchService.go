package core

import (
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"sync"
	"time"
)

/**
Description:循环把等待队列中每个队列拿出撮合

Parameter:

Return:
*/
//循环撮合等待队列与行情队列数据
func Match() {
	go func() {
		for {
			lens := len(EXCore.Engine.Handle.MatchQueueMap)
			if lens == 0 {
				time.Sleep(1 / 100 * time.Second)
				continue
			}
			var wait sync.WaitGroup

			EXCore.Engine.Handle.RLock()
			for _, list := range EXCore.Engine.Handle.MatchQueueMap {
				wait.Add(1)
				EXCore.Engine.Handle.RUnlock()
				go func() {
					EXCore.Engine.Handle.Lock()
					EXCore.Engine.DealMatch(list)
					EXCore.Engine.Handle.Unlock()
					wait.Done()
				}()
				EXCore.Engine.Handle.RLock()
			}
			EXCore.Engine.Handle.RUnlock()
			wait.Wait()
		}
	}()
}

/**
Description:循环撮合等待队列与行情队列数据

Parameter:
	sml:一个MarketCode对应的等待队列

Return:
*/
func (e *Engine) DealMatch(sml *matchpool.StkMatchList) {
	for {
		if sml.List.SellList.Len() == 0 && sml.List.BuyList.Len() == 0 {
			//线程休眠
			break
		}

		for {
			order := sml.GetOrder(configs.BUYInt)
			if order == nil {
				break
			}
			//订单时间大于行情时间不能进行本次撮合,(老行情不能撮合新订单)
			//if order.OrderTime >= q.QuoteTime {
			//	continue
			//}

			e.Handle.MatchQueueMap[order.MatchSymbol].RemoveOrderByObj(*order)

			err := e.PutOrder(*order)
			if err != nil {
				log.Errorw("PutOrder", "error", err)
				break
			}

		}
		for {
			order := sml.GetOrder(configs.SELLInt)
			if order == nil {
				//线程休眠
				break
			}
			//订单时间大于行情时间不能进行本次撮合,(老行情不能撮合新订单)
			//if order.OrderTime >= q.QuoteTime {
			//	continue
			//}

			e.Handle.MatchQueueMap[order.MatchSymbol].RemoveOrderByObj(*order)

			order.IsHandle = true
			err := e.PutOrder(*order)
			if err != nil {
				log.Errorw("PutOrder", "error", err)
				break
			}
		}
	}
}
