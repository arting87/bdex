package matchpool

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"

	"bdex.in/bdex/bdex-ex-backend/configs"
)

type StkMatchList struct {
	List MatchList
	l    sync.RWMutex //single write ,more read  slock
	//记录每个节点指针，用于直接删除
	IDMap map[uint64]*list.Element // k:orderId  v:*order
	Code  string
}

func (stm *StkMatchList) InitList(pcode string) {
	stm.Code = pcode
	stm.IDMap = make(map[uint64]*list.Element)
}

//向队列中插入CommonOrder
func (stm *StkMatchList) InsertOrder(o *CommonOrder) {
	//log := log.Named("InsertOrder")
	fmt.Println("++++++++++++++++++++InsertToMatchPool+++++++++++++++++++++++++++")
	log.Infow("插入到"+stm.Code+"订单池队列的订单为", "id:", o.Oid, "价格:", o.Price, "数量:", o.Count, "代币名称:", o.MatchSymbol, "买卖类型:", o.BsFlag)
	stm.l.Lock()
	defer stm.l.Unlock()
	if o.BsFlag == configs.BUYInt { //买 价格从低到高 时间由新及旧
		var el *list.Element
		if stm.List.BuyList.Len() == 0 { //list为空
			oElement := stm.List.BuyList.PushBack(o)
			stm.IDMap[o.Oid] = oElement
		} else {
			e := stm.List.BuyList.Front()
			for {
				if e.Value.(*CommonOrder).Price > o.Price {
					oElement := stm.List.BuyList.InsertBefore(o, e)
					stm.IDMap[o.Oid] = oElement
					break
				} else if e.Value.(*CommonOrder).Price == o.Price { //相同价格比较时间
					if e.Value.(*CommonOrder).OrderTime < o.OrderTime {
						oElement := stm.List.BuyList.InsertBefore(o, e)
						stm.IDMap[o.Oid] = oElement
						break
					}
				}
				el = e.Next()
				if el == nil { //结束
					oElement := stm.List.BuyList.InsertAfter(o, e)
					stm.IDMap[o.Oid] = oElement
					break
				}
				e = el
			}
		}

	} else if o.BsFlag == configs.SELLInt { //卖  价格从高到低 时间由新及旧
		var el *list.Element
		if stm.List.SellList.Len() == 0 { //list为空
			oElement := stm.List.SellList.PushBack(o)
			stm.IDMap[o.Oid] = oElement
		} else {
			e := stm.List.SellList.Front()
			for {
				if e.Value.(*CommonOrder).Price < o.Price {
					oElement := stm.List.SellList.InsertBefore(o, e)
					stm.IDMap[o.Oid] = oElement
					break
				} else if e.Value.(*CommonOrder).Price == o.Price {
					if e.Value.(*CommonOrder).OrderTime < o.OrderTime { //相同价格比较时间
						oElement := stm.List.SellList.InsertBefore(o, e)
						stm.IDMap[o.Oid] = oElement
						break
					}
				}
				el = e.Next()
				if el == nil { //结束
					oElement := stm.List.SellList.InsertAfter(o, e)
					stm.IDMap[o.Oid] = oElement
					break
				}
				e = el
			}
		}
	}
	log.Infow("插入后当前买队列", "Symbol", stm.Code, "大小", stm.List.BuyList.Len())
	log.Infow("插入后当前卖队列", "Symbol", stm.Code, "大小", stm.List.SellList.Len())
}

//删除买或卖队列最后一个订单
func (stm *StkMatchList) DelOrder(i int32) {
	stm.l.Lock()
	defer stm.l.Unlock()
	if i == configs.BUYInt { //买
		if stm.List.BuyList.Len() == 0 {
			log.Infow("DelOrder", "error", "buy mark is empty-------")
		} else {
			b := stm.List.BuyList.Back()
			stm.List.BuyList.Remove(b)
		}

	} else if i == configs.SELLInt { //卖
		if stm.List.SellList.Len() == 0 {
			log.Infow("DelOrder", "error", "sell mark is empty-------")
		} else {
			b := stm.List.SellList.Back()
			stm.List.SellList.Remove(b)
		}
	} else {
		log.Errorw("DelOrder", "error", "buy/sell mark is wrong-------")
	}
	log.Infow("删除后当前买队列", "Symbol", stm.Code, "大小", stm.List.BuyList.Len())
	log.Infow("删除后当前卖队列", "Symbol", stm.Code, "大小", stm.List.SellList.Len())
}

//获取买/卖队列最后一个订单
func (stm *StkMatchList) GetOrder(i int32) (c *CommonOrder) {
	if i == configs.BUYInt { //买
		if stm.List.BuyList.Len() == 0 {
			log.Infow("GetOrder", "error", "buy mark is empty-------")
		} else {
			return stm.List.BuyList.Back().Value.(*CommonOrder)
		}

	} else if i == configs.SELLInt { //卖
		if stm.List.SellList.Len() == 0 {
			log.Infow("GetOrder", "error", "sell mark is empty-------")
		} else {
			return stm.List.SellList.Back().Value.(*CommonOrder)
		}
	} else {
		log.Errorw("GetOrder", "error", "buy/sell mark is wrong-------")
	}
	return nil
}

//通过OrderId获取订单
func (stm *StkMatchList) GetOrderByOid(oid uint64) (c *list.Element) {
	return stm.IDMap[oid]
}

//删除订单通过订单
func (stm *StkMatchList) RemoveOrderByObj(c CommonOrder) bool {
	//log := log.Named("RemoveOrderByObj")
	stm.l.Lock()
	defer stm.l.Unlock()
	for {
		//根据订单号查询出待删除的订单
		delObj := stm.GetOrderByOid(c.Oid)
		if delObj == nil {
			return true
		}
		if c.BsFlag == configs.BUYInt { //买
			if stm.List.BuyList.Len() == 0 {
				break
			} else {
				o := stm.List.BuyList.Remove(delObj)
				delete(stm.IDMap, c.Oid)
				log.Infow("删除订单"+strconv.FormatUint(o.(*CommonOrder).Oid, 10)+"后当前买队列", "Symbol", stm.Code, "大小", stm.List.BuyList.Len())
				log.Infow("删除订单"+strconv.FormatUint(o.(*CommonOrder).Oid, 10)+"后当前卖队列", "Symbol", stm.Code, "大小", stm.List.SellList.Len())
				return true
			}

		} else if c.BsFlag == configs.SELLInt { //卖
			if stm.List.SellList.Len() == 0 {
				break
			} else {
				o := stm.List.SellList.Remove(delObj)
				delete(stm.IDMap, c.Oid)
				log.Infow("删除订单"+strconv.FormatUint(o.(*CommonOrder).Oid, 10)+"后当前买队列", "Symbol", stm.Code, "大小", stm.List.BuyList.Len())
				log.Infow("删除订单"+strconv.FormatUint(o.(*CommonOrder).Oid, 10)+"后当前卖队列", "Symbol", stm.Code, "大小", stm.List.SellList.Len())
				return true
			}
		} else {
			log.Errorw("RemoveOrderByObj", "error", "buy/sell mark is wrong-------")
			break
		}
	}
	return true
}

func (stm *StkMatchList) RemoveAll() {
	var n *list.Element
	for e := stm.List.BuyList.Front(); e != nil; e = n {
		n = e.Next()
		stm.List.BuyList.Remove(e)
	}

	for e := stm.List.SellList.Front(); e != nil; e = n {
		n = e.Next()
		stm.List.SellList.Remove(e)
	}
	for k, _ := range stm.IDMap {
		delete(stm.IDMap, k)
	}

	log.Infow("buy/sell queue " + stm.Identify() + " clean empty-------")
}

func (self *StkMatchList) Identify() string {
	return self.Code
}
