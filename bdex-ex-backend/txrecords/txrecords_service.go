package txrecords

import (
	"strconv"
	"time"
)

/**
Description:查询今日零点之后的基本价格，最新价格，最高价格，最低价格，成交量

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:离此时间之后的交易信息

Return:
	零点价
	最新价
	最高价
	最低价
	成交量
	错误
*/
//查询今日零点之后的基本价格，最新价格，最高价格，最低价格，成交量
func GetDatasInToday(marketCode, symbol string, baseTime time.Time) (basePrice, newestPrice, highestPrice, lowestPrice, volume float64, err error) {
	compareTime := time.Now()

	//从DB获取对应Symbol的basePrice
	basePrice, err = getBasePriceBySymbol(marketCode, symbol, baseTime)
	if err != nil {
		log.Errorw("getBasePriceBySymbol", "error", err)
		return
	}

	//查询该marketCode的零点之后所有的交易记录
	txRecords, err := GetTxRecordsAfterBTime(marketCode, symbol, baseTime)
	if err != nil {
		log.Errorw("TransactionRecords find", "error", err)
		return
	}

	//fmt.Println("len(txRecords)-----------------------:", len(txRecords))
	if len(txRecords) == 0 {
		newestPrice = basePrice
		highestPrice = basePrice
		lowestPrice = basePrice
		volume = 0
	} else {
		for k, txRec := range txRecords {
			if k == 0 {
				newestPrice = txRec.Price
				highestPrice = txRec.Price
				lowestPrice = txRec.Price
				compareTime = txRec.UpdateAt
				volume = txRec.Count
				continue
			}
			if compareTime.Before(txRec.UpdateAt) {
				newestPrice = txRec.Price
			}
			if highestPrice < txRec.Price {
				highestPrice = txRec.Price
			}
			if lowestPrice > txRec.Price {
				lowestPrice = txRec.Price
			}
			volume += txRec.Count
		}
	}

	return
}

/**
Description:查询今日零点之后的基本价格

Parameter:
	marketCode:此代币属于哪一条链 1:eos 2:eth
	symbol:此代币的名称 ps:DDD
	baseTime:离此时间之后的交易信息

Return:
	零点价
	错误
*/
//从TxRecord表获取对应Symbol的basePrice
func getBasePriceBySymbol(marketCode, symbol string, baseTime time.Time) (float64, error) {

	var basePrice float64

	//查询离当天0点之前最近的一笔订单价格
	beforBasePrice, err := GetClosestTxRecordPriceBeforeBTime(marketCode, symbol, baseTime)
	if err != nil {
		log.Errorw("db.Query", "error", err)
		return 0, err
	}

	if len(beforBasePrice) == 0 {
		//查询离当天0点之后最近的一笔订单价格
		afterBasePrice, err := GetClosestTxRecordPriceAfterBTime(marketCode, symbol, baseTime)
		//afterBaseResult, err := datas.db.Query("select min(update_at) as a,price from TransactionRecords where market_type=" + marketCode + " and symbol='" + symbol + "' and update_at >" + baseTime.Format("2006-01-02"))
		if err != nil {
			log.Errorw("db.Query", "error", err)
			return 0, err
		}

		if len(afterBasePrice) == 0 {
			//返回basePrice为0
			basePrice = 0
		} else {
			//设置basePrice为当天0点之后最近的一笔订单价格
			//fmt.Println("afterBaseResult[0]['price']", afterBaseResult[0]["price"])
			basePrice, err = strconv.ParseFloat(string(afterBasePrice), 64)
			if err != nil {
				log.Errorw("ParseFloat price", "error", err)
				return 0, err
			}
		}
	} else {
		//设置basePrice为当天0点之前最近的一笔订单价格
		basePrice, err = strconv.ParseFloat(string(beforBasePrice), 64)
		if err != nil {
			log.Errorw("ParseFloat price", "error", err)
			return 0, err
		}
	}
	return basePrice, nil
}
