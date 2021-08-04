package chainConn

import (
	"github.com/astaxie/beego/logs"
	"strconv"

	"github.com/GyhzzZ/eos-go"
	"github.com/GyhzzZ/eos-go/ecc"
)

type Transaction struct {
	Transaction *eos.Transaction `json:"transaction"`
	Signatures  []ecc.Signature  `json:"signatures"`
}

const ContractCode = "tbcexchange1"

var AN = eos.AN
var PN = eos.PN
var ActN = eos.ActN

func init() {
	eos.RegisterAction(AN(ContractCode), ActN("addtoken"), AddToken{})
	eos.RegisterAction(AN(ContractCode), ActN("deletoken"), DeleteToken{})
}

type AddToken struct {
	TokenContract string `json:"token_contract"`
	Symbol        string `json:"symbol"`
	Decimal       string `json:"decimal"`
}

type DeleteToken struct {
	Id string `json:"id"`
}

func NewAddTokenAction(tokenAddress string, symbol string, decimal uint64) *eos.Action {
	return &eos.Action{
		Account: AN(ContractCode),
		Name:    ActN("addtoken"),
		Authorization: []eos.PermissionLevel{
			{Actor: ContractCode, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(AddToken{
			TokenContract: tokenAddress,
			Symbol:        symbol,
			Decimal:       strconv.FormatUint(decimal, 10),
		}),
	}
}

func NewDeleteTokenAction(tokenId string) *eos.Action {
	return &eos.Action{
		Account: AN(ContractCode),
		Name:    ActN("deletetoken"),
		Authorization: []eos.PermissionLevel{
			{Actor: ContractCode, Permission: PN("active")},
		},
		ActionData: eos.NewActionData(DeleteToken{
			Id: tokenId,
		}),
	}
}

/**
Description:构建一笔交易发送至EOS链上

Parameter:
	a:构建好的eos的action
	chainaddr:第几条链的标记

Return:
	eos上链的返回值
	后端内部错误
	链上返回错误
*/
func BuildTx(a *eos.Action) (*eos.PushTransactionFullResp, error) {
	action := make([]*eos.Action, 1)
	action[0] = a

	opt := &eos.TxOptions{}
	api := eos.New(EosIp)
	err := opt.FillFromChain(api)
	if err != nil {
		return nil, err
	}

	tx := eos.NewTransaction(action, opt)
	sigtx := eos.NewSignedTransaction(tx)
	key := eos.NewKeyBag()
	err = key.ImportPrivateKey(EosOwnerPrivateKey)
	if err != nil {
		return nil, err
	}
	pub, _ := ecc.NewPublicKey(EosOwnerPublicKey)
	signTx, err := key.Sign(sigtx, opt.ChainID, pub)
	if err != nil {
		return nil, err
	}

	signedTx := &eos.SignedTransaction{
		Transaction: tx,
		Signatures:  signTx.Signatures,
	}
	//发送交易到链上
	packedTx, err := signedTx.Pack(eos.CompressionNone)
	if err != nil {
		logs.Error("Pack transaction", "error", err)
		return nil, err
	}

	r, err := api.PushTransaction(packedTx)
	if err != nil {
		logs.Error("Push transaction", "error", err)
		return nil, err
	}
	return r, nil
}
