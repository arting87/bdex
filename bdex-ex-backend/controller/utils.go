package controller

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// temporary function to fake quote data
func AddFakeQuoteData(tokenName, marketCode, tradeType string, lens int) []ApiOrder {
	fakeApiOrder := make([]ApiOrder, 0)

	if tradeType == "buy" {
		price := switchFakeMarket(marketCode, tokenName)
		fakeApiOrder = makeFakeQuoteData(price, lens, true)
	} else {
		price := switchFakeMarket(marketCode, tokenName)
		fakeApiOrder = makeFakeQuoteData(price, lens, false)
	}
	return fakeApiOrder
}

func makeFakeQuoteData(price float64, lens int, tradeType bool) []ApiOrder {
	fakeApiOrder := make([]ApiOrder, 0)
	fakeData := ApiOrder{}

	if tradeType {

		for i := 0; i < 30-lens; i++ {
			fakeData.Price = price
			fakeData.Amount = makeRandAmount()
			fakeApiOrder = append(fakeApiOrder, fakeData)
			price = price * 99 / 100
		}
	} else {

		for i := 0; i < 30-lens; i++ {
			fakeData.Price = price - 8
			fakeData.Amount = makeRandAmount()
			fakeApiOrder = append(fakeApiOrder, fakeData)
			price = price * 100 / 99
		}
	}

	return fakeApiOrder
}

func makeRandAmount() float64 {
	r := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	a, err := strconv.ParseFloat(r, 64)
	if err != nil {
		log.Error("line 173: ", "parse float64 error", err)
	}
	return a / 10000
}

func switchFakeMarket(marketCode, tokenName string) float64 {
	switch marketCode {
	case "1":
		symbol := strings.ToLower(tokenName) + "eos"
		f, err := GetTrade(symbol)
		if err != nil {
			log.Error("line 151: ", "get huobi quote error", err)
		}
		return f
	case "2":
		symbol := strings.ToLower(tokenName) + "eth"
		f, err := GetTrade(symbol)
		if err != nil {
			log.Error("line 151: ", "get huobi quote error", err)
		}
		return f
	case "4":
		symbol := "eos" + strings.ToLower(tokenName)
		f, err := GetTrade(symbol)
		if err != nil {
			log.Error("line 151: ", "get huobi quote error", err)
		}
		return 1 / f
	default:
		return 0
	}
}
