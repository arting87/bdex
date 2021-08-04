package controller

import (
	"errors"

	"bdex.in/bdex/bdex-ex-backend/configs"
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/deleorder"

	"github.com/GyhzzZ/eos-go"
)

/*
与user withdraweos 请求有关的结构体
*/
//  eos撤单api
type WithdrawOrderData struct {
	OrderId   string `json:"orderId"`
	ApproveTx string `json:"approveTx"`
}

type EosWithdrawOrderReq struct {
	Namespace string            `json:"namespace"`
	Action    string            `json:"action"`
	Data      WithdrawOrderData `json:"data"`
}

type EosWithdrawOrderResp struct {
	OrderId string `json:"orderId"`
	Status  bool   `json:"status"`
}

func UserRecallEosOrder(req *WithdrawOrderData) (string, error) {
	log := log.Context("recall eos order")

	//解包交易
	body, bytes, err := core.UnPackEosTransaction(req.ApproveTx)
	if err != nil {
		log.Errorw("UnPackEosTransaction", "error", err)
		return req.OrderId, err
	}

	//解包 withdraw action
	withdrawAct := eosContract.Withdraw{}
	err = eos.NewDecoder(bytes).Decode(&withdrawAct)
	if err != nil {
		log.Errorw("Decode Withdraw", "error", err)
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

	//发送交易到链上
	_, err = core.SendEosTransaction(body.Transaction, body.Signatures)
	if err != nil {
		log.Errorw("Send Transaction", "error", err)
		return req.OrderId, err
	}

	//修改数据库此订单的状态
	order.SellTokenAmount = "-" + order.SellTokenAmount
	order.WithdrawAble = false
	order.ExSuccess = true
	order.Status = configs.CANCLE
	err = deleorder.UpdateDeleOrderById(order)
	if err != nil {
		log.Errorw("DeleOrder Table update", "error", err)
		return req.OrderId, err
	}

	return req.OrderId, nil
}
