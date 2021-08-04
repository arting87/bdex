package core

import (
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	ethContract "bdex.in/bdex/bdex-ex-backend/contract/ethereum"
	"bdex.in/bdex/bdex-ex-backend/core/ethaccountlist"
	"bdex.in/bdex/bdex-ex-backend/core/ethputlist"
	"github.com/GyhzzZ/eos-go"
	"github.com/GyhzzZ/eos-go/ecc"
	"github.com/spf13/viper"
)

type ExService struct {
	Engine *Engine                        //撮合引擎
	Epl    *ethputlist.EthPutList         //ETH缓冲下单队列
	EOS    *eosContract.EosClient         //EOS客户端
	ETH    *ethContract.EthClient         //ETH客户端
	EAL    *ethaccountlist.EthAccountList //以太坊账户队列
}

type ReqBody struct {
	Transaction *eos.Transaction `json:"transaction"`
	Signatures  []ecc.Signature  `json:"signatures"`
}

var (
	AN   = eos.AN
	ActN = eos.ActN
	PN   = eos.PN
)

// EOSStr service
func NewService() (*ExService, error) {
	cfg := viper.Sub("eos")
	creator := AN(cfg.GetString("creator.name"))
	signer := eos.NewKeyBag()
	err := signer.ImportPrivateKey(cfg.GetString("creator.privateKey"))
	if err != nil {
		log.Errorw("ImportPrivateKey", "error:", err)
		return nil, err
	}

	eos := eosContract.NewEosClient(viper.GetString("eos.url"))
	eos.SetSigner(creator, signer)

	//构建封装的ETH客户端
	eth, err := ethContract.NewEthClient(ethContract.ETH_CONN_ADDRESS)
	if err != nil {
		log.Errorf("new eth client error:", err)
		return nil, err
	}

	return &ExService{
		EOS:    eos,
		Engine: new(Engine),
		ETH:    eth,
	}, nil
}
