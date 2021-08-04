package txrecords

import (
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
)

type TransactionRecords struct {
	OrderId       string    `xorm:"'order_id' varchar(255) index" `
	Name          string    `xorm:"'name' varchar(255)"`
	NameChainCode string    `xorm:"'name_chain_code' varchar(10)"`
	TradeType     string    `xorm:"'trade_type' varchar(8)" ` //1 is sell, 0 is buy
	MarketType    int       `xorm:"'market_type' int index" `
	Symbol        string    `xorm:"'symbol' varchar(20)"`
	Count         float64   `xorm:"'count' double"`
	Price         float64   `xorm:"'price' double" `
	Fee           float64   `xorm:"'fee' double"`
	UpdateAt      time.Time `xorm:"'update_at' index"`
	MatchOrderId  string    `xorm:"'match_orderid' varchar(255) index"`
	TxHash        string    `xorm:"'tx_hash' varchar(255)"`
}

func InsertTxRecords(txRecord *TransactionRecords) error {
	_, err := db.Table("TransactionRecords").Insert(txRecord)
	if err != nil {
		log.Errorw("TransactionRecords table insert", "error", err)
		return err
	}
	return nil
}

/**
Description:查询某种代币在某时间之后的所有交易记录

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:在此时间之后的所有交易记录

Return:
	交易记录数组
	错误
*/
//查询某种代币在某时间之后的所有交易记录
func GetTxRecordsAfterBTime(marketCode, symbol string, baseTime time.Time) ([]TransactionRecords, error) {
	txRecords := make([]TransactionRecords, 0)

	err := db.Table("TransactionRecords").Where("market_type = ?", marketCode).And("trade_type = ?", configs.SELLStr).And("symbol = ?", symbol).And("update_at > ?", baseTime).Find(&txRecords)
	if err != nil {
		log.Errorw("TransactionRecords find", "error", err)
		return nil, err
	}

	return txRecords, nil
}

/**
Description:获取某种代币查询离某时间之前最近的一笔订单价格

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:离此时间之前最近的一笔订单价格

Return:
	价格的字节类型
	错误
*/
//查询离当天0点之前最近的一笔订单价格
func GetClosestTxRecordPriceBeforeBTime(marketCode, symbol string, baseTime time.Time) ([]byte, error) {
	price := make([]byte, 0)

	beforeBaseResult, err := db.Query("select price from TransactionRecords where market_type=" + marketCode + " and symbol='" + symbol + "' and update_at <= " + baseTime.Format("2006-01-02") + " order by update_at desc limit 0,1;")
	if err != nil {
		log.Errorw("查询离当天0点之前最近的一笔订单价格失败", "error", err)
		return price, err
	}

	if len(beforeBaseResult) != 0 {
		price = beforeBaseResult[0]["price"]
	}
	return price, nil
}

/**
Description:查询离某时间之后最近的一笔订单价格

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:离此时间之后最近的一笔订单价格

Return:
	价格的字节类型
	错误
*/
//查询离某时间之后最近的一笔订单价格
func GetClosestTxRecordPriceAfterBTime(marketCode, symbol string, baseTime time.Time) ([]byte, error) {
	price := make([]byte, 0)

	afterBaseResult, err := db.Query("select price from TransactionRecords where market_type=" + marketCode + " and symbol='" + symbol + "' and update_at >" + baseTime.Format("2006-01-02") + " order by update_at limit 0,1;")
	if err != nil {
		log.Errorw("查询离某时间之后最近的一笔订单价格", "error", err)
		return price, err
	}

	if len(afterBaseResult) != 0 {
		price = afterBaseResult[0]["price"]
	}
	return price, nil
}

/**
Description:找出交易记录中指定代币的最新的30个订单

Parameter:
	marketCode(Int):此代币属于哪一条链 1:eos 2:eth
	symbol(String):此代币的名称 ps:DDD
	tradeType:交易类型  1:sell  0:buy

Return:
	30个交易记录数组
	错误
*/
//找出交易记录中指定代币的最新的30个订单
func GetNewestTxRecords(marketCode int, symbol, tradeType string) ([]TransactionRecords, error) {
	txRecords := make([]TransactionRecords, 0)

	err := db.Table("TransactionRecords").Where("market_type = ?", marketCode).And("symbol = ?", symbol).And("trade_type = ?", tradeType).Desc("update_at").Limit(30, 0).Find(&txRecords)
	if err != nil {
		log.Errorw("TransactionRecords Find", "error", err)
		return txRecords, err
	}

	return txRecords, nil
}

/**
Description:查询指定market、指定symbol的卖单交易记录

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:离此时间之后订单交易记录

Return:
	交易记录数组
	错误
*/
//查询指定market、指定symbol的卖单交易记录
func GetSellTxRecordsBySymbol(marketCode, timeType int, symbol string, baseTime time.Time) (float64, []TransactionRecords, error) {
	closestStartPrice := float64(0)

	session := db.Table("TransactionRecords").Where("market_type = ?", marketCode).And("symbol = ?", symbol).And("trade_type = ?", "1")

	//从baseTime开始获取,K线图的update数据
	session = session.And("update_at >= ?", baseTime)

	//需要查询距离开始时间最近的一笔交易记录
	resultsSlice, err := db.Query("select price from TransactionRecords where update_at < '" + baseTime.Format("2006-01-02 15:04:05") + "' and trade_type = 1 and market_type = " + strconv.Itoa(marketCode) + " and symbol = '" + symbol + "' order by update_at desc limit 0,1;")
	if err != nil {
		log.Errorw("查询距离开始时间最近的一笔交易记录失败", "error", err)
		return 0, nil, err
	}

	if len(resultsSlice) != 0 {
		//update_at, err := time.Parse("2006-01-02 15:04:05", string(resultsSlice[0]["update_at"]))
		//if err != nil {
		//	log.Errorw("转化距离开始时间最近的交易记录时间失败", "error", err)
		//	return closestStartTxRecord, nil, err
		//}
		//closestStartTxRecord.UpdateAt = update_at
		f, err := strconv.ParseFloat(string(resultsSlice[0]["price"]), 64)
		if err != nil {
			log.Errorw("转化距离开始时间最近的交易价格失败", "error", err)
			return 0, nil, err
		}
		closestStartPrice = f
	}

	txRecords := make([]TransactionRecords, 0)

	err = session.Asc("update_at").Find(&txRecords)
	if err != nil {
		log.Errorw("TransactionRecords table find", "error", err)
		return 0, nil, err
	}

	return closestStartPrice, txRecords, nil

}
