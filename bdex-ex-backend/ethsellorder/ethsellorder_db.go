package ethsellorder

import (
	"time"
)

/**
Description:插入订单至EthSellOrder表

Parameter:
	order:EthSellOrder订单结构体

Return:
	错误
*/
//插入ETH的卖单到EthSellOrder数据库
func InsertEthSellOrder(ethSellOrder *EthSellOrder) error {
	_, err := db.Table("EthSellOrder").Insert(&ethSellOrder)
	if err != nil {
		log.Errorw("EthSellOrder table insert", "error", err)
		return err
	}

	return err
}

/**
Description:通过订单ID删除EthSellOrder订单

Parameter:
	orderId:EthSellOrder订单ID

Return:
	错误
*/
//通过订单ID删除EthSellOrder订单
func DeleteEthSellOrderById(orderId string) error {
	ccc := EthSellOrder{}

	_, err := db.Where("order_id = ?", orderId).Delete(&ccc)
	if err != nil {
		log.Errorw("delete  EthSellOrder error", err)
		return err
	}
	return nil
}

/**
Description:获取所有的EthSellOrder订单

Parameter:

Return:
	EthSellOrder数组
	错误
*/
//获取所有的EthSellOrder订单
func GetAllEthSellOrder() ([]EthSellOrder, error) {
	orders := make([]EthSellOrder, 0)
	err := db.Table("EthSellOrder").Desc("create_at").Find(&orders)
	if err != nil {
		log.Errorw("EthSellOrder Table find", "error", err)
		return nil, err
	}

	return orders, nil
}

type EthSellOrder struct {
	OrderId           string    `xorm:"'order_id' varchar(255) pk" json:"i"`
	Name              string    `xorm:"'name' varchar(255)"`
	NameChainCode     string    `xorm:"'name_chain_code' varchar(10)"`
	TradeType         string    `xorm:"'trade_type' varchar(8)" json:"a"` //1 is sell, 0 is buy
	MarketType        int       `xorm:"'market_type' int" `
	SellTokenSymbol   string    `xorm:"'sell_token_symbol' varchar(20)"`
	SellTokenAmount   string    `xorm:"'sell_token_amount' varchar(20)"`
	SellTokenContract string    `xorm:"'sell_token_contract' varchar(255)"`
	SellDecimals      uint      `xorm:"'sell_decimals' uint"`
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
}
