package controller

import (
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/txrecords"
	"bdex.in/bdex/bdex-ex-backend/utils"
)

/*
与api get 请求有关的结构体
*/
// 行情订单请求
type QuoteApiReq struct {
	Namespace string `json:"namespace"`
	Action    string `json:"action"`
}

type ContractAddress struct {
	Address string `json:"address"`
	Decimal int8   `json:"decimal"`
}

type MarketApi struct {
	Price        float64         `json:"price"`
	HighestPrice float64         `json:"highestPrice"`
	LowestPrice  float64         `json:"lowestPrice"`
	BasePrice    float64         `json:"basePrice"`
	Volume       float64         `json:"volume"`
	Address      ContractAddress `json:"contractAddress"`
}

type QuoteData struct {
	MarketCode string    `json:"marketCode"`
	TokenName  string    `json:"tokenName"`
	Market     MarketApi `json:"marketApi"`
	Order      OrderApi  `json:"orderApi"`
}

type OrderApi struct {
	Buy      []ApiOrder      `json:"buyOrder"`
	Sell     []ApiOrder      `json:"sellOrder"`
	Complete []CompleteOrder `json:"completeOrder"` //已完成的买单
}

type ApiOrder struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amt"`
}

type CompleteOrder struct {
	Price  float64   `json:"price"`
	Amount float64   `json:"amt"`
	Date   time.Time `json:"date"`
}

type QuoteApiResp struct {
	Namespace string      `json:"namespace"`
	Action    string      `json:"action"`
	Quote     []QuoteData `json:"data"`
}

var AllQuoteData []QuoteData

/**
Description:获取到所有代币的行情以及订单信息

Parameter:

Return:
	行情以及订单信息数组
	错误
*/
//获取到所有代币的行情以及订单信息
func GetAllQuoteData() ([]QuoteData, error) {
	//log := log.Named("GetAllQuoteData")
	quoteDatas := make([]QuoteData, 0)

	//计算当天零的时间
	now := time.Now()
	baseTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for marketCode, tokens := range tokens.MarketTokenMap {
		for _, token := range tokens {
			if (marketCode == configs.INEOSStr && token.Symbol == configs.EOSStr) ||
				(marketCode == configs.INETHStr && token.Symbol == configs.ETHStr) ||
				(marketCode == configs.EOSETHStr && token.Symbol == configs.EOSStr) ||
				(marketCode == configs.INBDEStr && token.Symbol == configs.BDEStr){
				continue
			}

			//获取24小时内零点价，最新价，最高价，最低价和成交量
			basePrice, newestPrice, highestPrice, lowestPrice, volume, err := txrecords.GetDatasInToday(marketCode, token.Symbol, baseTime)
			if err != nil {
				log.Errorw("获取获取24小时内零点价，最新价，最高价，最低价和成交量失败", "error", err)
				return nil, err
			}

			//从DB获取对应Symbol的所有OrderApi
			orderApi, err := GetOrderApisBySymbol(token)
			if err != nil {
				log.Errorw("getOrderApisBySymbol", "error", err)
				return nil, err
			}

			quoteDatas = append(quoteDatas, QuoteData{
				MarketCode: marketCode,
				TokenName:  token.Symbol,
				Market: MarketApi{
					Price:        newestPrice,
					HighestPrice: highestPrice,
					LowestPrice:  lowestPrice,
					BasePrice:    basePrice,
					Volume:       volume,
					Address: ContractAddress{
						Address: token.TokenContract,
						Decimal: int8(token.Decimal),
					},
				},
				Order: orderApi,
			})
		}
	}

	//for i := 0; i < len(quoteDatas); i++ {
	//
	//	if len(quoteDatas[i].Order.Buy) < 30 {
	//		fakeBuy := AddFakeQuoteData(quoteDatas[i].TokenName, quoteDatas[i].MarketCode, "buy", len(quoteDatas[i].Order.Buy))
	//		//fmt.Println("len(quoteDatas[i].Order.Buy)----:", len(fakeBuy))
	//		quoteDatas[i].Order.Buy = append(quoteDatas[i].Order.Buy, fakeBuy...)
	//	}
	//	if len(quoteDatas[i].Order.Sell) < 30 {
	//		fakeSell := AddFakeQuoteData(quoteDatas[i].TokenName, quoteDatas[i].MarketCode, "sell", len(quoteDatas[i].Order.Sell))
	//		//fmt.Println("len(quoteDatas[i].Order.Sell)----:", len(fakeSell))
	//		quoteDatas[i].Order.Sell = append(quoteDatas[i].Order.Sell, fakeSell...)
	//	}
	//}

	return quoteDatas, nil
}

/**
Description:从DeleOrder表获取对应Symbol的30个买单，卖单，未完成订单

Parameter:
	token:token的结构体

Return:
	此token的30个买单，卖单，未完成订单
	错误
*/
//从DeleOrder表获取对应Symbol的30个买单，卖单，未完成订单
func GetOrderApisBySymbol(token tokens.Tokens) (OrderApi, error) {

	var orderApi OrderApi

	completeOrders := make([]CompleteOrder, 0)
	buyOrders := make([]ApiOrder, 0)
	sellOrders := make([]ApiOrder, 0)

	//找出交易记录中指定代币的最新的30个买单
	txRecords, err := txrecords.GetNewestTxRecords(token.MarketCode, token.Symbol, configs.BUYStr)
	if err != nil {
		log.Errorw("TransactionRecords Find", "error", err)
		return orderApi, err
	}

	for _, txRecord := range txRecords {
		completeOrders = append(completeOrders, CompleteOrder{
			Price:  txRecord.Price,
			Amount: txRecord.Count,
			Date:   txRecord.UpdateAt,
		})
	}

	//找出订单表中指定代币的未完成的30个价格的sell订单，按价格降序
	sellOrders, err = GetApiOrders(token.MarketCode, token.Symbol, token.OrderDecimal, configs.SELLStr)
	if err != nil {
		log.Errorw("DeleOrder table query sell DeleOrder", "error", err)
		return orderApi, err
	}
	//fmt.Println("len(sellOrders)----:",len(sellOrders))

	//找出订单表中指定代币的未完成的30个价格的buy订单，按价格升序
	buyOrders, err = GetApiOrders(token.MarketCode, token.Symbol, token.OrderDecimal, configs.BUYStr)
	//dbBuyOrders, err := datas.db.Query("select price,sum(sell_token_amount) as sellAmount,sell_token_symbol from DeleOrder where market_type=" + strconv.Itoa(token.MarketCode) + " and trade_type=0 and buy_token_symbol='" + token.Symbol + "," + strconv.Itoa(int(token.OrderDecimal)) + "' and ex_success=0 group by price order by price limit 0,30")
	if err != nil {
		log.Errorw("DeleOrder table query sell DeleOrder", "error", err)
		return orderApi, err
	}
	//fmt.Println("len(sellOrders)----:",len(buyOrders))

	orderApi.Buy = buyOrders
	orderApi.Sell = sellOrders
	orderApi.Complete = completeOrders

	return orderApi, nil
}

/**
Description:找出订单表中指定代币的未完成的30个价格的订单，根据买卖类型按价格排序

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	decimal:此代币的精度
	tradeType:交易类型  1:sell  0:buy

Return:
	30个ApiOrder数组
	错误
*/
//找出订单表中指定代币的未完成的30个价格的订单，根据买卖类型按价格排序
func GetApiOrders(marketCode int, symbol string, decimal uint, tradeType string) ([]ApiOrder, error) {

	apiOrders := make([]ApiOrder, 0)

	orders, err := deleorder.GetThirtyUnfinishedDeleOrder(marketCode, symbol, decimal, tradeType)
	if err != nil {
		log.Errorw("DeleOrder table query sell DeleOrder", "error", err)
		return apiOrders, err
	}

	if len(orders) != 0 {
		for _, order := range orders {
			sellAmount := string(order["sellAmount"])
			sellTokenSymbol := string(order["sell_token_symbol"])

			price, err := utils.AtofPrice(string(order["price"]))
			if err != nil {
				log.Errorw("getOrderApisBySymbol AtofParice", "error", err)
				return apiOrders, nil
			}
			_, sellCount := utils.AtoAmount(sellTokenSymbol, sellAmount)

			if tradeType == configs.BUYStr {
				sellCount = sellCount / price
			}
			//amount := sellCount / price
			apiOrders = append(apiOrders, ApiOrder{
				Price:  price,
				Amount: sellCount,
			})
		}
	}
	return apiOrders, nil
}
