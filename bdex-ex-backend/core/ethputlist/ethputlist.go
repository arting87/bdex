package ethputlist

import (
	"container/list"
	"sync"

	"bdex.in/bdex/bdex-ex-backend/ethsellorder"
	"bdex.in/bdex/bdex-ex-backend/log"
)

type EthPutList struct {
	List list.List
	l    sync.RWMutex //single write ,more read  slock
	//记录每个节点指针，用于直接删除
	IDMap map[string]*list.Element // k:orderId  v:*order
}

func (etl *EthPutList) InitList() {
	etl.IDMap = make(map[string]*list.Element)
}

//向队列中插入EthSellOrder
func (etl *EthPutList) InsertEthSellOrder(o ethsellorder.EthSellOrder) {
	log.Infow("插入到以太坊卖单缓冲队列的" + o.Price + "价格中")
	etl.l.Lock()
	defer etl.l.Unlock()
	oElement := etl.List.PushBack(o)
	etl.IDMap[o.OrderId] = oElement
	log.Infow("当前EthPutSellOrder队列", "大小", etl.List.Len())
	log.Infof("当前IDMap:%s %s:%s", "内容", etl.IDMap)
}

//获取买/卖队列最后一个订单
func (etl *EthPutList) GetFirstEthPutOrder() (c ethsellorder.EthSellOrder) {
	if etl.List.Len() == 0 {
		//log.Infow("GetOrder", "error", "buy mark is empty-------")
	} else {
		return etl.List.Front().Value.(ethsellorder.EthSellOrder)
	}

	return
}

//删除买或卖队列第一个订单
func (etl *EthPutList) DelFirstEthPutOrder() {
	etl.l.Lock()
	defer etl.l.Unlock()
	if etl.List.Len() == 0 {
		log.Infow("DelOrder", "error", "buy mark is empty-------")
	} else {
		b := etl.List.Front()
		etl.List.Remove(b)
		delete(etl.IDMap, b.Value.(ethsellorder.EthSellOrder).OrderId)
	}
}

//通过OrderId获取订单
func (etl *EthPutList) GetEthPutOrderByOid(oid string) (c *list.Element) {
	return etl.IDMap[oid]
}

//删除订单通过订单
func (etl *EthPutList) RemoveOrderByObj(c ethsellorder.EthSellOrder) bool {
	etl.l.Lock()
	defer etl.l.Unlock()
	for {
		//根据订单号查询出待删除的订单
		delObj := etl.GetEthPutOrderByOid(c.OrderId)
		if delObj == nil {
			return true
		}
		if etl.List.Len() == 0 {
			break
		} else {
			o := etl.List.Remove(delObj)
			delete(etl.IDMap, c.OrderId)
			log.Infof("删除EthPutSellOrder队列中元素:%s %s:%v", o.(*ethsellorder.EthSellOrder).OrderId)
			log.Infof("当前EthPutSellOrder队列:%s %s:%d", "大小", etl.List.Len())
			log.Infof("当前IDMap:%s %s:%s", "内容", etl.IDMap)
			return true
		}

	}
	return true
}

func (etl *EthPutList) RemoveAll() {
	var n *list.Element
	for e := etl.List.Front(); e != nil; e = n {
		n = e.Next()
		etl.List.Remove(e)
	}
	for k, _ := range etl.IDMap {
		delete(etl.IDMap, k)
	}

	log.Infow("buy/sell queue  clean empty-------")
}

func LoadEthSellOrderFromDB() (*EthPutList, error) {
	orders, err := ethsellorder.GetAllEthSellOrder()
	if err != nil {
		log.Errorw("EthSellOrder Table find", "error", err)
		return nil, err
	}

	epl := EthPutList{}
	epl.InitList()

	for _, sellOrder := range orders {
		epl.InsertEthSellOrder(sellOrder)
	}
	return &epl, nil
}
