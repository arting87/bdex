package errorder

import (
	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"time"
)

func HandleErrOrder(order deleorder.DeleOrder, txhash string, handleErr error) error {
	order.Status = configs.Error
	order.ExSuccess = true
	err := deleorder.UpdateDeleOrderById(&order)
	if err != nil {
		log.Error("HandleErrOrder update to deleordertable err", "error", err)
		return err
	}

	errOrder := ErrorOrder{
		OrderId:           order.OrderId,
		Name:              order.Name,
		NameChainCode:     order.NameChainCode,
		TradeType:         order.TradeType,
		MarketType:        order.MarketType,
		MatchSymbol:       order.MatchSymbol,
		DeleVol:           order.DeleVol,
		SellTokenSymbol:   order.SellTokenSymbol,
		SellTokenAmount:   order.SellTokenAmount,
		SellTokenContract: order.SellTokenContract,
		BuyTokenSymbol:    order.BuyTokenSymbol,
		BuyTokenAmount:    order.BuyTokenAmount,
		BuyTokenContract:  order.BuyTokenContract,
		Price:             order.Price,
		Fee:               order.Fee,
		WithdrawAble:      order.WithdrawAble,
		ExSuccess:         order.ExSuccess,
		TargetAddress:     order.TargetAddress,
		TargetChainCode:   order.TargetChainCode,
		CreateAt:          order.CreateAt,
		UpdateAt:          time.Now(),
		TxHash:            order.TxHash,
		Status:            configs.Error,
		Err:               handleErr.Error(),
	}

	err = InsertErrOrder(errOrder)
	if err != nil {
		log.Error("HandleErrOrder insert to errOrdertable err", "error", err)
		return err
	}
	return nil
}
