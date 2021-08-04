package chainConn

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/ethereum/go-ethereum/common"
)

func AddEthToken(tokenName string, tokenAddress string, decimal int, marketCode int) error {
	if tokenName == "" || tokenAddress == "" || tokenAddress == Address_Nill || marketCode <= 0 || decimal <= 0 {
		return fmt.Errorf("参数为空或参数错误")
	}
	if marketCode == 4 {
		return fmt.Errorf("不支持添加的交易类型")
	}

	if len(tokenAddress) != 42 {
		return fmt.Errorf("以太坊地址错误")
	}

	tx, err := ExchangeClient.ExDataContract.AddTokens(ExchangeClient.Auth, tokenName, common.HexToAddress(tokenAddress), uint8(decimal), uint64(marketCode))
	if err != nil {
		logs.Error("AddEthToken error", err)
		return err
	}
	istrue, err := ExchangeClient.IsTransactionSuccess(tx)
	if err != nil || !istrue {
		return fmt.Errorf("链上交易失败")
	}
	return nil
}

func DeleteEthToken(tokenAddress string, marketCode int) error {
	if tokenAddress == "" || tokenAddress == Address_Nill || marketCode <= 0 {
		return fmt.Errorf("参数为空或参数错误")
	}
	if marketCode == 4 {
		return fmt.Errorf("不支持删除的交易类型")
	}
	tx, err := ExchangeClient.ExDataContract.DeleteToken(ExchangeClient.Auth, common.HexToAddress(tokenAddress), uint64(marketCode))
	if err != nil {
		logs.Error("AddEthToken error", err)
		return err
	}
	istrue, err := ExchangeClient.IsTransactionSuccess(tx)
	if err != nil || !istrue {
		return fmt.Errorf("链上交易失败")
	}
	return nil
}

func AddEosToken(tokenAddress string, symbol string, decimal uint64) error {
	if tokenAddress == "" || symbol == "" || decimal <= 0 {
		return fmt.Errorf("参数为空或参数错误")
	}
	addTokenAction := NewAddTokenAction(tokenAddress, symbol, decimal)

	txresp, err := BuildTx(addTokenAction)
	if err != nil {
		logs.Error("Send addToken to eos Chain", "ErrorInMem", err)
		return err
	}
	logs.Info("添加代币成功", "交易HASH", txresp.TransactionID)

	return nil
}

func DeleteEosToken(tokenId string) error {
	if tokenId == "" {
		return fmt.Errorf("参数为空")
	}
	deleteTokenAction := NewDeleteTokenAction(tokenId)

	txresp, err := BuildTx(deleteTokenAction)
	if err != nil {
		logs.Error("Send Deletoken to eos Chain", "ErrorInMem", err)
		return err
	}
	logs.Info("添加代币成功", "交易HASH", txresp.TransactionID)

	return nil
}
