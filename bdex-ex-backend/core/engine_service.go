package core

import (
	"strconv"

	"bdex.in/bdex/bdex-ex-backend/configs"
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/errorder"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/GyhzzZ/eos-go"
	"github.com/spf13/viper"
)

/**
Description:回滚一个订单

Parameter:
	order:需要回滚的订单结构体，CommonOrder类型

Return:
	错误
*/
func BackOneOrder(order *matchpool.CommonOrder) error {
	deleOrder, err := deleorder.GetEOSOrderFromDbById(strconv.FormatUint(order.Oid, 10))
	if err != nil {
		log.Errorw("Use GetEOSOrderFromDbById", "error", err)
		return err
	}
	log.Infow("backorder订单", "id:", order.Oid)
	var handleErr error
	txHash := ""
	switch order.MarketCode {
	case configs.INEOS:
		txHash, handleErr = BackEosOrder(order, configs.INEOSStr)
		if handleErr != nil {
			log.Error("订单数据回滚失败", "error", handleErr, "订单ID", order.Oid)
			goto ERR
		}
	case configs.INETH, configs.INBDE:
		txHash, handleErr = BackEthOrder(order.Oid)
		if handleErr != nil {
			log.Error("订单数据回滚失败", "error", handleErr, "订单ID", order.Oid)
			goto ERR
		}
	case configs.EOSETH:
		if order.BsFlag == configs.BUYInt {
			txHash, handleErr = BackEosOrder(order, configs.INEOSStr)
			if handleErr != nil {
				log.Error("订单数据回滚失败", "error", handleErr, "订单ID", order.Oid)
				goto ERR
			}
		} else {
			txHash, handleErr = BackEthOrder(order.Oid)
			if handleErr != nil {
				log.Error("订单数据回滚失败", "error", handleErr, "订单ID", order.Oid)
				goto ERR
			}
		}
		//case configs.INBDE:
		//	txHash, handleErr = BackEthOrder(order.Oid)
		//	if handleErr != nil {
		//		log.Error("订单数据回滚失败", "error", handleErr, "订单ID", order.Oid)
		//		goto ERR
		//	}
	}
	return nil

ERR:
	err = errorder.HandleErrOrder(deleOrder, txHash, handleErr)
	if err != nil {
		log.Error("HandleErrOrder错误", "error", err)
	}
	return handleErr
}

/**
Description:发送一笔订单的转账

Parameter:
	order:需要回滚的订单结构体，CommonOrder类型
	count:这笔订单撮合了多少个自身的Selltoken的代币
	dealPrice:撮合的价格

Return:
	错误
*/
func TransNeedOneOrder(order *matchpool.CommonOrder, count, dealPrice float64) error {
	deleOrder, err := deleorder.GetEOSOrderFromDbById(strconv.FormatUint(order.Oid, 10))
	if err != nil {
		log.Errorw("Use GetEOSOrderFromDbById", "error", err)
		return err
	}
	log.Infow("transfer订单:", "id:", order.Oid)
	var handleErr error
	txHash := ""

	switch order.MarketCode {
	case configs.INEOS:
		txHash, handleErr = order.TransNeedEosOrder(deleOrder, count, dealPrice, configs.INEOSStr)
		if handleErr != nil {
			log.Error("订单链上转账失败", "error", handleErr, "订单ID", order.Oid)
			goto ERR
		}
	case configs.INETH, configs.INBDE:
		txHash, handleErr = TransNeedEthOrder(order, deleOrder, count, dealPrice)
		if handleErr != nil {
			log.Error("订单链上转账失败", "error", handleErr, "订单ID", order.Oid)
			goto ERR
		}
	case configs.EOSETH:
		if order.BsFlag == configs.BUYInt {
			txHash, handleErr = order.TransNeedEosOrder(deleOrder, count, dealPrice, configs.INEOSStr)
			if handleErr != nil {
				log.Error("订单链上转账失败", "error", handleErr, "订单ID", order.Oid)
				goto ERR
			}
		} else {
			txHash, handleErr = TransNeedEthOrder(order, deleOrder, count, dealPrice)
			if handleErr != nil {
				log.Error("订单链上转账失败", "error", handleErr, "订单ID", order.Oid)
				goto ERR
			}
		}
		//case configs.INBDE:
		//	txHash, handleErr = TransNeedEthOrder(order, deleOrder, count, dealPrice)
		//	if handleErr != nil {
		//		log.Error("订单链上转账失败", "error", handleErr, "订单ID", order.Oid)
		//		goto ERR
		//	}
	}
	return nil

ERR:
	err = errorder.HandleErrOrder(deleOrder, txHash, handleErr)
	if err != nil {
		log.Error("HandleErrOrder错误", "error", err)
	}
	return handleErr
}

/**
Description:发送一笔订单的成交

Parameter:
	order:需要回滚的订单结构体，CommonOrder类型
	count:这笔订单撮合了多少个自身的Selltoken的代币
	dealPrice:撮合的价格
	orderbId:与submit撮合的订单Id

Return:
	错误
*/
func SubmitOneOrder(order *matchpool.CommonOrder, count, dealPrice float64, orderbId uint64) error {
	deleOrder, err := deleorder.GetEOSOrderFromDbById(strconv.FormatUint(order.Oid, 10))
	if err != nil {
		log.Errorw("Use GetEOSOrderFromDbById", "error", err)
		return err
	}
	var handleErr error
	log.Infow("submit订单：", "id:", order.Oid)
	txHash := ""

	switch order.MarketCode {
	case configs.INEOS:
		//EOS链上
		txHash, handleErr = order.SubmitEosOrder(deleOrder, count, dealPrice, configs.INEOSStr, orderbId)
		if handleErr != nil {
			log.Errorw("SubmitEosOrder", "error", handleErr)
			goto ERR
		}

	case configs.INETH,configs.INBDE:
		//以太坊链上
		txHash, handleErr = SubmitEthOrder(order, deleOrder, count, dealPrice, orderbId)
		if handleErr != nil {
			log.Errorf("submit eth error:", handleErr)
			goto ERR
		}
	//case :
	//	//Bde
	//	txHash, handleErr = SubmitBdeOrder(order, deleOrder, count, dealPrice, orderbId)
	//	if handleErr != nil {
	//		log.Errorf("submit eth error:", handleErr)
	//		goto ERR
	//	}

	case configs.EOSETH:
		//跨链
		if deleOrder.TradeType == configs.SELLStr { //卖单,发送至ETH
			txHash, handleErr = SubmitEthOrder(order, deleOrder, count, dealPrice, orderbId)
			if handleErr != nil {
				log.Errorf("submit eth error:", handleErr)
				goto ERR
			}
		} else { //买单,发送至EOS
			txHash, handleErr = order.SubmitEosOrder(deleOrder, count, dealPrice, "1", orderbId)
			if handleErr != nil {
				log.Errorw("SubmitEosOrder", "error", handleErr)
				goto ERR
			}
		}
	}

	//将交易成功部分插入到交易记录表中
	handleErr = matchpool.InsertTxRecordByCommOrder(order, deleOrder.Name, count, dealPrice, txHash, orderbId)
	if handleErr != nil {
		log.Errorw("InsertTxRecordByCommOrder", "error", handleErr)
		goto ERR
	}

	return nil

ERR:
	err = errorder.HandleErrOrder(deleOrder, txHash, handleErr)
	if err != nil {
		log.Error("HandleErrOrder错误", "error", err)
	}
	return handleErr
}

/**
Description:发送一笔订单的Upstate修改状态

Parameter:
	order:需要回滚的订单结构体，CommonOrder类型
	orderbId:与submit撮合的订单Id
	withdrawable:需要修改的状态为true或者false

Return:
	错误
*/
func UpstateOneOrder(order *matchpool.CommonOrder, orderbId uint64, withdrawable bool) error {

	eosOrder, err := deleorder.GetEOSOrderFromDbById(strconv.FormatUint(order.Oid, 10))
	if err != nil {
		log.Errorw("Use GetEOSOrderFromDbById", "error", err)
		return err
	}
	log.Infow("更新订单状态 内部：", "orderid", eosOrder.OrderId)
	var handleErr error
	txHash := ""
	switch order.MarketCode {
	case configs.INEOS:
		//EOS链上
		txHash, handleErr = UpstateEosOrder(eosOrder.OrderId, orderbId, configs.INEOSStr, withdrawable)
		if handleErr != nil {
			log.Errorw("UpstateEosOrder", "error", handleErr)
			goto ERR
		}
	case configs.INETH, configs.INBDE:
		//以太坊链上
		txHash, handleErr = UpstatusEthOrder(order.MarketCode, eosOrder.OrderId, withdrawable)
		if handleErr != nil {
			log.Errorf("upstate eth error:", handleErr)
			goto ERR
		}
	case configs.EOSETH:
		//跨链
		if eosOrder.TradeType == configs.SELLStr { //卖单,发送至ETH
			txHash, handleErr = UpstatusEthOrder(order.MarketCode, eosOrder.OrderId, withdrawable)
			if handleErr != nil {
				log.Errorf("upstate eth error:", handleErr)
				goto ERR
			}
		} else if eosOrder.TradeType == configs.BUYStr { //买单,发送至EOS
			txHash, handleErr = UpstateEosOrder(eosOrder.OrderId, orderbId, configs.INEOSStr, withdrawable)
			if handleErr != nil {
				log.Errorw("UpstateEosOrder", "error", handleErr)
				goto ERR
			}
		}
	}

	return nil

ERR:
	err = errorder.HandleErrOrder(eosOrder, txHash, handleErr)
	if err != nil {
		log.Error("HandleErrOrder错误", "error", err)
	}
	return handleErr
}

func UpstateEosOrder(oId string, obId uint64, chainnum string, withdrawable bool) (string, error) {
	from := eos.AN(viper.GetString("chain." + chainnum + ".name"))
	orderId, err := strconv.ParseUint(oId, 0, 64)
	if err != nil {
		return "", err
	}

	upstate := eosContract.NewUpstate(from, orderId, uint64(obId), withdrawable)

	txHash, err := eosContract.BuildTx(upstate, chainnum)
	if err != nil {
		return "", err
	}
	return txHash.TransactionID, nil
}

/**
Description:更新以太坊订单状态

Parameter:
    orderid:订单id

Return:
	错误
*/
//更新以太坊订单状态
func UpstatusEthOrder(marketCode int, orderid string, withdrawable bool) (string, error) {
	id, err := strconv.ParseUint(orderid, 10, 64)
	if err != nil {
		log.Errorf("orderid parseUint error:", err)
		return "", err
	}

	//获取账户空闲队列的第一个账户
	account := EXCore.EAL.DelFirstLeisureAccount()
	//账户使用完后再次插入队列
	defer EXCore.EAL.InsertLeisureAccount(account)

	txHash, err := EXCore.ETH.UpState(marketCode, id, withdrawable, account)
	if err != nil {
		log.Errorf("connect to eth contract method upstate error:", err)
		return txHash, err
	}
	return txHash, nil
}

func (e *Engine) DeleteOrderFromMemory(order *deleorder.DeleOrder) (bool, error) {
	orderId, err := strconv.ParseUint(order.OrderId, 10, 64)
	if err != nil {
		log.Errorw("strconv.ParseInt", "error", err)
		return false, err
	}

	//在等待队列中查找
	list := e.Handle.MatchQueueMap[order.MatchSymbol]
	element := list.GetOrderByOid(orderId)
	if element != nil {
		//删除此订单
		if element.Value.(*matchpool.CommonOrder).IsHandle == true {
			return false, nil
		}
		if order.TradeType == configs.BUYStr {
			list.List.BuyList.Remove(element)
			return true, nil
		} else {
			list.List.SellList.Remove(element)
			return true, nil
		}
	}

	//在撮合队列中查找
	e.Quote.Lock()
	defer e.Quote.Unlock()

	quote := e.Quote.Quotes[order.MatchSymbol]

	price, err := utils.AtofPrice(order.Price)
	if err != nil {
		log.Errorw("AtofPrice", "error", err)
		return false, err
	}

	if order.TradeType == configs.BUYStr {
		cOrdersLen := len(quote.Buy[price])
		for k := 0; k < cOrdersLen; {
			cOrder := quote.Buy[price][k]
			if orderId == cOrder.Oid {
				if cOrder.IsHandle == true {
					return false, nil
				}
				if k == len(quote.Buy[price])-1 {
					quote.Buy[price] = quote.Buy[price][:k]
				} else {
					quote.Buy[price] = append(quote.Buy[price][:k], quote.Buy[price][k+1:]...)
				}
				k--
				break
			}
			cOrdersLen = len(quote.Buy[price])
			k++
		}
	} else {
		cOrdersLen := len(quote.Sell[price])
		for k := 0; k < cOrdersLen; {
			cOrder := quote.Sell[price][k]
			if orderId == cOrder.Oid {
				if cOrder.IsHandle == true {
					return false, nil
				}
				if k == len(quote.Sell[price])-1 {
					quote.Sell[price] = quote.Sell[price][:k]
				} else {
					quote.Sell[price] = append(quote.Sell[price][:k], quote.Sell[price][k+1:]...)
				}
				k--
				break
			}
			cOrdersLen = len(quote.Sell[price])
			k++
		}
	}
	return true, nil
}
