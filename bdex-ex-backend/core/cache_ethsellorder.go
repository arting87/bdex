package core

/**
Description:缓冲ETHSellOrder下单

Parameter:

Return:
*/
//缓冲ETHSellOrder下单
/*func CacherEthSellOrder() {
	var c ethsellorder.EthSellOrder
	for {
		sellOrder := EXCore.Epl.GetFirstEthPutOrder()
		if sellOrder == c {
			continue
		}
		s, _ := utils.ParseSymbol(sellOrder.SellTokenSymbol)
		deleOrder := deleorder.DeleOrder{
			OrderId:           sellOrder.OrderId,
			Name:              sellOrder.Name,
			NameChainCode:     sellOrder.NameChainCode,
			TradeType:         sellOrder.TradeType,
			MarketType:        sellOrder.MarketType,
			MatchSymbol:       s + "," + strconv.Itoa(sellOrder.MarketType),
			SellTokenSymbol:   sellOrder.SellTokenSymbol,
			SellTokenAmount:   sellOrder.SellTokenAmount,
			SellTokenContract: sellOrder.SellTokenContract,
			BuyTokenSymbol:    sellOrder.BuyTokenSymbol,
			BuyTokenAmount:    sellOrder.BuyTokenAmount,
			BuyTokenContract:  sellOrder.BuyTokenContract,
			Price:             sellOrder.Price,
			Fee:               sellOrder.Fee,
			WithdrawAble:      sellOrder.WithdrawAble,
			ExSuccess:         sellOrder.ExSuccess,
			TargetAddress:     sellOrder.TargetAddress,
			TargetChainCode:   sellOrder.TargetChainCode,
			CreateAt:          sellOrder.CreateAt,
			UpdateAt:          sellOrder.UpdateAt,
			TxHash:            sellOrder.TxHash,
			Status:            configs.PUTORDERSUCCESS,
		}

		//构造订单，发送到链上
		targetType, err := strconv.ParseBool(sellOrder.TradeType)
		if err != nil {
			log.Errorw("ParseBool", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		sellTokenAmount, err := strconv.ParseUint(sellOrder.SellTokenAmount, 10, 64)
		if err != nil {
			log.Errorw("ParseUint", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		price, err := strconv.ParseUint(sellOrder.Price, 10, 64)
		if err != nil {
			log.Errorw("ParseUint", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		deposits, err := ethereum.NewDepositsOrderSell(sellOrder.OrderId, sellOrder.Name, sellOrder.TargetAddress, targetType, sellTokenAmount, sellOrder.SellDecimals, sellOrder.SellTokenContract, price, &sellOrder.CreateAt)
		if err != nil {
			log.Errorw("NewDepositsOrderSell", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		fmt.Println("NewDepositsOrderSell :", deposits)
		txHash, err := EXCore.ETH.DepositsOrderSell(deposits)
		if err != nil || txHash == "" {
			log.Errorw("DepositsOrderSell", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
			return
		}

		deleOrder.TxHash = txHash

		//计算委托第三方代币的总量
		p, err := utils.AtofPrice(deleOrder.Price)
		if err != nil {
			log.Errorw("eosPutOrder AtofPrice", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
			return
		}

		_, money := utils.AtoAmount(deleOrder.SellTokenSymbol, deleOrder.SellTokenAmount)

		if deleOrder.TradeType == configs.BUYStr {
			deleOrder.DeleVol = money / p
		} else {
			deleOrder.DeleVol = money
		}
		//订单存入数据库
		for {
			err := deleorder.InsertDeleOrder(&deleOrder)
			if err != nil {
				log.Errorw("db insert DeleOrder", "error", err)
				deleOrder.OrderId = "-" + deleOrder.OrderId
				continue
			}
			break
		}

		err = ethsellorder.DeleteEthSellOrderById(sellOrder.OrderId)
		if err != nil {
			log.Errorw("delete  EthSellOrder error", err)
			continue
		}

		EXCore.Epl.DelFirstEthPutOrder()
		if deleOrder.Status != configs.PUTORDERFAILED {
			//触发撮合引擎
			commonOrder, err := matchpool.ParseOrder(deleOrder)
			if err != nil {
				log.Errorw("ParseOrder", "error", err)
				continue
			}

			EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

		}

	}
}*/

/*func BulkCacheEthSellOrder() {
	var c ethsellorder.EthSellOrder
	for {
		sellOrder := EXCore.Epl.GetFirstEthPutOrder()
		if sellOrder == c {
			continue
		}
		s, _ := utils.ParseSymbol(sellOrder.SellTokenSymbol)
		deleOrder := deleorder.DeleOrder{
			OrderId:           sellOrder.OrderId,
			Name:              sellOrder.Name,
			NameChainCode:     sellOrder.NameChainCode,
			TradeType:         sellOrder.TradeType,
			MarketType:        sellOrder.MarketType,
			MatchSymbol:       s + "," + strconv.Itoa(sellOrder.MarketType),
			SellTokenSymbol:   sellOrder.SellTokenSymbol,
			SellTokenAmount:   sellOrder.SellTokenAmount,
			SellTokenContract: sellOrder.SellTokenContract,
			BuyTokenSymbol:    sellOrder.BuyTokenSymbol,
			BuyTokenAmount:    sellOrder.BuyTokenAmount,
			BuyTokenContract:  sellOrder.BuyTokenContract,
			Price:             sellOrder.Price,
			Fee:               sellOrder.Fee,
			WithdrawAble:      sellOrder.WithdrawAble,
			ExSuccess:         sellOrder.ExSuccess,
			TargetAddress:     sellOrder.TargetAddress,
			TargetChainCode:   sellOrder.TargetChainCode,
			CreateAt:          sellOrder.CreateAt,
			UpdateAt:          sellOrder.UpdateAt,
			TxHash:            sellOrder.TxHash,
			Status:            configs.PUTORDERSUCCESS,
		}

		//构造订单，发送到链上
		targetType, err := strconv.ParseBool(sellOrder.TradeType)
		if err != nil {
			log.Errorw("ParseBool", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		sellTokenAmount, err := strconv.ParseUint(sellOrder.SellTokenAmount, 10, 64)
		if err != nil {
			log.Errorw("ParseUint", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		price, err := strconv.ParseUint(sellOrder.Price, 10, 64)
		if err != nil {
			log.Errorw("ParseUint", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		deposits, err := ethereum.NewDepositsOrderSell(sellOrder.OrderId, sellOrder.Name, sellOrder.TargetAddress, targetType, sellTokenAmount, sellOrder.SellDecimals, sellOrder.SellTokenContract, price, &sellOrder.CreateAt)
		if err != nil {
			log.Errorw("NewDepositsOrderSell", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
		}

		fmt.Println("NewDepositsOrderSell :", deposits)
		txHash, err := EXCore.ETH.DepositsOrderSell(deposits)
		if err != nil || txHash == "" {
			log.Errorw("DepositsOrderSell", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
			return
		}

		deleOrder.TxHash = txHash

		//计算委托第三方代币的总量
		p, err := utils.AtofPrice(deleOrder.Price)
		if err != nil {
			log.Errorw("eosPutOrder AtofPrice", "error", err)
			deleOrder.WithdrawAble = false
			deleOrder.ExSuccess = true
			deleOrder.Status = configs.PUTORDERFAILED
			return
		}

		_, money := utils.AtoAmount(deleOrder.SellTokenSymbol, deleOrder.SellTokenAmount)

		if deleOrder.TradeType == configs.BUYStr {
			deleOrder.DeleVol = money / p
		} else {
			deleOrder.DeleVol = money
		}
		//订单存入数据库
		for {
			err := deleorder.InsertDeleOrder(&deleOrder)
			if err != nil {
				log.Errorw("db insert DeleOrder", "error", err)
				deleOrder.OrderId = "-" + deleOrder.OrderId
				continue
			}
			break
		}

		err = ethsellorder.DeleteEthSellOrderById(sellOrder.OrderId)
		if err != nil {
			log.Errorw("delete  EthSellOrder error", err)
			continue
		}

		EXCore.Epl.DelFirstEthPutOrder()
		if deleOrder.Status != configs.PUTORDERFAILED {
			//触发撮合引擎
			commonOrder, err := matchpool.ParseOrder(deleOrder)
			if err != nil {
				log.Errorw("ParseOrder", "error", err)
				continue
			}

			EXCore.Engine.Handle.GetInnerSktList(commonOrder.MatchSymbol).InsertOrder(commonOrder)

		}
	}
}*/
