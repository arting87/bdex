package controller

import (
	"math"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/utils"
)

/*
与user get 请求有关的结构体
*/
type UserReqData struct {
	Account   string `json:"account"`
	ChainCode string `json:"chainCode"`
}

// 用户订单请求
type UserReq struct {
	Namespace string        `json:"namespace"`
	Action    string        `json:"action"`
	Data      []UserReqData `json:"data"`
}

type UserRespDatas struct {
	Account   string         `json:"account"`
	ChainCode string         `json:"chainCode"`
	Datas     []UserRespData `json:"datas"`
}

type UserRespData struct {
	MarketCode string      `json:"marketCode"`
	TokenName  string      `json:"tokenName"`
	Current    []OrderData `json:"currentOrder"`
	History    []OrderData `json:"historyOrder"`
}

type OrderData struct {
	OrderId    string    `json:"orderId"`
	Date       time.Time `json:"date"`
	Type       string    `json:"type"`
	OrderPrice float64   `json:"orderPrice"`
	OrderVol   float64   `json:"orderVol"`
	DealtVol   float64   `json:"dealtVol"`
	PendingVol float64   `json:"pendingVol"`
	Status     int       `json:"status"`
}

type UserResp struct {
	Namespace string         `json:"namespace"`
	Action    string         `json:"action"`
	UserData  []UserRespData `json:"data"`
}

/**
Description:获取用户订单数据

Parameter:
	req:请求的用户信息
		account:用户地址
		chainCode:用户属于哪条链 1:eos 2:eth

Return:
	用户的当前订单历史订单数组 每一个元素是关于一种代币
	错误
*/
//获取用户订单数据
func GetUserOrderRecords(req *UserReqData) ([]UserRespData, error) {
	log := log.Context("GetUserOrderRecords")

	userRespDatas := make([]UserRespData, 0)

	//获取某用户所有的当前订单和历史订单
	userRespDatasMap, err := GetPersonDeleOrderDates(req.Account, req.ChainCode)
	if err != nil {
		log.Errorw("获取某用户所有的当前订单和历史订单失败", "error", err)
		return nil, err
	}

	for _, tokenmap := range userRespDatasMap {
		userRespDatas = append(userRespDatas, *tokenmap)
	}

	return userRespDatas, nil
}

/**
Description:获取某用户指定代币的当前订单和历史订单

Parameter:
	account:用户的地址
	chainCode:用户属于那条链 1:eos 2:eth

Return:
	用户当前订单和历史订单的map key:token名称 value:当前订单和历史订单
	错误
*/
//获取某用户的所有当前订单和历史订单
func GetPersonDeleOrderDates(account string, chainCode string) (map[string]*UserRespData, error) {
	//获取某用户的所有订单
	orders, err := deleorder.GetDeleOrderByUser(account, chainCode)
	if err != nil {
		log.Errorw("DeleOrder table find", "error", err)
		return nil, err
	}

	userRespDatasMap := make(map[string]*UserRespData) //key:MatchSymbol

	for _, order := range orders {

		market_type := strconv.Itoa(order.MarketType)

		tradeType, err := strconv.ParseBool(order.TradeType)
		if err != nil {
			log.Errorw("ParseBool", "error", err)
			return nil, err
		}

		price, err := utils.AtofPrice(order.Price)
		if err != nil {
			log.Errorw("AtofPrice", "error", err)
			return nil, err
		}

		sellTokenSymbol, sellTokenPrecision := utils.ParseSymbol(order.SellTokenSymbol)
		buyTokenSymbol, buyTokenPrecision := utils.ParseSymbol(order.BuyTokenSymbol)

		sellTokenAmount, err := strconv.ParseFloat(order.SellTokenAmount, 64)
		if err != nil {
			log.Errorw("ParseFloat", "error", err)
			return nil, err
		}

		buyTokenAmount, err := strconv.ParseFloat(order.BuyTokenAmount, 64)
		if err != nil {
			log.Errorw("ParseFloat", "error", err)
			return nil, err
		}

		orderData := OrderData{
			OrderId:    order.OrderId,
			Date:       order.CreateAt,
			Type:       order.TradeType,
			OrderPrice: price,
		}

		//if sellTokenAmount < 0 {
		//	orderData.Status = CANCLE
		//	sellTokenAmount = -sellTokenAmount
		//} else if sellTokenAmount == 0 {
		//	orderData.Status = SUCCESS
		//} else {
		//	orderData.Status = UNSUCCESS
		//}

		orderData.Status = order.Status

		sellMoney := sellTokenAmount / math.Pow(10, float64(sellTokenPrecision))
		buyMoney := buyTokenAmount / math.Pow(10, float64(buyTokenPrecision))

		tokenName := ""
		if !tradeType { //买单 selltoken->EOS  buytoken->MV
			tokenName = buyTokenSymbol
			orderData.OrderVol = order.DeleVol
			orderData.DealtVol = buyMoney
			orderData.PendingVol = sellMoney / price
		} else { //卖单 selltoken->MV     buytoken->EOS
			tokenName = sellTokenSymbol
			orderData.OrderVol = order.DeleVol
			orderData.DealtVol = buyMoney / price
			orderData.PendingVol = sellMoney
		}

		if orderData.Status == configs.UNSUCCESS || orderData.Status == configs.PUTORDERSUCCESS {
			if _, has := userRespDatasMap[order.MatchSymbol]; has {
				userRespDatasMap[order.MatchSymbol].Current = append(userRespDatasMap[order.MatchSymbol].Current, orderData)
			} else {
				userRespData := newUserRespData(market_type, tokenName)
				userRespData.Current = append(userRespData.Current, orderData)
				userRespDatasMap[order.MatchSymbol] = userRespData
			}
		} else {
			if _, has := userRespDatasMap[order.MatchSymbol]; has {
				userRespDatasMap[order.MatchSymbol].History = append(userRespDatasMap[order.MatchSymbol].History, orderData)
			} else {
				userRespData := newUserRespData(market_type, tokenName)
				userRespData.History = append(userRespData.History, orderData)
				userRespDatasMap[order.MatchSymbol] = userRespData
			}
		}

	}

	return userRespDatasMap, nil
}

//初始化一个新的某代币当前订单和历史订单的结构体
func newUserRespData(market_type string, token string) *UserRespData {
	return &UserRespData{
		MarketCode: market_type,
		TokenName:  token,
		Current:    make([]OrderData, 0),
		History:    make([]OrderData, 0),
	}
}
