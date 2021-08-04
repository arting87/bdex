package matchpool

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/GyhzzZ/eos-go"
	"github.com/spf13/viper"
)

/**
Description:发送EOSSubmit上链并修改数据

Parameter:
	eosOrder:需要TransNeed的订单，DeleOrder类型
	count:这笔订单撮合了多少个自身的Selltoken的代币
	dealPrice:撮合的价格
	chainNum:发送到哪条链
	orderbId:与此订单撮合的订单Id

Return:
	交易Hash
	错误
*/
func (o *CommonOrder) SubmitEosOrder(eosOrder deleorder.DeleOrder, count, dealPrice float64, chainNum string, orderbId uint64) (string, error) {

	eosOrderB, err := deleorder.GetEOSOrderFromDbById(strconv.FormatUint(orderbId, 10))
	if err != nil {
		log.Errorw("Use GetEOSOrderFromDbById", "error", err)
		return "", err
	}

	fmt.Println("submitEosOrder,count-------:", count, ",code----:", eosOrder.SellTokenSymbol)

	name := eos.AN(viper.GetString("chain." + chainNum + ".name"))
	to := eos.AN(eosOrderB.TargetAddress)

	orderId, err := strconv.ParseUint(eosOrder.OrderId, 0, 64)
	if err != nil {
		log.Errorw("strconv.ParseUint", "error", err)
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

	fmt.Println("submitEosOrder,countAmount-------:", countAmount, ",code----:", eosOrder.SellTokenSymbol)

	dealPriceStr := strconv.FormatFloat(dealPrice*math.Pow(10, 8), 'f', -1, 64)
	submit := eosContract.NewSubmit(name, to, orderId, countAsset, eosOrder.SellTokenContract, uint64(orderbId), dealPriceStr)

	txResp, err := eosContract.BuildTx(submit, chainNum)
	if err != nil {
		log.Errorw("Send Submit to "+chainNum+" Chain", "ErrorInMem", err)
		return "", err
	}

	return txResp.TransactionID, nil
}

func (o *CommonOrder) ChangeDBOrder(eosOrder deleorder.DeleOrder, count eos.Asset, dealPrice float64) error {
	f, err := strconv.ParseFloat(eosOrder.SellTokenAmount, 64)
	if err != nil {
		log.Errorw("ParseFloat", "error", err)
		return err
	}

	sellTokenAmount := int64(f) - count.Amount

	eosOrder.SellTokenAmount = strconv.FormatInt(sellTokenAmount, 10)

	//price, err := AtofPrice(eosOrder.Price)
	//if err != nil {
	//	log.Errorw("AtofPrice", "error", err)
	//	return err
	//}

	fee := math.Ceil(float64(count.Amount)/configs.FeeBase*configs.FeePersents) / math.Pow(10, float64(count.Symbol.Precision))

	oldFee, err := strconv.ParseFloat(eosOrder.Fee, 64)
	if err != nil {
		log.Errorw("ParseFloat", "error", err)
		return err
	}

	eosOrder.Fee = strconv.FormatFloat(oldFee+fee, 'f', -1, 64)

	if eosOrder.TradeType == configs.BUYStr {
		realSellTokenAmount := float64(count.Amount - (count.Amount / configs.FeeBase * configs.FeePersents))
		realSellTokenCount := realSellTokenAmount / math.Pow(10, float64(count.Symbol.Precision))

		buyTokenAmout, err := strconv.ParseFloat(eosOrder.BuyTokenAmount, 64)
		if err != nil {
			log.Errorw("ParseFloat BuyTokenAmount", "error", err)
			return err
		}
		_, buytokenPer := utils.ParseSymbol(eosOrder.BuyTokenSymbol)

		eosOrder.BuyTokenAmount = strconv.FormatFloat(buyTokenAmout+(realSellTokenCount/dealPrice)*math.Pow(10, float64(buytokenPer)), 'f', -1, 64)
	} else {
		realSellTokenAmount := float64(count.Amount - (count.Amount / configs.FeeBase * configs.FeePersents))
		realSellTokenCount := realSellTokenAmount / math.Pow(10, float64(count.Symbol.Precision))

		buyTokenAmout, err := strconv.ParseFloat(eosOrder.BuyTokenAmount, 64)
		if err != nil {
			log.Errorw("ParseFloat BuyTokenAmount", "error", err)
			return err
		}
		_, buytokenPer := utils.ParseSymbol(eosOrder.BuyTokenSymbol)

		eosOrder.BuyTokenAmount = strconv.FormatFloat(buyTokenAmout+(realSellTokenCount*dealPrice)*math.Pow(10, float64(buytokenPer)), 'f', -1, 64)
	}

	withdrawable := true
	ex_success := false
	status := configs.UNSUCCESS

	if sellTokenAmount*3 <= 2000 {
		eosOrder.SellTokenAmount = "0"
		withdrawable = false
		ex_success = true
		status = configs.SUCCESS
	}

	eosOrder.WithdrawAble = withdrawable
	eosOrder.ExSuccess = ex_success
	eosOrder.Status = status
	eosOrder.UpdateAt = time.Now()
	fmt.Println("-----------------------------")
	if len(eosOrder.Fee) >= 18 {
		eosOrder.Fee = eosOrder.Fee[:17]
	}
	fmt.Println(eosOrder.Fee)
	fmt.Println("-----------------------------")

	err = deleorder.UpdateDeleOrderById(&eosOrder)
	if err != nil {
		log.Errorw("Update eosOrder", "error", err)
		return err
	}
	cOrder, err := ParseOrder(eosOrder)
	if err != nil {
		log.Errorw("ParseOrder", "error", err)
		return err
	}
	o.Count = cOrder.Count
	o.IsCancle = cOrder.IsCancle
	fmt.Println("**************************o.oid:", o.Oid, ",o.selltoken:", o.Count, ",o.withdrawable:", o.IsCancle)

	return nil
}
