package controller

import (
	"errors"
	"fmt"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
)

/*
与user withdraweth 请求有关的结构体
*/
//eth withdraw order api
type EthWithdrawOrderReq struct {
	Namespace string            `json:"namespace"`
	Action    string            `json:"action"`
	Data      WithdrawOrderData `json:"data"`
}

type EthWithdrawOrderResp struct {
	OrderId string `json:"orderId"`
	Status  bool   `json:"status"`
}

/**
Description:ETH撤销订单

Parameter:
    req：WithdrawOrderData结构体指针(前端传来的要撤销的订单数据)

Return:
    交易哈希
    错误
*/
//ETH撤销订单
func UserRecallEthOrder(req *WithdrawOrderData) (string, error) {
	//解析订单
	orderid, err := core.DecodeEthTransactionWithdraw(req.ApproveTx)
	log.Debug("analysis order :", orderid)
	if err != nil {
		log.Errorw("DecodeEthTransactionWithdraw", "error", err)
		return req.OrderId, err
	}

	//触发撮合引擎处理此订单
	order, err := deleorder.GetOrderFromDbById(req.OrderId)
	if err != nil {
		log.Errorw("GetOrderFromDbById", "error", err)
		return req.OrderId, err
	}

	//查询交易是否正在处理,删除内存中订单
	b, err := core.EXCore.Engine.DeleteOrderFromMemory(order)
	if err != nil {
		log.Errorw("DeleteOrderFromMemory", "error", err)
		return req.OrderId, err
	}

	if !b {
		log.Errorw("DeleteOrderFromMemory", "error", "Order is Handle")
		return req.OrderId, errors.New("this Order is Handling")
	}

	//发送订单到链上
	_, err = core.SendTransaction(req.ApproveTx)
	if err != nil {
		log.Errorw("DepositsOrderSell", "error", err)
		return req.OrderId, err
	}

	//修改数据库此订单的状态
	order.SellTokenAmount = "-" + order.SellTokenAmount
	order.WithdrawAble = false
	order.ExSuccess = true
	order.Status = configs.CANCLE

	ccc := deleorder.DeleOrder{OrderId: order.OrderId, SellTokenAmount: order.SellTokenAmount, BuyTokenAmount: order.BuyTokenAmount, Fee: order.Fee, WithdrawAble: order.WithdrawAble, ExSuccess: order.ExSuccess, UpdateAt: order.UpdateAt, Status: order.Status}
	fmt.Println("cancleorder----:", ccc)
	err = deleorder.UpdateDeleOrderById(&ccc)
	if err != nil {
		log.Errorw("DeleOrder Table update", "error", err)
		return req.OrderId, err
	}

	return req.OrderId, nil
}
