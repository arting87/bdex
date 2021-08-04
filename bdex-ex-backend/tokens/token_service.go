package tokens

import (
	"strconv"
	"strings"

	"bdex.in/bdex/bdex-ex-backend/configs"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

var MarketTokenMap = make(map[string][]Tokens) //市场-代币

func LoadTokenFromDB() (string, error) {
	tokens, err := GetAllToken()
	if err != nil {
		log.Errorw("Token Table finde", "error", err)
		return "", err
	}

	bdeAddress := ""

	for _, dbtoken := range tokens {
		if MarketTokenMap[strconv.Itoa(dbtoken.MarketCode)] == nil {
			tokenSlice := make([]Tokens, 0)
			tokenSlice = append(tokenSlice, dbtoken)
			MarketTokenMap[strconv.Itoa(dbtoken.MarketCode)] = tokenSlice
		} else {
			memHas := false
			for _, token := range MarketTokenMap[strconv.Itoa(dbtoken.MarketCode)] {
				if token.TokenContract == dbtoken.TokenContract && token.Symbol == dbtoken.Symbol {
					memHas = true
				}
			}
			if !memHas {
				MarketTokenMap[strconv.Itoa(dbtoken.MarketCode)] = append(MarketTokenMap[strconv.Itoa(dbtoken.MarketCode)], dbtoken)
			}
		}

		if dbtoken.Symbol == configs.BDEStr && dbtoken.MarketCode == configs.INBDE {
			bdeAddress = dbtoken.TokenContract
		}
	}

	//加载到内存中的Token
	log.Infow("加载到内存中的Token...")
	for market, tokens := range MarketTokenMap {
		log.Infow("行情为" + market + "的Tokens...")
		for _, token := range tokens {
			log.Infow(strconv.Itoa(token.MarketCode) + "_" + token.Symbol + "_" + token.TokenContract)
		}
	}

	return bdeAddress, nil
}

func QueryTokenExistFromMem(tokenContract, symbol string, tokenMarketCode int) error {
	tokenMc := strconv.Itoa(tokenMarketCode)
	for _, token := range MarketTokenMap[tokenMc] {
		//忽略大小写比对
		if strings.EqualFold(token.TokenContract, tokenContract) && token.Symbol == symbol {
			return nil
		}
	}
	err := errors.New(tokenMc + "_" + symbol + "_" + tokenContract + "此token不存在")
	return err
}

/**
Description:获取代币的名称

Parameter:
	addr:发布token的地址

Return:
	代币的名称
*/
//获取代币的名称
func GetEthTokenSymbleFromMem(marketCode int, addr common.Address) (string, error) {
	mc := strconv.Itoa(marketCode)
	//TODO：添加新的交易区修改
	if addr.String() == configs.AddressNil && marketCode == configs.EOSETH {
		return "EOS" + ",4", nil
	} else {
		for _, token := range MarketTokenMap[mc] {
			if strings.EqualFold(token.TokenContract, addr.String()) {
				return token.Symbol + ",8", nil
			}
		}
	}
	err := errors.New(mc + "_" + addr.String() + "没有代币")
	return "", err
}

/**
Description:根据合约地址获取合约小数点位数

Parameter:
    addr：合约地址

Return:
    交易哈希
	错误
*/
//获取代币的精度
func GetEthDecimalsFromMem(addr string) (uint, error) {
	ethMarketCode := strconv.Itoa(configs.INETH)
	for market, tokens := range MarketTokenMap {
		if market == ethMarketCode {
			for _, token := range tokens {
				if strings.EqualFold(token.TokenContract, addr) {
					return uint(token.Decimal), nil
				}
			}
		}
	}
	log.Errorw(ethMarketCode + "_" + addr + "此token不存在")
	return 0, errors.New(ethMarketCode + "_" + addr + "此token不存在")
}

func GetEthDecimalsByAddressAndMarketCode(addr string,marketCode int) (uint, error) {
	ethMarketCode := strconv.Itoa(marketCode)
	for market, tokens := range MarketTokenMap {
		if market == ethMarketCode {
			for _, token := range tokens {
				if strings.EqualFold(token.TokenContract, addr) {
					return uint(token.Decimal), nil
				}
			}
		}
	}
	err := errors.New(ethMarketCode + "_" + addr + "此token不存在")
	log.Errorw(ethMarketCode + "_" + addr + "此token不存在")
	return 0, err
}
