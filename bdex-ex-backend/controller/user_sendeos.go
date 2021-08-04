package controller

import (
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"fmt"

	"github.com/GyhzzZ/eos-go"
	"github.com/GyhzzZ/eos-go/token"
)

/*
与user sendeos 请求有关的结构体
*/
//eos下单api
type EosPutOrderData struct {
	Order      string `json:"order"`
	MarketCode string `json:"marketCode"`
	ChainCode  string `json:"chainCode"`
	TargetCode string `json:"targetCode"`
}

type EosPutOrderReq struct {
	Namespace string          `json:"namespace"`
	Action    string          `json:"action"`
	Data      EosPutOrderData `json:"data"`
}

type EosPutOrderResp struct {
	OrderId string `json:"orderId"`
	Status  bool   `json:"status"`
}

func UserSendEOSOrder(req *EosPutOrderData) (string, error) {
	log := log.Context("User send eos order")

	//解包交易
	body, bytes, err := core.UnPackEosTransaction(req.Order)
	if err != nil {
		log.Errorw("UnPackEosTransaction", "error", err)
		return "", err
	}

	//解包 transfer action
	transferAct := token.Transfer{}
	err = eos.NewDecoder(bytes).Decode(&transferAct)
	if err != nil {
		log.Errorw("Decode Transfer", "error", err)
		return "", err
	}

	/*{                             //1
		"i": "123456789",           //16  order_id
		"a": "0或1",                //11  trade_type
		"p": 0000000000000,        //21  price
	    "bs": (ABCD,4),             //    buy_token_symbol 4是精度
		"bc": "abcd12345abc",       //20  buy_token_contract
		"ta": "abcd12345abc"        //19  target_address
	  }*/
	//组装DeleOrder
	eosOrder, err := deleorder.BuildEosDeleOrder(transferAct, req.MarketCode, string(body.Transaction.Actions[0].Account))
	if err != nil {
		log.Errorw("BuildEosDeleOrder", "error", err)
		return "", err
	}

	//发送交易到链上
	txresp, err := core.SendEosTransaction(body.Transaction, body.Signatures)
	if err != nil {
		log.Errorw("Send Transaction", "error", err)
		return "", err
	}

	fmt.Println(*txresp)
	for key, cosole := range txresp.Processed.ActionTraces {
		fmt.Println("发送订单上链的返回打印console,下标：", key, "   ", cosole.Console)
	}

	eosOrder.TxHash = txresp.TransactionID
	eosOrder.Status = configs.PUTORDERSUCCESS

	//存入数据库
	err = deleorder.InsertDeleOrder(&eosOrder)
	if err != nil {
		log.Errorw("db insert DeleOrder", "error", err)
		return "", err
	}

	//解析eosOrder
	commonOrder, err := matchpool.ParseOrder(eosOrder)
	if err != nil {
		log.Errorw("ParseOrder", "error", err)
		return "", err
	}

	//放入等待队列
	core.EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

	return eosOrder.OrderId, nil
}
