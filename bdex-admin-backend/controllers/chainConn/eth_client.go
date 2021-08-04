package chainConn

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
)

var ExchangeClient *EthClient

type EthClient struct {
	*rpc.Client
	Eth            *ethclient.Client
	ExDataContract *ExchangeData
	Auth           *bind.TransactOpts
}

func init() {
	conn, err := NewEthClient()
	if err != nil {
		logs.Error("以太坊连接失败", err)
		return
	}
	ExchangeClient = conn
}

func NewEthClient() (*EthClient, error) {
	client, err := rpc.Dial(Address_Eth_Conn)
	if err != nil {
		return nil, err
	}
	logs.Info("以太坊网络连接成功")

	ethClient := ethclient.NewClient(client)

	//创建Exchange合约实例
	exchange, err := NewExchangeData(common.HexToAddress(Address_ExchnageData), ethClient)
	if err != nil {
		logs.Error("创建exchange合约实例失败", "error:", err)
		return nil, err
	}

	//加载Owner私钥
	PrivateKey, err := crypto.HexToECDSA(PrivateKey_ExchnageData_Owner)
	if err != nil {
		logs.Error("加载PrivateKey失败：", "error", err)
		return nil, err

	}
	logs.Info("加载PrivateKey成功", "INFO", "success")

	//构造Owner的签名对象
	Auth := bind.NewKeyedTransactor(PrivateKey)
	logs.Info("构造ETH的签名对象成功", "INFO", "success")
	Auth.GasLimit = 3000000

	return &EthClient{
		Client:         client,
		Eth:            ethClient,
		ExDataContract: exchange,
		Auth:           Auth,
	}, nil
}

//发送已签名交易上链
func (client *EthClient) SendSignedTransaction(signedTx string) (*types.Transaction, error) {
	var hash *common.Hash

	err := client.Call(hash, "eth_sendRawTransaction", padHexPrefix(signedTx))

	if err != nil {
		return nil, err
	}
	//解析交易
	s := types.Transaction{}
	data, err := hexutil.Decode(signedTx)
	if err != nil {
		return nil, err
	}
	err = rlp.DecodeBytes(data, &s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

//去链上获取GasPrice
func (client *EthClient) GetGasPrice() (string, error) {
	b, err := client.Eth.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

//去链上获取当前笔交易的Gas消耗
func (client *EthClient) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	return client.Eth.EstimateGas(context.Background(), msg)
}

//去链上获取当前区块高度
func (client *EthClient) GetBlockNumber() (string, error) {
	var b *big.Int
	err := client.Call(b, "eth_blockNumber")
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

//处理已签名交易，在交易数据前添加0x
func padHexPrefix(str string) string {
	if len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
		return str
	} else {
		return "0x" + str
	}
}

func (client *EthClient) IsTransactionSuccess(tx *types.Transaction) (bool, error) {
	logs.Info("正在确认交易是否上链....", "INFO", tx.Hash().String())
	begin := time.Now()
	receipt, err := bind.WaitMined(context.Background(), client.Eth, tx)
	if err != nil {
		logs.Error("等待挖矿失败", "error", err)
		return false, err
	}
	//再判断receipt.Status是否成功
	if receipt.Status != types.ReceiptStatusSuccessful {
		err = errors.New("挖矿失败")
		logs.Error("挖矿失败", "返回状态码错误：", receipt.Status)
		return false, err
	}
	logs.Info("挖矿结束", "返回状态码：", receipt.Status)
	logs.Info("此次挖矿耗时", "分钟", time.Now().Sub(begin).Minutes())
	return true, nil
}
