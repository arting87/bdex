package controller

import (
	"fmt"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/txrecords"
)

/*
与api k 请求有关的结构体
*/
//K线图数据请求
type KLineReqData struct {
	TimeType   int    `json:"timeType"`
	Symbol     string `json:"symbol"`
	MarketCode int    `json:"marketCode"`
}

type KLineRespData struct {
	GapStart   int64   `json:"gapStart"`   //当前时间段开始时间戳
	GapEnd     int64   `json:"gapEnd"`     //当前时间段结束时间戳
	LowerPrice float64 `json:"lowerPrice"` //当前时间段最低价
	HighPrice  float64 `json:"highPrice"`  //当前时间段最高价
	OpenPrice  float64 `json:"openPrice"`  //当前时间段开盘价
	ClosePrice float64 `json:"closePrice"` //当前时间段收盘价
	QuoteVol   float64 `json:"quoteVol"`   //当前时间段交易总量
}

/**
Description:插入订单至DeleOrder表

Parameter:
	req:指定代币的信息
		timeType:K线时间间隔类型，以分钟为单位
		symbol:代币名称
		marketCode:代币属于哪条链 1:eos 2:eth
	baseTime:查找K线数据的开始时间
	lastClosePrice:上次查找的收盘价

Return:
	K线数据的数组，每一个元素为一个时间间隔的数据
	此次查找的结束时间
	错误
*/
//获取指定代币的K线图信息
func GetKLineData(req KLineReqData, baseTime time.Time, lastClosePrice float64) ([]KLineRespData, time.Time, error) {
	log := log.Context("GetKLineData")

	resp := make([]KLineRespData, 0)

	//判断是取1000个时间片段的时间还是从baseTime开始取
	if !time.Now().After(baseTime) {
		//从baseTime开始获取,K线图的update数据
		baseTime = time.Now().Add(time.Duration(-configs.KLineGapCount*req.TimeType) * time.Minute)
	}
	//查询指定market、指定symbol的1000条卖单交易记录,和距离开始时间最近的交易的价格
	closestStartPrice, txRecords, err := txrecords.GetSellTxRecordsBySymbol(req.MarketCode, req.TimeType, req.Symbol, baseTime)
	nextBaseTime := time.Now()
	if err != nil {
		log.Errorw("TransactionRecords table find", "error", err)
		return nil, nextBaseTime, err
	}

	//如果此次未查询出数据
	if len(txRecords) == 0 {
		//如果在此时间之前也没有价格，则说明此代币没有交易记录
		if closestStartPrice == 0 {
			//返回空的K线数据
			if lastClosePrice < 0 {
				return resp, nextBaseTime, nil
			}
		} else {
			//如果之前有价格，则使用上一次的价格
			if lastClosePrice > 0 {
				kLineRespData := KLineRespData{
					GapStart:   baseTime.UnixNano() / 1e6,
					GapEnd:     baseTime.Add(time.Duration(req.TimeType) * time.Minute).UnixNano() / 1e6,
					LowerPrice: lastClosePrice,
					HighPrice:  lastClosePrice,
					OpenPrice:  lastClosePrice,
					ClosePrice: lastClosePrice,
					QuoteVol:   0,
				}
				nextBaseTime = baseTime.Add(time.Duration(req.TimeType) * time.Minute)
				resp = append(resp, kLineRespData)
				return resp, nextBaseTime, nil
			}
			//如果在此时间之前有价格，则使用此价格贯穿所有的时间片段，使用下面的代码
			kLines := CreatKlineWithoutRecords(baseTime, closestStartPrice, req.TimeType)
			return kLines, baseTime.Add(time.Duration(configs.KLineGapCount*req.TimeType) * time.Minute), nil
		}
	}

	//计算第一个时间间隔的始末
	gapTimeStart := baseTime
	//if lastClosePrice < 0 {
	//	gapTimeStart = txRecords[0].UpdateAt
	//}

	gapTimeEnd := gapTimeStart.Add(time.Duration(req.TimeType) * time.Minute)

	kLineRespData := KLineRespData{
		LowerPrice: closestStartPrice,
		HighPrice:  closestStartPrice,
		OpenPrice:  closestStartPrice,
	}
	fmt.Println(len(txRecords))
	//a := 0
	for key, txRecord := range txRecords {
		txTime := txRecord.UpdateAt
		currencyPrice := txRecord.Price
		//a++
		//fmt.Println(a)
		if txTime.Before(gapTimeEnd) || txTime == gapTimeEnd {
			//放入当前时间段内
			if currencyPrice < kLineRespData.LowerPrice {
				kLineRespData.LowerPrice = currencyPrice
			}
			if currencyPrice > kLineRespData.HighPrice {
				kLineRespData.HighPrice = currencyPrice
			}
			kLineRespData.QuoteVol = kLineRespData.QuoteVol + txRecord.Count
			kLineRespData.ClosePrice = currencyPrice
		} else { //如果这笔记录超出当前时间段
			//添加上一个时间段到数组
			kLineRespData.GapStart = gapTimeStart.UnixNano() / 1e6
			kLineRespData.GapEnd = gapTimeEnd.UnixNano() / 1e6
			resp = append(resp, kLineRespData)

			//判断超出几个时间段
			gapCount := txTime.Sub(gapTimeEnd).Minutes() / float64(req.TimeType)
			//fmt.Println("当前超过的时间段数量", gapCount)
			//为每一个时间段赋值
			for i := gapCount; i > 0; i-- {
				if i > 1 {
					//为当前时间段赋同样的值,赋值为上一个时间段的收盘价
					kLineRespData.LowerPrice = kLineRespData.ClosePrice
					kLineRespData.HighPrice = kLineRespData.ClosePrice
					kLineRespData.OpenPrice = kLineRespData.ClosePrice
					//kLineRespData.ClosePrice = currencyPrice
					kLineRespData.QuoteVol = 0
					//添加当前空时间段放入数组
					kLineRespData.GapStart = gapTimeStart.UnixNano() / 1e6
					kLineRespData.GapEnd = gapTimeEnd.UnixNano() / 1e6
					resp = append(resp, kLineRespData)
				} else {
					//此为最后一个时间段，需要用于判断之后一组是否属于这一组
					kLineRespData.LowerPrice = currencyPrice
					kLineRespData.HighPrice = currencyPrice
					kLineRespData.OpenPrice = currencyPrice
					kLineRespData.QuoteVol = txRecord.Count
					kLineRespData.ClosePrice = currencyPrice
				}
				//重置时间段
				gapTimeStart = gapTimeEnd
				gapTimeEnd = gapTimeStart.Add(time.Duration(req.TimeType) * time.Minute)
			}
		}
		//若当前为最后一笔订单,则添加
		if key == len(txRecords)-1 {
			kLineRespData.GapStart = gapTimeStart.UnixNano() / 1e6
			kLineRespData.GapEnd = gapTimeEnd.UnixNano() / 1e6
			resp = append(resp, kLineRespData)

			//判断当前时间距离下次返回k线图时间距离几个时间段
			//判断超出几个时间段
			gapCount := nextBaseTime.Sub(gapTimeEnd).Minutes() / float64(req.TimeType)
			for i := gapCount; i >= 0; i-- {

				if i > 1 {
					gapTimeStart = gapTimeEnd
					gapTimeEnd = gapTimeStart.Add(time.Duration(req.TimeType) * time.Minute)
					//为当前时间段赋同样的值,赋值为上一个时间段的收盘价
					kLineRespData.LowerPrice = kLineRespData.ClosePrice
					kLineRespData.HighPrice = kLineRespData.ClosePrice
					kLineRespData.OpenPrice = kLineRespData.ClosePrice
					kLineRespData.QuoteVol = 0
					//添加当前空时间段放入数组
					kLineRespData.GapStart = gapTimeStart.UnixNano() / 1e6
					kLineRespData.GapEnd = gapTimeEnd.UnixNano() / 1e6
					resp = append(resp, kLineRespData)

				}
			}
		}
	}
	nextBaseTime = gapTimeEnd
	return resp, nextBaseTime, nil
}

func CreatKlineWithoutRecords(baseTime time.Time, price float64, timeType int) []KLineRespData {
	resp := make([]KLineRespData, 0)
	//endTime := baseTime.Add(time.Duration(1000*timeType) * time.Minute)
	gapTimeStart := baseTime
	//创建1000个时间片段
	for i := 0; i < configs.KLineGapCount; i++ {
		gapTimeEnd := gapTimeStart.Add(time.Duration(timeType) * time.Minute)
		resp = append(resp, KLineRespData{
			GapStart:   gapTimeStart.UnixNano() / 1e6,
			GapEnd:     gapTimeEnd.UnixNano() / 1e6,
			LowerPrice: price,
			HighPrice:  price,
			OpenPrice:  price,
			ClosePrice: price,
			QuoteVol:   0,
		})
		gapTimeStart = gapTimeEnd.Add(time.Duration(timeType) * time.Minute)
	}
	return resp
}
