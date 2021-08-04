package core

import (
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	ethContract "bdex.in/bdex/bdex-ex-backend/contract/ethereum"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

//解析Bde订单数据结构体
type DepositsBdeDecoder struct {
	OrderId     uint64
	UserAddress common.Address
	Receiver    string
	TradeType   bool
	SellAmount  *big.Int
	SellAddress common.Address
	Price       *big.Int
	CreateTime  *big.Int
}

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
func SubmitBdeOrder(o *matchpool.CommonOrder, order deleorder.DeleOrder, count, dealPrice float64, orderid uint64) (string, error) {
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
	buy, err := ComputerBdeSubmitBuyAmount(order.MarketType, order.TradeType, count, dealPrice)
	if err != nil {
		log.Errorf("computerSubmitBuyAmount error:", err)
		return "", err
	}

	var selldecimal uint
	var buyDecimal uint

	//TODO：添加新的交易区修改
	tradetype, err := strconv.ParseInt(order.TradeType, 10, 64)
	if err != nil {
		return "", err
	}
	if tradetype == configs.BUYInt { //买单
		selldecimal = 18
		buyDecimals, err := tokens.GetEthDecimalsFromMem(order.BuyTokenContract)
		if err != nil {
			return "", err
		}
		buyDecimal = buyDecimals
	} else { //卖单
		buyDecimal = 18
		sellDecimals, err := tokens.GetEthDecimalsFromMem(order.SellTokenContract)
		if err != nil {
			return "", err
		}
		selldecimal = sellDecimals
	}

	//发送到链上
	fmt.Println("********************************")
	fmt.Println("交易数据:", id, sell, selldecimal, buy, buyDecimal)
	//TODO:确认合约已经修改
	submit := ethContract.NewSubmit(id, addressTo, sell, selldecimal, buy, buyDecimal)
	fmt.Println("处理后的数据：", submit.OrderId, submit.To.String(), submit.Sell.String(), submit.Buy.String(), submit.Fee.String())
	fmt.Println("********************************")

	//获取账户空闲队列的第一个账户
	account := EXCore.EAL.DelFirstLeisureAccount()
	//账户使用完后再次插入队列
	defer EXCore.EAL.InsertLeisureAccount(account)

	//TODO:确认合约已经修改
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
func ComputerBdeSubmitBuyAmount(markettype int, trandetypes string, count float64, dealPrice float64) (uint64, error) {
	fmt.Println("computerSubmitBuyAmount---tradetype:", trandetypes)
	t, err := strconv.ParseInt(trandetypes, 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseFloat error:", err)
		return 0, err
	}
	count = count * configs.FeeRate / configs.FeeBase

	//Bde买单
	if t == configs.BUYInt {
		sell := count * configs.FLOAT_MATH_RATE
		buy := uint64(sell / (dealPrice * configs.FLOAT_MATH_RATE) * configs.FLOAT_MATH_RATE)
		return buy, nil

	} else {
		//Bde卖单
		buy := dealPrice * configs.FLOAT_MATH_RATE * count
		return uint64(buy), nil
	}

}

/**
Description:解析Bde交易数据

Parameter:
    tx：已签名的交易

Return:
    DepositsSellDecoder结构体指针(从交易数据中解析出的订单数据结构体)
    错误
*/
//解析以太坊交易数据
func DecodeBdeOrderTransaction(tx string) (*DepositsBdeDecoder, error) {
	data, err := hexutil.Decode(tx)
	if err != nil {
		log.Errorw("Decode-Bde订单失败", "error", err)
		return nil, err

	}
	//解析transaction
	var tras types.Transaction
	err = rlp.DecodeBytes(data, &tras)
	if err != nil {
		log.Errorw("DecodeBytes-Bde订单失败", "error", err)
		return nil, err
	}

	//解析data
	contractAbi, err := abi.JSON(strings.NewReader(ethContract.ExchangeABI))
	meth, ok := contractAbi.Methods["sendBdeOrder"]
	if ok != true {
		log.Errorw("通过合约ABI解析交易传递参数失败", "error", err)
		return nil, err
	}
	//将合约参数解析到interface
	params, err := meth.Inputs.UnpackValues(tras.Data()[4:])
	if err != nil {
		log.Errorw("解包Bde订单交易内容失败", "error", err)
		return nil, err
	}
	return ChangeBdeOrderType(params), nil
}

/**
Description:通过BDE订单数据构建一个DeleOrder卖单结构体

Parameter:
	order:数据不完整的DeleOrder结构体，
	data:eth交易中打包的订单信息

Return:
*/
//通过ETH订单数据构建一个DeleOrder卖单结构体
func BuildBdeDeleOrder(order *deleorder.DeleOrder, data *DepositsBdeDecoder) error {
	order.OrderId = strconv.FormatUint(data.OrderId, 10)
	order.Name = data.UserAddress.String()
	order.CreateAt = time.Unix(data.CreateTime.Int64(), 0)
	order.UpdateAt = time.Unix(data.CreateTime.Int64(), 0)

	decimal, err := tokens.GetEthDecimalsFromMem(data.SellAddress.String())
	if err != nil {
		log.Errorw("获取代币精度失败：", err)
		return err
	}

	tokenSymbol, err := tokens.GetEthTokenSymbleFromMem(order.MarketType, data.SellAddress)
	if err != nil {
		log.Errorw("获取代币名称失败", "error", err)
		return err
	}

	matchSymbol := ""
	tradeType := ""
	buyTokenSymbol := ""
	sellTokenSymbol := ""
	buyTokenContract := ""
	sellTokenContract := ""
	price := strconv.FormatUint(data.Price.Uint64(), 10)
	sellTokenAmount := utils.ComputeBigintToStringOfSell(data.SellAmount, int64(decimal))

	if data.TradeType {
		tradeType = configs.SELLStr
		sellTokenSymbol = tokenSymbol
		sellTokenContract = data.SellAddress.String()
		buyTokenSymbol = configs.BDEStr + ",8"
		buyTokenContract = BdeAddress
		symbol, _ := utils.ParseSymbol(tokenSymbol)
		matchSymbol = symbol + "," + strconv.Itoa(order.MarketType)
	} else {
		tradeType = configs.BUYStr
		sellTokenSymbol = configs.BDEStr + ",8"
		sellTokenContract = BdeAddress
		buyTokenSymbol = tokenSymbol
		buyTokenContract = data.SellAddress.String()
		matchSymbol = configs.BDEStr + "," + strconv.Itoa(order.MarketType)
	}

	//计算委托第三方代币的总量
	fPrice, err := utils.AtofPrice(price)
	if err != nil {
		log.Errorw("utils.AtofPrice", "error", err)
		return err
	}

	_, money := utils.AtoAmount(sellTokenSymbol, sellTokenAmount)

	if order.TradeType == configs.BUYStr {
		order.DeleVol = money / fPrice
	} else {
		order.DeleVol = money
	}

	order.SellTokenSymbol = sellTokenSymbol
	order.SellTokenContract = sellTokenContract
	order.SellTokenAmount = sellTokenAmount

	order.BuyTokenSymbol = buyTokenSymbol
	order.BuyTokenContract = buyTokenContract
	order.BuyTokenAmount = "0"

	order.MatchSymbol = matchSymbol
	order.TradeType = tradeType

	order.Price = price
	order.Fee = strconv.FormatFloat(0, 'f', -1, 64)
	order.WithdrawAble = true
	order.ExSuccess = false
	order.TargetAddress = data.Receiver

	return nil
}

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
func ChangeBdeOrderType(v []interface{}) *DepositsBdeDecoder {
	return &DepositsBdeDecoder{
		OrderId:     v[0].(uint64),
		UserAddress: v[1].(common.Address),
		Receiver:    v[2].(string),
		TradeType:   v[3].(bool),
		SellAmount:  v[4].(*big.Int),
		SellAddress: v[5].(common.Address),
		Price:       v[6].(*big.Int),
		CreateTime:  v[7].(*big.Int),
	}
}
