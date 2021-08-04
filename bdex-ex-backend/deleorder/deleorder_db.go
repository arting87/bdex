package deleorder

import (
	"errors"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
)

/*
存储到Mysql数据的结构体都在此文件中
*/
type DeleOrder struct {
	OrderId           string    `xorm:"'order_id' varchar(255) pk" json:"i"`
	Name              string    `xorm:"'name' varchar(255)"`
	NameChainCode     string    `xorm:"'name_chain_code' varchar(10)"`
	TradeType         string    `xorm:"'trade_type' varchar(8)" json:"a"` //1 is sell, 0 is buy
	MarketType        int       `xorm:"'market_type' int" `
	MatchSymbol       string    `xorm:"'match_symbol' varchar(8)"` //1AAA   2DDD
	DeleVol           float64   `xorm:"'dele_vol' Double"`         //除EOS或ETH以外的数量
	SellTokenSymbol   string    `xorm:"'sell_token_symbol' varchar(20)"`
	SellTokenAmount   string    `xorm:"'sell_token_amount' varchar(20)"`
	SellTokenContract string    `xorm:"'sell_token_contract' varchar(255)"`
	BuyTokenSymbol    string    `xorm:"'buy_token_symbol' varchar(20)" json:"bs"`
	BuyTokenAmount    string    `xorm:"'buy_token_amount' varchar(20)"`
	BuyTokenContract  string    `xorm:"'buy_token_contract' varchar(255)" json:"bc"`
	Price             string    `xorm:"'price' varchar(255)" json:"p"`
	Fee               string    `xorm:"'fee' varchar(20)"`
	WithdrawAble      bool      `xorm:"'withdraw_able' bool"`
	ExSuccess         bool      `xorm:"'ex_success' bool"`
	TargetAddress     string    `xorm:"'target_address' varchar(255)" json:"ta"`
	TargetChainCode   string    `xorm:"'target_chain_code' varchar(10)"`
	CreateAt          time.Time `xorm:"'create_at' index"`
	UpdateAt          time.Time `xorm:"'update_at'"`
	TxHash            string    `xorm:"'tx_hash' varchar(255)"`
	Status            int       `xorm:"'status' int"`
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
func GetThirtyUnfinishedDeleOrder(marketCode int, symbol string, decimal uint, tradeType string) ([]map[string][]byte, error) {

	//dbSellOrders, err := db.Query("select * from DeleOrder as a inner join(select price,market_type,buy_token_symbol,trade_type,ex_success from DeleOrder where market_type=", token.MarketCode, " and trade_type=", 1, " and buy_token_symbol=", token.Symbol, ",", token.Decimal, " and ex_success=0 group by price order by price desc limit 0,30) as b on a.price=b.price and a.market_type=b.market_type and a.trade_type=b.trade_type and a.buy_token_symbol=b.buy_token_symbol and a.ex_success=b.ex_success")
	if tradeType == configs.SELLStr {
		orders, err := db.Query("select price,sum(sell_token_amount) as sellAmount,sell_token_symbol from DeleOrder where market_type=" + strconv.Itoa(marketCode) + " and trade_type=1 and sell_token_symbol='" + symbol + "," + strconv.Itoa(int(decimal)) + "' and ex_success=0 group by price order by price desc limit 0,30")
		if err != nil {
			log.Errorw("DeleOrder table query sell DeleOrder", "error", err)
			return nil, err
		}
		return orders, nil
	} else {
		orders, err := db.Query("select price,sum(sell_token_amount) as sellAmount,sell_token_symbol from DeleOrder where market_type=" + strconv.Itoa(marketCode) + " and trade_type=0 and buy_token_symbol='" + symbol + "," + strconv.Itoa(int(decimal)) + "' and ex_success=0 group by price order by price limit 0,30")
		if err != nil {
			log.Errorw("DeleOrder table query sell DeleOrder", "error", err)
			return nil, err
		}
		return orders, nil
	}
}

/**
Description:查询某用户的所有订单

Parameter:
	account:用户地址
	chainCode:此代币的名称 1:eos 2:eth

Return:
	DeleOrder订单数组
	错误
*/
//查询某用户的所有订单
func GetDeleOrderByUser(account, chainCode string) ([]DeleOrder, error) {
	orders := make([]DeleOrder, 0)

	err := db.Table("DeleOrder").Where("name = ? and (name_chain_code = ? or name_chain_code = '')", account, chainCode).Find(&orders)
	if err != nil {
		log.Errorw("DeleOrder table find", "error", err)
		return nil, err
	}

	return orders, nil
}

/**
Description:插入订单至DeleOrder表

Parameter:
	order:DeleOrder订单结构体

Return:
	DeleOrder订单数组
	错误
*/
//插入订单至DeleOrder表
func InsertDeleOrder(order *DeleOrder) error {
	_, err := db.Table("DeleOrder").Insert(order)
	if err != nil {
		log.Errorw("db insert DeleOrder", "error", err)
		return err
	}
	return nil
}

/**
Description:修改DeleOrder表中的订单

Parameter:
	order:DeleOrder订单结构体

Return:
	错误
*/
//修改DeleOrder表中的订单
func UpdateDeleOrderById(order *DeleOrder) error {
	_, err := db.Table("DeleOrder").ID(order.OrderId).Cols("sell_token_amount", "buy_token_amount", "fee", "withdraw_able", "ex_success", "update_at", "status").Update(order)
	if err != nil {
		log.Errorw("DeleOrder Table update", "error", err)
		return err
	}
	return err
}

/**
Description:通过ID获取DeleOrder表中的订单

Parameter:
	orderId:DeleOrder订单的ID

Return:
	DeleOrder订单结构体
	错误
*/
//通过ID获取DeleOrder表中的订单
func GetOrderFromDbById(orderId string) (*DeleOrder, error) {
	order := DeleOrder{}
	_, err := db.Table("DeleOrder").Where("order_id = ?", orderId).Get(&order)
	if err != nil {
		log.Errorw("DeleOrder get by id", "error", err)
		return nil, err
	}

	if order.OrderId == "" {
		return nil, errors.New("this order does not exist")
	}

	return &order, nil
}

/**
Description:获取DeleOrder表中所有未完成的订单，按订单价格在前更新时间在后的顺序排序

Parameter:

Return:
	DeleOrder订单数组
	错误
*/
//获取所有未完成的订单，按订单价格在前更新时间在后的顺序排序
func GetAllUnfinishedDeleOrder() ([]DeleOrder, error) {
	orders := make([]DeleOrder, 0)

	err := db.Table("DeleOrder").OrderBy("price").OrderBy("update_at").Where("ex_success = ?", 0).And("status != ?", configs.Error).Find(&orders)
	if err != nil {
		log.Errorw("DeleOrder table orderby find", "error", err)
		return nil, err
	}

	return orders, nil
}

/**
Description:获取DeleOrder表中最新创建的一笔订单

Parameter:

Return:
	DeleOrder订单结构体
	错误
*/
//获取DeleOrder表中最新创建的一笔订单
func GetNewestDeleOrder() (DeleOrder, error) {
	newestOrder := DeleOrder{}
	_, err := db.Table("DeleOrder").Desc("create_at").Get(&newestOrder)
	if err != nil {
		log.Errorw("init ID", "error", err)
		return newestOrder, err
	}
	return newestOrder, nil
}

/**
Description:通过订单ID获取DeleOrder表中订单

Parameter:
	id:订单ID

Return:
	DeleOrder订单结构体
	错误
*/
func GetEOSOrderFromDbById(id string) (DeleOrder, error) {
	eosOrder := DeleOrder{
		OrderId: id,
	}
	_, err := db.Table("DeleOrder").Get(&eosOrder)
	if err != nil {
		log.Errorw("GetEOSOrderFromDbById", "error", err)
		return eosOrder, err
	}

	return eosOrder, nil
}

/**
Description:获取交易目标地址

Parameter:
    id ：订单id

Return:
    目标地址
	错误
*/
//获取交易目标地址
func GetTargetAddressFromDb(id uint64) (string, error) {
	orderid := strconv.FormatUint(id, 10)

	a := DeleOrder{}
	var addressTo string

	has, err := db.Table(&a).Where("order_id = ?", orderid).Cols("target_address").Get(&addressTo)
	if err != nil {
		log.Errorf("get Eth target_address error:", err)
		return "", nil
	}
	if has == false {
		log.Errorf("get Eth target_address not exist:", err)
		return "", nil
	}
	return addressTo, nil
}
