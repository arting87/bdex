package core

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	ethContract "bdex.in/bdex/bdex-ex-backend/contract/ethereum"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

/**
Description:解析以太坊交易数据

Parameter:
    tx：已签名的交易

Return:
    Transaction结构体指针(解析后以太坊的交易数据)
    DepositsSellDecoder结构体指针(从交易数据中解析出的订单数据结构体)
    错误
*/
//解析以太坊交易数据
/*func DecodeEthTransactionDepositsBuy(tx string) (*types.Transaction, *DepositsBuyDecoder, error) {
	data, err := hexutil.Decode(tx)
	if err != nil {
		return nil, nil, fmt.Errorf("eth hexutil.Decode error", err)

	}
	//解析transaction
	var tras types.Transaction
	err = rlp.DecodeBytes(data, &tras)
	if err != nil {
		return nil, nil, fmt.Errorf("eth rlp.DecodeBytes", err)
	}

	//解析data
	abi, err := abi.JSON(strings.NewReader(ethContract.ExchangeABI))
	meth, ok := abi.Methods["depositsOrder"]
	if ok != true {
		return nil, nil, fmt.Errorf("abi.Methods error", err)
	}
	//将合约参数解析到interface
	params, err := meth.Inputs.UnpackValues(tras.Data()[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("meth.Inputs.UnpackValues", err)
	}
	return &tras, ChangeBuyOrderType(params), nil
}*/

/**
Description:解析以太坊交易数据

Parameter:
    tx：已签名的交易

Return:
    Transaction结构体指针(解析后以太坊的交易数据)
    DepositsSellDecoder结构体指针(从交易数据中解析出的订单数据结构体)
    错误
*/
//解析以太坊交易数据
/*func DecodeEthTransactionDepositsSell(tx string) (*types.Transaction, *DepositsSellDecoder, error) {
	data, err := hexutil.Decode(tx)
	if err != nil {
		return nil, nil, fmt.Errorf("eth hexutil.Decode error", err)

	}
	//解析transaction
	var tras types.Transaction
	err = rlp.DecodeBytes(data, &tras)
	if err != nil {
		return nil, nil, fmt.Errorf("eth rlp.DecodeBytes", err)
	}

	//解析data
	abi, err := abi.JSON(strings.NewReader(ethContract.ExchangeABI))
	meth, ok := abi.Methods["depositsSell"]
	if ok != true {
		return nil, nil, fmt.Errorf("abi.Methods error", err)
	}
	//将合约参数解析到interface
	params, err := meth.Inputs.UnpackValues(tras.Data()[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("meth.Inputs.UnpackValues", err)
	}
	return &tras, ChangeSellOrderType(params), nil
}*/

func DecodeEthTransactionDepositsOrder(tx string) (*DepositsOrderDecoder, error) {
	data, err := hexutil.Decode(tx)
	if err != nil {
		return nil, fmt.Errorf("eth hexutil.Decode error", err)

	}
	//解析transaction
	var tras types.Transaction
	err = rlp.DecodeBytes(data, &tras)
	if err != nil {
		return nil, fmt.Errorf("eth rlp.DecodeBytes", err)
	}
	//解析from
	addrFrom, err := types.Sender(types.HomesteadSigner{}, &tras)
	if err != nil {
		log.Errorw("解析From失败：", "error", err)
		return nil, err
	}

	//解析data
	abi, err := abi.JSON(strings.NewReader(ethContract.ExchangeABI))
	method, ok := abi.Methods["depositsOrder"]
	if ok != true {
		return nil, fmt.Errorf("abi.Methods error", err)
	}
	//将合约参数解析到interface
	params, err := method.Inputs.UnpackValues(tras.Data()[4:])
	if err != nil {
		return nil, fmt.Errorf("meth.Inputs.UnpackValues", err)
	}

	return ChangeEthOrderType(params, addrFrom), nil
}

/**
Description:发送链内卖单的交易

Parameter:
    tx：已签名的交易

Return:
    交易哈希
    错误
*/
//发送链内卖单的交易
func SendTransaction(tx string) (string, error) {
	txId, err := EXCore.ETH.SendEthTransaction(tx)
	if err != nil {
		return "", err
	}
	return txId, nil
}

/**
Description:解析ETH撤单数据

Parameter:
    tx：已签名的交易

Return:
    解析出来的订单id
    错误
*/
//解析ETH撤单数据
func DecodeEthTransactionWithdraw(tx string) (uint64, error) {
	data, err := hexutil.Decode(tx)
	if err != nil {
		return 0, fmt.Errorf("ethhexutil.Decode error", err)
	}
	//解析transaction
	var tras types.Transaction
	err = rlp.DecodeBytes(data, &tras)
	if err != nil {
		return 0, fmt.Errorf("eth rlp.DecodeBytes error", err)
	}

	//解析data
	abi, err := abi.JSON(strings.NewReader(ethContract.ExchangeABI))
	meth, ok := abi.Methods["withdrawal"]
	if ok != true {
		return 0, fmt.Errorf("abi.Methods error", err)
	}
	//将合约参数解析到interface
	params, err := meth.Inputs.UnpackValues(tras.Data()[4:])
	if err != nil {
		return 0, fmt.Errorf("meth.Inputs.UnpackValues", err)
	}

	return params[0].(uint64), nil
}

//解析买单数据结构体
/*type DepositsBuyDecoder struct {
	OrderId     uint64
	UserAddress common.Address
	Receiver    string
	TradeType   bool
	SellAmount  *big.Int
	Price       *big.Int
	BuyAddress  common.Address
	CreateTime  *big.Int
}*/

//解析卖单数据结构体
/*type DepositsSellDecoder struct {
	OrderId     uint64
	UserAddress common.Address
	Receiver    string
	TradeType   bool
	SellAmount  *big.Int
	SellAddress common.Address
	Price       *big.Int
	CreateTime  *big.Int
}*/

type DepositsOrderDecoder struct {
	OrderId      uint64
	UserAddress  common.Address
	Receiver     string
	TradeType    bool
	SellAmount   *big.Int
	SellContract common.Address
	Price        *big.Int
	BuyContract  common.Address
	MarketCode   uint64
}

/**
Description:通过ETH订单数据构建一个DeleOrder买单结构体

Parameter:
	order:数据不完整的DeleOrder结构体，
	data:eth交易中打包的订单信息

Return:
*/
//通过ETH订单数据构建一个DeleOrder买单结构体
/*func BuildEthDeleOrderBuy(order *deleorder.DeleOrder, data *DepositsBuyDecoder) error {

	order.OrderId = strconv.FormatUint(data.OrderId, 10)
	order.Name = data.UserAddress.String()
	order.CreateAt = time.Unix(data.CreateTime.Int64(), 0)
	order.UpdateAt = time.Unix(data.CreateTime.Int64(), 0)

	if data.TradeType { //卖单
		order.TradeType = configs.SELLStr
	} else { //买单
		order.TradeType = configs.BUYStr
	}

	order.SellTokenAmount = utils.ComputeBigIntToStringOfBuy(data.SellAmount)
	order.SellTokenSymbol = "ETH" + ",8"
	order.SellTokenContract = configs.AddressNil
	order.BuyTokenAmount = "0"

	buyTokenSymbol, err := tokens.GetEthTokenSymbleFromMem(order.MarketType, data.BuyAddress)
	if err != nil {
		log.Errorw("获取代币名称失败", "error", err)
		return err
	}
	order.BuyTokenSymbol = buyTokenSymbol

	//TODO：添加新的交易区修改,已审核过，无需修改
	if order.MarketType == configs.EOSETH {
		order.MatchSymbol = "ETH," + strconv.Itoa(order.MarketType)
	} else {
		s, _ := utils.ParseSymbol(order.BuyTokenSymbol)
		order.MatchSymbol = s + "," + strconv.Itoa(order.MarketType)
	}

	order.BuyTokenContract = data.BuyAddress.String()
	order.Price = strconv.FormatUint(data.Price.Uint64(), 10)
	order.Fee = strconv.FormatFloat(0, 'f', -1, 64)
	order.WithdrawAble = true
	order.ExSuccess = false
	order.TargetAddress = data.Receiver

	return nil
}*/

/**
Description:通过ETH订单数据构建一个DeleOrder卖单结构体

Parameter:
	order:数据不完整的DeleOrder结构体，
	data:eth交易中打包的订单信息

Return:
*/
//通过ETH订单数据构建一个DeleOrder卖单结构体
func BuildEthDeleOrder(order *deleorder.DeleOrder, data *DepositsOrderDecoder) error {
	order.OrderId = strconv.FormatUint(data.OrderId, 10)
	order.Name = data.UserAddress.String()
	order.CreateAt = time.Now()
	order.UpdateAt = time.Now()

	if data.TradeType {
		order.TradeType = configs.SELLStr
	} else {
		order.TradeType = configs.BUYStr
	}
	order.MarketType = int(data.MarketCode)
	order.SellTokenContract = data.SellContract.String()
	var SellTokenSymbol string
	var SellTokenAmount string
	var BuyTokenSymbol string
	var MatchSymbol string

	if data.SellContract.String() == configs.AddressNil {
		SellTokenSymbol = "ETH" + ",8"
		SellTokenAmount = utils.ComputeBigintToStringOfSell(data.SellAmount, configs.ETHDecimal)
	} else {
		decimalsSell, err := tokens.GetEthDecimalsByAddressAndMarketCode(data.SellContract.String(), int(data.MarketCode))
		if err != nil {
			log.Errorw("获取代币精度失败：", "error", err.Error())
			return err
		}
		sellSymble, err := tokens.GetEthTokenSymbleFromMem(int(data.MarketCode), data.SellContract)
		if err != nil {
			log.Errorw("获取代币名称失败：", "error", err.Error())
			return err
		}
		SellTokenSymbol = sellSymble
		SellTokenAmount = utils.ComputeBigintToStringOfSell(data.SellAmount, int64(decimalsSell))
	}

	if data.BuyContract.String() == configs.AddressNil {
		if data.MarketCode == 2 && data.TradeType && data.SellContract.String() == configs.AddressNil {
			BuyTokenSymbol = "EOS" + ",4"
		} else {
			BuyTokenSymbol = "ETH" + ",8"
		}
	} else {
		buySymble, err := tokens.GetEthTokenSymbleFromMem(int(data.MarketCode), data.BuyContract)
		if err != nil {
			log.Errorw("获取代币名称失败：", "error", err)
			return err
		}
		BuyTokenSymbol = buySymble
	}

	if data.TradeType {
		MatchSymbol = ParseSymble(SellTokenSymbol) + "," + strconv.Itoa(order.MarketType)
	} else {
		if data.BuyContract.String() == configs.AddressNil && data.SellContract.String() == configs.AddressNil {
			MatchSymbol = ParseSymble(SellTokenSymbol) + "," + strconv.Itoa(order.MarketType)
		} else {
			MatchSymbol = ParseSymble(BuyTokenSymbol) + "," + strconv.Itoa(order.MarketType)
		}
	}

	order.MatchSymbol = MatchSymbol
	order.SellTokenSymbol = SellTokenSymbol
	order.SellTokenAmount = SellTokenAmount
	order.BuyTokenSymbol = BuyTokenSymbol
	order.BuyTokenContract = data.BuyContract.String()
	order.BuyTokenAmount = "0"
	order.Price = strconv.FormatUint(data.Price.Uint64(), 10)
	order.Fee = strconv.FormatFloat(0, 'f', -1, 64)
	order.WithdrawAble = true
	order.ExSuccess = false
	order.TargetAddress = data.Receiver

	return nil
}

func ParseSymble(symble string) (string) {
	str := strings.Split(symble, ",")
	return str[0]
}

//

func ChangeEthOrderType(v []interface{}, address common.Address) *DepositsOrderDecoder {
	return &DepositsOrderDecoder{
		OrderId:      v[0].(uint64),
		UserAddress:  address,
		Receiver:     v[1].(string),
		TradeType:    v[2].(bool),
		SellAmount:   v[3].(*big.Int),
		SellContract: v[4].(common.Address),
		Price:        v[5].(*big.Int),
		BuyContract:  v[6].(common.Address),
		MarketCode:   v[7].(uint64),
	}
}

//将interface转换成结构体
/*func ChangeBuyOrderType(v []interface{}) *DepositsBuyDecoder {
	return &DepositsBuyDecoder{
		OrderId:     v[0].(uint64),
		UserAddress: v[1].(common.Address),
		Receiver:    v[2].(string),
		TradeType:   v[3].(bool),
		SellAmount:  v[4].(*big.Int),
		Price:       v[5].(*big.Int),
		BuyAddress:  v[6].(common.Address),
		CreateTime:  v[7].(*big.Int),
	}
}*/

/**
			uint64 orderid,
            address useraddress,
            string receiver,
            bool tradetype,
            uint256 sellamount,
            address selladdress,
            uint256 price,
            uint256 createTime
*/
//将interface转换成结构体
/*func ChangeSellOrderType(v []interface{}) *DepositsSellDecoder {
	return &DepositsSellDecoder{
		OrderId:     v[0].(uint64),
		UserAddress: v[1].(common.Address),
		Receiver:    v[2].(string),
		TradeType:   v[3].(bool),
		SellAmount:  v[4].(*big.Int),
		SellAddress: v[5].(common.Address),
		Price:       v[6].(*big.Int),
		CreateTime:  v[7].(*big.Int),
	}
}*/

/**
Description:撮合成功后submit以太坊订单

Parameter:
    order：订单数据结构体
    count；交易数量
    dealPrice：单价
    orderid：交易id

Return:
    交易哈希
	错误
*/
//submit 以太坊订单
func SubmitEthOrder(o *matchpool.CommonOrder, order deleorder.DeleOrder, count, dealPrice float64, orderid uint64) (string, error) {
	//查询交易对象的目标address
	fmt.Println("订单：", order.OrderId, orderid)
	addressTo, err := deleorder.GetTargetAddressFromDb(orderid)
	if err != nil {
		log.Errorf("GetTargetAddressFromDb error:", err)
		return "", err
	}
	//交易id
	id, err := strconv.ParseUint(order.OrderId, 10, 64)
	if err != nil {
		log.Errorf("orderid parseUint error:", err)
		return "", err
	}

	//卖方个数
	sell := uint64(count * configs.FLOAT_MATH_RATE)
	//买方个数
	buy, err := ComputerSubmitBuyAmount(order.MarketType, order.TradeType, count, dealPrice)
	if err != nil {
		log.Errorf("computerSubmitBuyAmount error:", err)
		return "", err
	}

	var SellDecimals uint
	var BuyDecimals uint

	if order.MarketType == configs.EOSETH { //跨链
		SellDecimals = 18
		BuyDecimals = 8

	} else if order.MarketType == configs.INETH { //链内
		tradeType, err := strconv.ParseInt(order.TradeType, 10, 64)
		if err != nil {
			return "", err
		}
		if tradeType == configs.BUYInt { //买单
			SellDecimals = 18
			buyDecimal, err := tokens.GetEthDecimalsFromMem(order.BuyTokenContract)
			if err != nil {
				return "", err
			}
			BuyDecimals = buyDecimal
		} else { //卖单
			BuyDecimals = 18
			sellDecimal, err := tokens.GetEthDecimalsFromMem(order.SellTokenContract)
			if err != nil {
				return "", err
			}
			SellDecimals = sellDecimal
		}

	} else {
		sellDecimal, err := tokens.GetEthDecimalsFromMem(order.SellTokenContract)
		if err != nil {
			log.Errorw("GetEthDecimalsFromMem error", "error", err)
			return "", err
		}
		buyDecimal, err := tokens.GetEthDecimalsFromMem(order.BuyTokenContract)
		if err != nil {
			log.Errorw("GetEthDecimalsFromMem error", "error", err)
			return "", err
		}
		SellDecimals = sellDecimal
		BuyDecimals = buyDecimal
	}


	//发送到链上
	fmt.Println("********************************")
	fmt.Println("交易数据:", id, sell, SellDecimals, buy, BuyDecimals)
	submit := ethContract.NewSubmit(id, addressTo, sell, SellDecimals, buy, BuyDecimals)
	fmt.Println("处理后的数据：", submit.OrderId, submit.To.String(), submit.Sell.String(), submit.Buy.String(), submit.Fee.String())
	fmt.Println("********************************")

	//获取账户空闲队列的第一个账户
	account := EXCore.EAL.DelFirstLeisureAccount()
	//账户使用完后再次插入队列
	defer EXCore.EAL.InsertLeisureAccount(account)

	txhash, err := EXCore.ETH.SubmitOrder(order.MarketType, submit, account)
	if err != nil {
		log.Errorf("connect to eth contract method submit error:", err)
		return "", err
	}

	//修改内存数据的应该提前，并且单线程处理

	//更新数据库
	/*err = o.UpdataEthOrder(order, sell, buy, count, submit.Fee, selldecimal)
	if err != nil {
		log.Errorf("updata eth database order error:", err)
		return "", err
	}*/

	return txhash, nil
}

/**
Description:计算买的代币数量

Parameter:
    trandetypes：买卖类型
    count：卖方交易数量
    dealPrice：单价

Return:
    买的代币数量
	错误
*/
//计算买的代币数量
func ComputerSubmitBuyAmount(markettype int, trandetypes string, count float64, dealPrice float64) (uint64, error) {
	fmt.Println("computerSubmitBuyAmount---tradetype:", trandetypes)
	t, err := strconv.ParseInt(trandetypes, 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseFloat error:", err)
		return 0, err
	}
	count = count * configs.FeeRate / configs.FeeBase
	//跨链卖单
	if markettype == configs.EOSETH && t == configs.SELLInt {
		buy := uint64((count * dealPrice) * 10000)
		fmt.Println("buy:", buy)
		return buy, nil
		//链内买单
	} else if markettype == configs.INETH && t == configs.BUYInt {
		sell := count * configs.FLOAT_MATH_RATE
		buy := uint64(sell / (dealPrice * configs.FLOAT_MATH_RATE) * configs.FLOAT_MATH_RATE)
		return buy, nil

		//链内卖单
	} else {
		buy := dealPrice * configs.FLOAT_MATH_RATE * count
		return uint64(buy), nil
	}

}

/**
Description:以太坊订单回滚

Parameter:
    orderbId：交易id

Return:
    交易哈希
	错误
*/
//以太坊订单回滚
func BackEthOrder(orderbId uint64) (string, error) {
	//获取账户空闲队列的第一个账户
	account := EXCore.EAL.DelFirstLeisureAccount()
	//账户使用完后再次插入队列
	defer EXCore.EAL.InsertLeisureAccount(account)

	txHash, err := EXCore.ETH.RollBackEthOrder(orderbId, account)
	if err != nil {
		log.Errorw("RollBackEthOrder", "error", err)
		return "", err
	}
	return txHash, nil
}

/**
Description:以太坊submit成功后转账

Parameter:
    o：CommonOrder结构体指针
    order：DeleOrder结构体
    count：此此交易数量
    prise：此次交易价格

Return:
    交易哈希
	错误
*/
//以太坊submit成功后转账
func TransNeedEthOrder(o *matchpool.CommonOrder, order deleorder.DeleOrder, count float64, prise float64) (string, error) {

	//获取账户空闲队列的第一个账户
	account := EXCore.EAL.DelFirstLeisureAccount()
	//账户使用完后再次插入队列
	defer EXCore.EAL.InsertLeisureAccount(account)
	//交易id
	id, err := strconv.ParseUint(order.OrderId, 10, 64)
	if err != nil {
		log.Errorw("orderid parseUint", "error", err)
		return "", err
	}

	txhash, err := EXCore.ETH.TransferEthMoney(id, account)
	if err != nil {
		log.Errorw("TransferEthMoney", "error", err, "id", id, )
		return "", err
	}

	//var selldecimal uint

	//TODO:添加新的交易区判断
	//if order.MarketType == configs.EOSETH { //跨链
	//	selldecimal = 18
	//
	//	//} else if order.MarketType == configs.INETH { //链内
	//} else {
	//	tradetype, err := strconv.ParseInt(order.TradeType, 10, 64)
	//	if err != nil {
	//		return "", err
	//	}
	//	if tradetype == configs.BUYInt { //买单
	//		selldecimal = 18
	//
	//	} else { //卖单
	selldecimal, err := tokens.GetEthDecimalsFromMem(order.SellTokenContract)
	if err != nil {
		return "", err
	}
	//selldecimal = sellDecimals
	//}
	//}

	sell := uint64(count * configs.FLOAT_MATH_RATE)
	//买方个数
	buy, err := ComputerSubmitBuyAmount(order.MarketType, order.TradeType, count, prise)
	if err != nil {
		log.Errorf("computerSubmitBuyAmount error:", err)
		return "", err
	}
	fee := computerFee(sell, selldecimal)

	err = o.UpdataEthOrder(order, sell, buy, count, fee, selldecimal)
	if err != nil {
		log.Errorf("updata eth database order error:", err)
		return "", err
	}
	return txhash, nil
}

/**
Description:计算手续费

Parameter:
    sellamount：交易的代币数量
    order：交易的代币精度

Return:
    手续费
*/
//计算手续费
func computerFee(sellamount uint64, selldecimal uint) *big.Int {
	compute_sell := ethContract.ComputeSellAmount(int64(sellamount), int64(selldecimal))
	compute_fee := ethContract.ComputerSellAndFee(sellamount, selldecimal, compute_sell)
	return compute_fee
}
