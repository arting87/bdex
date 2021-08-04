package sync

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"bdex.in/bdex/bdex-ex-backend/bdexerrors"
	"bdex.in/bdex/bdex-ex-backend/configs"
	eosContract "bdex.in/bdex/bdex-ex-backend/contract/eos"
	"bdex.in/bdex/bdex-ex-backend/contract/ethereum"
	"bdex.in/bdex/bdex-ex-backend/core"
	tokensDB "bdex.in/bdex/bdex-ex-backend/tokens"

	"github.com/GyhzzZ/eos-go"
	"github.com/spf13/viper"
)

type TokenTable struct {
	TokenID       uint64 `json:"id"`
	Symbol        string `json:"symbol"`
	TokenContract string `json:"tcontract"`
	Decimal       string `json:"decimal"`
	IsCancle      bool   `json:"is_cancle"`
}

type EthTokenTable struct {
	TokenID       uint64 `json:"id"`
	Symbol        string `json:"symbol"`
	TokenContract string `json:"tcontract"`
	Decimal       uint8  `json:"decilmal"`
	IsCancel      bool   `json:"is_cancle"`
	MarketCode    uint64 `json:"market_code"`

	/*ID            int    `xorm:"'id' int pk autoincr unique"`
	TokenID       uint64 `xorm:"'token_id' BigInt index"`
	Symbol        string `xorm:"'symbol' varchar(255)"`
	TokenContract string `xorm:"'token_contract' varchar(42)"`
	MarketCode    int    `xorm:"'market_code' int"`
	Decimal       uint   `xorm:"'decimal' int"`
	OrderDecimal  uint   `xorm:"'order_decimal' int"`
	IsCancel      bool   `xorm:"'iscancel' bool"`*/
}

type Syncer struct {
	eosClient *eosContract.EosClient
	ethClient *ethereum.EthClient

	wg sync.WaitGroup

	ctx    context.Context
	cancel context.CancelFunc

	limitContract  uint32
	limitDB        int
	batchInterval  int
	updateInterval int
	tokenID        uint64
}

func SyncServe(ctx context.Context) {
	setupSyncLog()

	ctx, cancel := context.WithCancel(ctx)
	syncCfg := viper.Sub("sync")

	eosCfg := viper.Sub("eos")

	creator := core.AN(eosCfg.GetString("creator.name"))
	log.Infow("加载eosCreater账户成功:" +  string(creator))

	signer := eos.NewKeyBag()
	eosCreaterPrivate := eosCfg.GetString("creator.privateKey")
	err := signer.ImportPrivateKey(eosCreaterPrivate)
	if err != nil {
		log.Errorf("ImportPrivateKey error:", err)
		return
	}
	log.Infow("加载eosCreater私钥成功...")

	eosUrl := viper.GetString("eos.url")
	log.Infow("EOS连接地址:" + eosUrl)

	eosClient := eosContract.NewEosClient(eosUrl)
	eosClient.SetSigner(creator, signer)
	log.Infow("生成eosClient成功...")

	ethClient, err := ethereum.NewEthClient(ethereum.ETH_CONN_ADDRESS)
	if err != nil {
		log.Errorf("生成ethClient失败", "error", err)
		return
	}
	log.Infow("生成ethClient成功...")

	syncer := Syncer{
		eosClient: eosClient,
		ethClient: ethClient,

		ctx:    ctx,
		cancel: cancel,

		limitContract:  uint32(syncCfg.GetInt32("limitContract")),
		limitDB:        syncCfg.GetInt("limitDB"),
		batchInterval:  syncCfg.GetInt("batchInterval"),
		updateInterval: syncCfg.GetInt("updateInterval"),

		tokenID: uint64(syncCfg.GetInt64("tokenID")),
	}

	syncer.wg.Add(1)
	go syncer.syncTokenTable()

	syncer.wg.Wait()
	log.Infow("need to update config field", "sync tokenID", syncer.tokenID)
}

func (s *Syncer) syncTokenTable() {
	log := log.Context("syncTokenTable")
	defer s.wg.Done()
	log.Info("syncTokenTable begin")
	defer log.Info("syncTokenTable end")
	for {
		err := s.syncEosToken()
		if err != nil {
			return
		}

		err = s.syncEthToken()
		if err != nil {
			return
		}
	}
}

func (s *Syncer) syncEosToken() error {
	var tokens []*TokenTable
	var err error
	log.Infow("开始同步EOS合约Token表...")
	tokens, err = s.getTokenTable(strconv.FormatUint(s.tokenID, 10), s.limitContract)
	if err != nil {
		log.Errorw("Sync stop: get contract table fail.", "error", err)
		s.cancel()
		return err
	}
	log.Infow("查询到的eosToken表数据", "tokens", &tokens)
	if len(tokens) > 0 {
		var insertTokens = make([]tokensDB.Tokens, 0)
		for _, token := range tokens {
			log.Infow("开始处理eosToken", "INFO", *token)
			tContract, err := strconv.ParseUint(token.TokenContract, 10, 64)
			if err != nil {
				log.Errorw("Parse tokenContract to uint64", "error", err)
				s.cancel()
				return err
			}
			tContractStr := eos.NameToString(tContract)

			has, err := tokensDB.QueryTokenExistFromDB(tContractStr, token.Symbol, configs.INEOS)
			if err != nil {
				log.Errorw("Sync stop: database error", "error", err)
				s.cancel()
				return err
			}

			if !has {
				log.Infow("此代币数据库不存在，准备插入数据库")

				tDecimal, err := strconv.Atoi(token.Decimal)
				if err != nil {
					log.Errorw("Atoi token.Decimal", "error", err)
					s.cancel()
					return err
				}

				iToken := tokensDB.Tokens{
					TokenID:       token.TokenID,
					Symbol:        token.Symbol,
					TokenContract: tContractStr,
					MarketCode:    configs.INEOS,
					Decimal:       uint(tDecimal),
					OrderDecimal:  uint(tDecimal),
					IsCancel:      token.IsCancle,
				}

				insertTokens = append(insertTokens, iToken)

				if iToken.Symbol == configs.EOSStr {
					//插入EOS跨链需要用的EOS代币
					iToken.MarketCode = configs.EOSETH
					insertTokens = append(insertTokens, iToken)

					//插入ETH跨链需要的EOS代币
					iToken.TokenContract = configs.AddressNil
					insertTokens = append(insertTokens, iToken)
				}
			}
		}

		err := tokensDB.InsertTokens(&insertTokens)
		if err != nil {
			log.Errorw("Sync stop: insert database error", "error", err)
			s.cancel()
			return err
		}
		s.tokenID = tokens[len(tokens)-1].TokenID + 1

		log.Infow("update eosTokenID to", "tokenID", s.tokenID)
	}
	timeChan := time.After(time.Duration(s.batchInterval) * time.Second)
	select {
	case <-timeChan:
		break
	case <-s.ctx.Done():
		log.Warnw("Stop sync", "error", s.ctx.Err())
		return bdexerrors.ErrNotFound
	}
	return nil
}

func (s *Syncer) syncEthToken() error {
	tokens := make([]EthTokenTable, 0)
	tokens, err := s.GetTokenTable()
	if err != nil {
		log.Errorw("查询eth链上Token表失败", "error", err)
		return err
	}
	log.Infow("查询到的ethToken表数据", "tokens", &tokens)

	if len(tokens) > 0 {
		var insertTokens = make([]tokensDB.Tokens, 0)
		for _, token := range tokens {
			log.Infow("开始处理ethToken", "INFO", token)
			has, err := tokensDB.QueryTokenExistFromDB(string(token.TokenContract), token.Symbol, int(token.MarketCode))
			if err != nil {
				log.Errorw("Sync stop: database error", "error", err)
				s.cancel()
				return err
			}
			if !has {
				log.Infow("此代币数据库不存在，准备插入数据库")
				iToken := tokensDB.Tokens{
					TokenID:       token.TokenID,
					Symbol:        token.Symbol,
					TokenContract: token.TokenContract,
					MarketCode:    int(token.MarketCode),
					Decimal:       uint(token.Decimal),
					OrderDecimal:  configs.ETHOrderDecimal,
					IsCancel:      token.IsCancel,
				}
				insertTokens = append(insertTokens, iToken)

				//如果是ETH则再添加一种token给跨链用
				if iToken.Symbol == configs.ETHStr {
					iToken.MarketCode = configs.EOSETH
					insertTokens = append(insertTokens, iToken)
				}

				//添加Bde代币
			}
		}

		err := tokensDB.InsertTokens(&insertTokens)
		if err != nil {
			log.Errorw("Sync stop: insert database error", "error", err)
			s.cancel()
			return err
		}
		s.tokenID = tokens[len(tokens)-1].TokenID + 1
		log.Infow("update tokenID to", "tokenID", s.tokenID)
	}
	timeChan := time.After(time.Duration(s.batchInterval) * time.Second)
	select {
	case <-timeChan:
		break
	case <-s.ctx.Done():
		log.Warnw("Stop sync", "error", s.ctx.Err())
		return bdexerrors.ErrNotFound
	}
	return nil
}

//从链上拿取所有代币
func (conn *Syncer) GetTokenTable() ([]EthTokenTable, error) {
	tokens := make([]EthTokenTable, 0)

	addressArray, err := conn.ethClient.GetTokenAddressArray()
	if err != nil {
		log.Errorw(".GetTokenAddressArray error：", "error", err)
		return nil, err
	}
	for _, v := range addressArray {
		fmt.Println("addressArray", v.String())
		resp, err := conn.ethClient.GetTokenInfoByAddress(v)
		if err != nil {
			log.Errorw("GetTokenInfoByAddress error", "error", err)
			return nil, err
		}
		for _, v := range resp.MarketCodeArray {
			if v > 0 {
				tokenInfo := EthTokenTable{
					Symbol:        resp.TokenName,
					TokenContract: resp.TokenAddress.String(),
					Decimal:       resp.Decimals,
					MarketCode:    v,
				}
				tokens = append(tokens, tokenInfo)
			}
		}
	}

	fmt.Println("ETH TOKENS :", tokens)
	return tokens, nil
}
