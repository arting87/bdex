package deleorder

import (
	"encoding/json"
	"math"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/GyhzzZ/eos-go/token"
)

/**
Description:通过EOS订单数据构建一个DeleOrder订单结构体

Parameter:
	transferAct:eos转账action，包含转账的数据
	marketCode:订单类型 1:eos链上 2:eth链上 3:bos链上 4:eos与eth跨链
	sellTokenCotract:Token属于的合约地址

Return:
	DeleOrder订单结构体
	错误
*/
//通过EOS订单数据构建一个DeleOrder订单结构体
func BuildEosDeleOrder(transferAct token.Transfer, marketCode string, sellTokenCotract string) (DeleOrder, error) {
	eosOrder := DeleOrder{}
	err := json.Unmarshal([]byte(transferAct.Memo), &eosOrder)
	if err != nil {
		log.Errorw("Unmarshal MemoOrder", "error", err)
		return eosOrder, err
	}

	//组装EosOrder
	eosOrder.Name = string(transferAct.From)
	eosOrder.NameChainCode = strconv.Itoa(configs.INEOS)
	marketcode, err := strconv.Atoi(marketCode)
	if err != nil {
		log.Errorw("strconv.Atoi", "error", err)
		return eosOrder, err
	}
	eosOrder.MarketType = marketcode
	if eosOrder.TradeType == configs.BUYStr {
		symbol, _ := utils.ParseSymbol(eosOrder.BuyTokenSymbol)
		eosOrder.MatchSymbol = symbol + "," + strconv.Itoa(marketcode)
	} else {
		eosOrder.MatchSymbol = transferAct.Quantity.Symbol.Symbol + "," + strconv.Itoa(marketcode)
	}
	eosOrder.SellTokenAmount = strconv.Itoa(int(transferAct.Quantity.Amount))
	eosOrder.SellTokenSymbol = transferAct.Quantity.Symbol.Symbol + "," + strconv.Itoa(int(transferAct.Quantity.Symbol.Precision))
	eosOrder.SellTokenContract = sellTokenCotract
	eosOrder.BuyTokenAmount = "0"
	eosOrder.Fee = "0"
	eosOrder.WithdrawAble = true
	eosOrder.ExSuccess = false
	eosOrder.CreateAt = time.Now()
	eosOrder.UpdateAt = time.Now()

	price, err := utils.AtofPrice(eosOrder.Price)
	if err != nil {
		log.Errorw("eosPutOrder AtofPrice", "error", err)
		return eosOrder, err
	}

	if eosOrder.TradeType == configs.BUYStr {
		eosOrder.DeleVol = float64(transferAct.Quantity.Amount) / math.Pow(10, float64(transferAct.Quantity.Precision)) / price
	} else {
		eosOrder.DeleVol = float64(transferAct.Quantity.Amount) / math.Pow(10, float64(transferAct.Quantity.Precision))
	}

	return eosOrder, nil
}
