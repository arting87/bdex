package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"math"
	"math/big"
	"runtime"
	"strconv"
	"strings"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/log"

	math2 "github.com/ethereum/go-ethereum/common/math"
)

//字符串转float64的price
func AtofPrice(price string) (float64, error) {
	i, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return 0, err
	}
	var fPrice float64
	fPrice = i / configs.PRICEUTILS
	return fPrice, nil
}

//处理代币精度和数额
func AtoAmount(symbol, amount string) (string, float64) {
	s, precision := ParseSymbol(symbol)

	a, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		log.Errorw("convert amount error", "error", err)
		return "", 0
	}
	var money float64
	if precision == 0 {
		log.Errorw("convert amount error", "error", "precision can not be zero")
		return "", 0
	}
	money = a / math.Pow(10, float64(precision))
	return s, money
}

func ParseSymbol(symbol string) (string, int) {
	comma := strings.Index(symbol, ",")
	s := symbol[:comma]
	precision, err := strconv.Atoi(symbol[comma+1:])
	if err != nil {
		log.Errorw("convert decimal error", "error", err)
		return "", 0
	}
	return s, precision
}

func ComputeBigintToStringOfSell(sell *big.Int, decimal int64) string {
	a := new(big.Int)
	a.Mul(sell, math2.BigPow(10, 8))
	a.Div(a, math2.BigPow(10, decimal))

	return strconv.FormatInt(a.Int64(), 10)
}

func ComputeBigIntToStringOfBuy(sell *big.Int) string {
	a := new(big.Int)
	a.Mul(sell, math2.BigPow(10, 8))
	a.Div(a, math2.BigPow(10, 18))

	return strconv.FormatInt(a.Int64(), 10)
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//哈希处理
func SigDigest(payload []byte) []byte {
	h := sha256.New()
	_, _ = h.Write(payload)
	return h.Sum(nil)
}

//读取JSON文件
func LoadJsonFile(filename string, v interface{}) error {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Errorw("读取json文件失败", "error", err)
		return err
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Errorw("json.Unmarshal失败", "error", err)
		return err
	}
	return nil
}