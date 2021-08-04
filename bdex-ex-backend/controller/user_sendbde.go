package controller

import (
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"fmt"
)

/**
Description:提交Bde交易对的订单

Parameter:
    req：EthPutOrderData结构体指针(前端传来的组装订单数据)

Return:
    交易哈希
    错误
*/
//提交链内卖单
func UserSendBdeOrder(req *EthPutOrderData) (string, error) {
	log := log.Context("SendBdeOrder")
	//解析订单
	fmt.Println(req.ApproveTx)
	data, err := core.DecodeBdeOrderTransaction(req.ApproveTx)
	if err != nil {
		return "", err
	}

	//组装订单
	order := deleorder.DeleOrder{}
	//order.MarketType = req.MarketType

	//构建订单数据
	//以太坊链内卖单
	err = core.BuildBdeDeleOrder(&order, data)
	if err != nil {
		return "", err
	}

	fmt.Println("------Bde-Order-----", order)
	//发送订单到链上
	txId, err := core.SendTransaction(req.ApproveTx)
	if err != nil {
		log.Errorw("发送BDE订单上链失败", "error", err)
		return order.OrderId, err
	}
	order.TxHash = txId
	order.Status = configs.PUTORDERSUCCESS

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
}
