package tokens

import (
	"bdex.in/bdex/bdex-ex-backend/configs"

	"github.com/ethereum/go-ethereum/common"
)

/**
Description:查找Token在数据库中是否存在

Parameter:
	tokencontract:发布token的合约地址
	symbol:token的名称
	marketCode:token属于哪一条链 1:eos 2:eth

Return:
	是否存在  true:存在  false:不存在
	错误
*/
//查找Token在数据库中是否存在
func QueryTokenExistFromDB(tokencontract, symbol string, marketCode int) (bool, error) {
	has := true
	var tokens []Tokens

	session := db.Where(nil)
	session = session.And("symbol = ?", symbol)
	session = session.And("token_contract = ?", tokencontract)
	if marketCode != 0 {
		session = session.And("market_code = ?", marketCode)
	}

	err := session.Find(&tokens)
	if err != nil {
		log.Errorw("query fail", "error", err)
		return has, err
	}

	if len(tokens) == 0 {
		has = false
	}
	return has, nil
}

/**
Description:插入token至Tokens表

Parameter:
	tokens:将要插入Tokens表的Token数组

Return:
	错误
*/
//插入Token
func InsertTokens(tokens *[]Tokens) error {
	_, err := db.Insert(tokens)
	if err != nil {
		log.Errorw("tokens insert", "error", err)
	}
	return err
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
func GetEthDecimals(addr string) (uint, error) {
	a := Tokens{}
	var decimal int
	has, err := db.Table(&a).Where("token_contract = ?", addr).Cols("decimal").Get(&decimal)
	if err != nil {
		log.Errorf("getEthDecimal error:", err)
		return 0, nil
	}
	if has == false {
		log.Errorf("getEthDecimal symbol not exist:", err)
		return 0, nil
	}
	return uint(decimal), nil
}

/**
Description:获取代币的名称

Parameter:
	addr:发布token的地址

Return:
	代币的名称
*/
//获取代币的名称
func GetBuyTokenSymble(addr common.Address) string {
	if addr.String() == configs.AddressNil {
		return "EOS" + ",4"
	} else {
		var symble string
		has, err := db.Table(&Tokens{}).Where("token_contract = ?", addr.String()).Cols("symbol").Get(&symble)
		if err != nil {
			log.Errorf("getEthSymble error:", err)
			return ""
		}
		if has == false {
			log.Errorf("getEthSymble not exist:", err)
			return ""
		}
		return symble + ",8"
	}
}

/**
Description:获取所有的代币

Parameter:

Return:
	Tokens数组
	错误
*/
//获取所有的代币
func GetAllToken() ([]Tokens, error) {
	tokens := make([]Tokens, 0)

	err := db.Table("Tokens").Where("iscancel = ?", 0).Find(&tokens)
	if err != nil {
		log.Errorw("Token Table finde", "error", err)
		return nil, err
	}
	return tokens, nil
}
