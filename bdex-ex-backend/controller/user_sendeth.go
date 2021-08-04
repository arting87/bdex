package controller

import (
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/utils"
	"fmt"
)

/*
与user sendeth 请求有关的结构体
*/

// Eth put order api
type EthPutOrderReq struct {
	Namespace string          `json:"namespace"`
	Action    string          `json:"action"`
	Data      EthPutOrderData `json:"data"`
}

type EthPutOrderData struct {
	//TradeType  bool   `json:"tradeType"`
	//MarketType int    `json:"marketType"`
	ApproveTx string `json:"approveTx"`
}

type EthPutOrderResp struct {
	OrderId string `json:"orderId"`
	Status  bool   `json:"status"`
}

func UserSendEthOrder(req *EthPutOrderData) (string, error) {
	log := log.Context("SendEthOrder")
	//解析订单
	fmt.Println(req.ApproveTx)
	data, err := core.DecodeEthTransactionDepositsOrder(req.ApproveTx)
	if err != nil {
		log.Errorw("DecodeEthTransactionDepositsOrder", "error", err)
		return "", err
	}
	//组装订单
	order := deleorder.DeleOrder{}
	//构建订单数据
	err = core.BuildEthDeleOrder(&order, data)
	if err != nil {
		log.Errorw("构建订单数据失败", "error", err.Error())
		return "", err
	}

	fmt.Println("eth------Order-----", order)
	//发送订单到链上
	txid, err := core.SendTransaction(req.ApproveTx)
	if err != nil {
		log.Errorw("DepositsOrderBuy", "error", err)
		return order.OrderId, err
	}
	order.TxHash = txid
	order.Status = configs.PUTORDERSUCCESS

	//计算委托第三方代币的总量
	price, err := utils.AtofPrice(order.Price)
	if err != nil {
		log.Errorw("eosPutOrder AtofPrice", "error", err)
		return order.OrderId, err
	}

	_, money := utils.AtoAmount(order.SellTokenSymbol, order.SellTokenAmount)

	if (data.SellContract.String() == configs.AddressNil && data.BuyContract.String() == configs.AddressNil) {
		//跨链eth/eos
		order.DeleVol = money
	} else {
		if data.TradeType {
			order.DeleVol = money
		} else {
			order.DeleVol = money / price
		}
	}

	/*if order.TradeType == configs.BUYStr {
		order.DeleVol = money / price
	} else {
		order.DeleVol = money
	}*/
	//插入数据库
	err = deleorder.InsertDeleOrder(&order)
	if err != nil {
		log.Errorw("db insert DeleOrder", "error", err)
		return order.OrderId, err
	}
	//触发撮合引擎
	commonOrder, err := matchpool.ParseOrder(order)
	if err != nil {
		log.Errorw("ParseOrder", "error", err)
		return order.OrderId, err
	}

	//放入等待队列
	core.EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

	return order.OrderId, nil
}

/**
Description://提交链内卖单

Parameter:
    req：EthPutOrderData结构体指针(前端传来的组装订单数据)

Return:
    交易哈希
    错误
*/
//提交链内卖单
//func UserSendEthSellOrder(req *EthPutOrderData) (string, error) {
//	log := log.Context("SendEthSellOrder")
//	//从数据库获取小数点位数
//	decimal, err := tokens.GetEthDecimalsFromMem(req.SellTokenAddress)
//	if err != nil {
//		log.Errorw("getEthDecimals", "error", err)
//		return req.OrderId, err
//	}
//	//组装订单
//	order, err := BuildEthSellOrder(req, decimal)
//
//	fmt.Println("--------approve-----tx", req.ApproveTx)
//	//调用erc20给交易所授权
//	txHash, err := core.SendTransaction(req.ApproveTx)
//	if err != nil {
//		log.Errorw("DepositsOrderApprove", "error", err)
//		return req.OrderId, err
//	}
//	order.TxHash = txHash
//	log.Debug("DepositsOrderApprove susscee:", order)
//	//sellOrder插入数据库
//	err = ethsellorder.InsertEthSellOrder(&order)
//	if err != nil {
//		log.Errorw("EthSellOrder table insert", "error", err)
//		return req.OrderId, err
//	}
//
//	//插入EthSellOrderList缓冲处理
//	core.EXCore.Epl.InsertEthSellOrder(order)
//
//	//返回前端数据
//	return req.OrderId, nil
//}

/**
Description://提交链内卖单

Parameter:
    req：EthPutOrderData结构体指针(前端传来的组装订单数据)

Return:
    交易哈希
    错误
*/
//提交链内卖单
/*func UserSendEthSellOrder(req *EthPutOrderData) (string, error) {
	log := log.Context("SendEthSellOrder")
	//解析订单
	fmt.Println(req.ApproveTx)
	_, data, err := core.DecodeEthTransactionDepositsSell(req.ApproveTx)
	if err != nil {
		log.Errorw("DecodeEthTransactionDepositsSell", "error", err)
		return "", err
	}

	//组装订单
	order := deleorder.DeleOrder{}
	order.MarketType = req.MarketType

	//构建订单数据
	//以太坊链内卖单
	err = core.BuildEthDeleOrderSell(&order, data)
	if err != nil {
		log.Errorw("构建卖单数据失败", "error", err)
		return "", err
	}

	fmt.Println("eth------sell-----", &order)
	//发送订单到链上
	txId, err := core.SendTransaction(req.ApproveTx)
	if err != nil {
		return order.OrderId, err
	}
	order.TxHash = txId
	order.Status = configs.PUTORDERSUCCESS

	//计算委托第三方代币的总量
	price, err := utils.AtofPrice(order.Price)
	if err != nil {
		log.Errorw("utils.AtofPrice", "error", err)
		return order.OrderId, err
	}

	_, money := utils.AtoAmount(order.SellTokenSymbol, order.SellTokenAmount)

	if order.TradeType == configs.BUYStr {
		order.DeleVol = money / price
	} else {
		order.DeleVol = money
	}

	//插入数据库
	err = deleorder.InsertDeleOrder(&order)
	if err != nil {
		log.Errorw("db insert DeleOrder", "error", err)
		return order.OrderId, err
	}

	//触发撮合引擎
	commonOrder, err := matchpool.ParseOrder(order)
	if err != nil {
		log.Errorw("ParseOrder", "error", err)
		return order.OrderId, err
	}

	//放入等待队列
	core.EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

	//返回前端数据
	return order.OrderId, nil
}*/

/**
Description://提交链内买单和跨链卖单

Parameter:
    req：EthPutOrderData结构体指针(前端传来的订单数据)

Return:
    交易哈希
    错误
*/
//提交链内买单和跨链卖单
/*func UserSendEthBuyOrder(req *EthPutOrderData) (string, error) {
	log := log.Context("SendEthBuyOrder")
	//解析订单
	fmt.Println(req.ApproveTx)
	_, data, err := core.DecodeEthTransactionDepositsBuy(req.ApproveTx)
	if err != nil {
		log.Errorw("DecodeEthTransactionDepositsBuy", "error", err)
		return "", err
	}
	//组装订单
	order := deleorder.DeleOrder{}
	order.MarketType = req.MarketType

	//构建订单数据
	err = core.BuildEthDeleOrderBuy(&order, data)
	if err != nil {
		log.Errorw("构建买单数据失败", "error", err)
		return "", err
	}

	fmt.Println("eth------buy-----", &order)
	//发送订单到链上
	txid, err := core.SendTransaction(req.ApproveTx)
	if err != nil {
		log.Errorw("DepositsOrderBuy", "error", err)
		return order.OrderId, err
	}
	order.TxHash = txid
	order.Status = configs.PUTORDERSUCCESS

	//计算委托第三方代币的总量
	price, err := utils.AtofPrice(order.Price)
	if err != nil {
		log.Errorw("eosPutOrder AtofPrice", "error", err)
		return order.OrderId, err
	}

	_, money := utils.AtoAmount(order.SellTokenSymbol, order.SellTokenAmount)

	if order.TradeType == configs.BUYStr {
		order.DeleVol = money / price
	} else {
		order.DeleVol = money
	}

	//插入数据库
	err = deleorder.InsertDeleOrder(&order)
	if err != nil {
		log.Errorw("db insert DeleOrder", "error", err)
		return order.OrderId, err
	}

	//触发撮合引擎
	commonOrder, err := matchpool.ParseOrder(order)
	if err != nil {
		log.Errorw("ParseOrder", "error", err)
		return order.OrderId, err
	}

	//放入等待队列
	core.EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

	return order.OrderId, nil
}*/

/**
Description:通过ETH订单数据构建一个DeleOrder订单结构体

Parameter:
	req:请求发送的订单信息
	decimal:token精度

Return:
	EthSellOrder订单结构体
	错误
*/
//通过数据组装EthSellOrder结构体
/*func BuildEthSellOrder(req *EthPutOrderData, decimal uint) (ethsellorder.EthSellOrder, error) {

	order := ethsellorder.EthSellOrder{}
	order.OrderId = req.OrderId
	order.SellDecimals = decimal
	order.Name = req.UserAddress
	order.MarketType = int(req.MarketType)
	order.SellTokenSymbol = req.SellTokenSymbol + ",8"
	order.SellTokenAmount = strconv.FormatUint(req.SellTokenAmount, 10)
	order.SellTokenContract = req.SellTokenAddress
	order.BuyTokenSymbol = "ETH" + ",8"
	order.BuyTokenAmount = strconv.FormatUint(0, 10)
	order.BuyTokenContract = configs.AddressNil
	order.Price = strconv.FormatUint(req.Price, 10)
	order.Fee = strconv.FormatFloat(0, 'f', -1, 64) //数据库存float类型
	order.WithdrawAble = true
	order.ExSuccess = false
	order.TargetAddress = req.TargetAddress
	order.CreateAt = time.Now()
	order.UpdateAt = time.Now()
	order.TradeType = strconv.FormatInt(1, 10)

	return order, nil
}*/

