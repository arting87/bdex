package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Tick struct {
	ID   int64 `json:"id"` // 消息ID
	Ts   int64 `json:"ts"` // 响应生成时间点
	Data []TickData
}

type TickData struct {
	Amount    float64 `json:"amount"` // 24小时成交量
	Ts        int64   `json:"ts"`
	ID        float64 `json:"id"` // 消息ID
	Price     float64 `json:"price"`
	Direction string  `json:"direction"`
}

type Trade struct {
	Status  string `json:"status"` // 请求状态
	Ts      int64  `json:"ts"`     // 响应生成时间点
	Tick    Tick   `json:"tick"`
	Ch      string `json:"ch"` // 数据所属的Channel, 格式: market.$symbol.depth.$type
	ErrCode string `json:"err-code"`
	ErrMsg  string `json:"err-msg"`
}

const (
	// SYMBOL          = "etheos"
	TradeUrl string = "http://api.huobi.pro/market/trade"
)

func GetTrade(symbol string) (float64, error) {

	fmt.Println("正在获取火币行情。。。")
	mapParams := make(map[string]string)
	mapParams["symbol"] = symbol
	var jsonMarketDetailReturn string
	var err error

	jsonMarketDetailReturn, err = HttpGetRequest(TradeUrl, mapParams)
	if err != nil {
		fmt.Println("获取火币行情第失败---", "error---:", err)
		return 0, err
	}

	marketDetailReturn := Trade{}
	err = json.Unmarshal([]byte(jsonMarketDetailReturn), &marketDetailReturn)
	if err != nil {
		fmt.Println("json解析火币行情失败", "error", err)
		return 0, err
	}

	if marketDetailReturn.Status != "ok" {
		fmt.Println("返回数据不为ok", "status", marketDetailReturn.Status)
		return 0, err
	}

	return marketDetailReturn.Tick.Data[0].Price, nil
}

func HttpGetRequest(strUrl string, mapParams map[string]string) (string, error) {
	httpClient := &http.Client{}

	var strRequestUrl string
	if nil == mapParams {
		strRequestUrl = strUrl
	} else {
		strParams := Map2UrlQuery(mapParams)
		strRequestUrl = strUrl + "?" + strParams
	}

	// 构建Request, 并且按官方要求添加Http Header
	request, err := http.NewRequest("GET", strRequestUrl, nil)
	if nil != err {
		fmt.Println("组装火币请求失败", "error", err)
		return "", err
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")

	// 发出请求
	response, err := httpClient.Do(request)
	if nil != err {
		fmt.Println("发送火币请求失败", "error", err)
		return "", err
	}
	defer response.Body.Close()
	// 解析响应内容
	body, err := ioutil.ReadAll(response.Body)
	if nil != err {
		fmt.Println("解析火币响应内容失败", "error", err)
		return "", err
	}
	return string(body), nil
}

func Map2UrlQuery(mapParams map[string]string) string {
	var strParams string
	for key, value := range mapParams {
		strParams += (key + "=" + value + "&")
	}

	if 0 < len(strParams) {
		strParams = string([]rune(strParams)[:len(strParams)-1])
	}
	return strParams
}
