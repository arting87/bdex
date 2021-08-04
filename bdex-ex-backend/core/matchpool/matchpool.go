package matchpool

import (
	"sync"

	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/utils"
)

type MatchPool struct {
	MatchQueueMap map[string]*StkMatchList //k:marketcode key:Code v:commonOrderlist
	sync.RWMutex
}

func NewMatchPool() *MatchPool {
	return &MatchPool{MatchQueueMap: make(map[string]*StkMatchList)}
}

//func (this MatchPool) BindMarketMatchQueue(marketSymbol string) {
//	this.Lock()
//	defer this.Unlock()
//	InitList
//	this.MatchQueueMap[marketSymbol] = qs
//}

func (this *MatchPool) GetMatchQueue() *map[string]*StkMatchList {
	this.Lock()
	defer this.Unlock()
	if this.MatchQueueMap == nil {
		this.MatchQueueMap = make(map[string]*StkMatchList)
	}
	return &this.MatchQueueMap
}

func (this *MatchPool) GetInnerSktList(marketSymbol string) *StkMatchList {
	if this.MatchQueueMap[marketSymbol] == nil {
		this.Lock()
		s, p := utils.ParseSymbol(marketSymbol)
		log.Infow("创建一个新的订单池队列", "market:", p, ",code:", s)
		stkMatchList := &StkMatchList{}
		stkMatchList.InitList(s)
		this.MatchQueueMap[marketSymbol] = stkMatchList
		this.Unlock()
	}
	return this.MatchQueueMap[marketSymbol]
}

func (this *MatchPool) GetAllMatchQueue() map[string]*StkMatchList {

	return this.MatchQueueMap
}

func LoadOrderFromDB(matchPool *MatchPool) error {
	orders, err := deleorder.GetAllUnfinishedDeleOrder()
	if err != nil {
		log.Errorw("DeleOrder table orderby find", "error", err)
		return err
	}

	for _, order := range orders {
		cOrder, err := ParseOrder(order)
		if err != nil {
			log.Errorw("ParseOrder", "error", err)
			return err
		}

		matchPool.GetInnerSktList(cOrder.MatchSymbol).InsertOrder(cOrder)

		matchPool.GetInnerSktList(cOrder.MatchSymbol).Code = cOrder.Code
	}

	return nil
}
