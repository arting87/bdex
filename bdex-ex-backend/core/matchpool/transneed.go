package matchpool

import (
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/utils"
	"math"
	"strconv"

	"github.com/GyhzzZ/eos-go"
)

/**
Description:发送EOSTransNeed上链并修改数据

Parameter:
	eosOrder:需要TransNeed的订单，DeleOrder类型
	count:这笔订单撮合了多少个自身的Selltoken的代币
	dealPrice:撮合的价格
	chainNum:发送到哪条链

Return:
	交易Hash
	错误
*/
func (o *CommonOrder) TransNeedEosOrder(eosOrder deleorder.DeleOrder, count, dealPrice float64, chainNum string) (string, error) {
	orderId, err := strconv.ParseUint(eosOrder.OrderId, 0, 64)
	if err != nil {
		log.Errorw("strconv.ParseUint", "error", err)
		return "", err
	}

	transNeed := eosContract.NewTransNeed(orderId)

	txResp, err := eosContract.BuildTx(transNeed, chainNum)
	if err != nil {
		log.Errorw("Send transNeed to "+chainNum+" Chain", "ErrorInMem", err)
		return "", err
	}

	s, p := utils.ParseSymbol(eosOrder.SellTokenSymbol)
	countAmount := int64(count * math.Pow(10, float64(p)))

	countSymbol := eos.Symbol{
		Precision: uint8(p),
		Symbol:    s,
	}

	countAsset := eos.Asset{
		Amount: countAmount,
		Symbol: countSymbol,
	}

	//修改数据库数据
	err = o.ChangeDBOrder(eosOrder, countAsset, dealPrice)
	if err != nil {
		log.Errorw("changeEOSDBOrder"+chainNum+" Chain", "error", err)
		return "", err
	}

	return txResp.TransactionID, nil
}
