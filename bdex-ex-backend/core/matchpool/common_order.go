package matchpool

import (
	"bdex.in/bdex/bdex-ex-backend/bdexerrors"
	"container/list"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/configs"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/txrecords"
	"bdex.in/bdex/bdex-ex-backend/utils"

	math2 "github.com/ethereum/go-ethereum/common/math"
)

const ()

//内存中Order模型
type CommonOrder struct {
	Oid         uint64  //订单id
	Price       float64 //订单价格
	MatchSymbol string  //交易代币    1AAA   2DDD
	Code        string  //合约代码
	Name        string  //合约名称
	BsFlag      int32   //买卖标记
	Count       float64 //数量
	MarketCode  int     //市场代码 EOS链上 ETH EOS跨链 跨链 1，2，3，4
	OrderTime   string  //下单时间
	IsCancle    bool    //是否撤单 初始值=false,撤单或者成交=true
	IsHandle    bool    // true  正在处理
}

type MatchList struct {
	BuyList  list.List //sorted list asc
	SellList list.List //sorted list desc
}

//Parse DeleOrder
func ParseOrder(o deleorder.DeleOrder) (*CommonOrder, error) {
	commonOrder := CommonOrder{}
	id, err := strconv.ParseUint(o.OrderId, 10, 64)
	if err != nil {
		log.Errorw("convert order id error", "error", err)
		return nil, err
	}

	commonOrder.Oid = id
	commonOrder.Price, err = utils.AtofPrice(o.Price)

	if err != nil {
		log.Errorw("convert order price error", "error", err)
		return nil, err
	}

	c, err := strconv.Atoi(o.TradeType)
	commonOrder.BsFlag = int32(c)
	if err != nil {
		log.Errorw("convert order trader type error", "error", err)
		return nil, err
	}

	sellTokenSymbol, sellTokenMoney := utils.AtoAmount(o.SellTokenSymbol, o.SellTokenAmount)
	if sellTokenSymbol == "" {
		log.Errorw("convert order amount", "error", "AtoAmount error")
		return nil, bdexerrors.ErrInvalidArg
	}

	buyTokenSymbol, _ := utils.ParseSymbol(o.BuyTokenSymbol)
	//buyTokenSymbol, _ := AtoAmount(o.BuyTokenSymbol, o.BuyTokenAmount)
	if buyTokenSymbol == "" {
		log.Errorw("convert order amount", "error", "AtoAmount error")
		return nil, bdexerrors.ErrInvalidArg
	}

	if commonOrder.BsFlag == configs.BUYInt { //买单 amount/price
		//查询Tokens是否存在
		err = tokens.QueryTokenExistFromMem(o.BuyTokenContract, buyTokenSymbol, o.MarketType)
		if err != nil {
			log.Errorw("此Token不存在", "error", err)
			return nil, err
		}

		commonOrder.Name = o.BuyTokenContract
		//fmt.Println("+++++++++++++sellTokenMoney :", sellTokenMoney, commonOrder.Price, sellTokenMoney/commonOrder.Price, (sellTokenMoney*100000000)/(commonOrder.Price*100000000))
		commonOrder.Count = (sellTokenMoney * 1000000) / (commonOrder.Price * 1000000)

		commonOrder.Code = buyTokenSymbol
	} else { //卖单  amount*price
		err = tokens.QueryTokenExistFromMem(o.BuyTokenContract, buyTokenSymbol, o.MarketType)
		if err != nil {
			log.Errorw("此Token不存在", "error", err)
			return nil, err
		}
		commonOrder.Name = o.SellTokenContract
		commonOrder.Count = sellTokenMoney
		commonOrder.Code = sellTokenSymbol
	}
	//fmt.Println("commonOrder.Count:", commonOrder.Count)
	commonOrder.MarketCode = o.MarketType

	commonOrder.MatchSymbol = o.MatchSymbol

	commonOrder.OrderTime = o.CreateAt.Format(time.RFC3339)

	commonOrder.IsCancle = o.ExSuccess

	return &commonOrder, nil
}

func InsertTxRecordByCommOrder(o *CommonOrder, name string, count, dealPrice float64, tx_id string, matchOrderId uint64) error {
	fee := count / configs.FeeBase * configs.FeePersents

	txRecord := txrecords.TransactionRecords{
		OrderId:      strconv.FormatUint(o.Oid, 10),
		Name:         name,
		TradeType:    fmt.Sprint(o.BsFlag),
		MarketType:   o.MarketCode,
		Symbol:       o.Code,
		Count:        count,
		Price:        dealPrice,
		Fee:          fee,
		UpdateAt:     time.Now(),
		MatchOrderId: strconv.FormatUint(matchOrderId, 10),
		TxHash:       tx_id,
	}

	err := txrecords.InsertTxRecords(&txRecord)
	if err != nil {
		log.Errorw("插入交易记录失败", "error", err)
		return err
	}
	return nil
}

/**
Description:更新数据库和内存中订单数据

Parameter:
    order：订单数据结构体
    sell：卖方当次交易数量(count*10**8)
    buy：买方当次交易数量
    count：卖方当次交易数量
    fees：当次交易手续费
    selldecimal：卖方精度

Return:
	错误
*/
//更新数据库和内存中订单数据
func (o *CommonOrder) UpdataEthOrder(order deleorder.DeleOrder, sell uint64, buy uint64, count float64, fees *big.Int, selldecimal uint) error {
	//计算卖的数量和订单状态

	sells, err := strconv.ParseUint(order.SellTokenAmount, 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseInt error", err)
		return err
	}

	if sells == sell {
		order.WithdrawAble = false
		order.ExSuccess = true
		order.Status = configs.SUCCESS
	} else if computeOrderIsTooSmall(sells, sell, selldecimal) {
		order.WithdrawAble = false
		order.ExSuccess = true
		order.Status = configs.SUCCESS
	} else {
		order.WithdrawAble = true
		order.ExSuccess = false
		order.Status = configs.UNSUCCESS
	}

	order.SellTokenAmount = strconv.FormatUint(sells-sell, 10)

	//计算买的数量
	buys, err := strconv.ParseUint(order.BuyTokenAmount, 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseInt error", err)
		return err
	}

	order.BuyTokenAmount = strconv.FormatUint(buys+buy, 10)

	//计算fee

	Fee, err := computeSubmitFee(fees, selldecimal)
	if err != nil {
		fmt.Println("computeSubmitFee error", err)
		return err
	}
	order.Fee = Fee

	order.UpdateAt = time.Now()

	err = deleorder.UpdateDeleOrderById(&order)
	if err != nil {
		log.Errorw("Update EthOrder", "error", err)
		return err
	}
	//如果是买单，o.count是buy的数量
	if o.BsFlag == 0 {
		o.Count -= count / o.Price
		o.IsCancle = order.ExSuccess
		if order.ExSuccess == true {
			o.Count = 0
		}
	} else {
		//更新内存中状态
		o.Count = o.Count - count
		o.IsCancle = order.ExSuccess
		if order.ExSuccess == true {
			o.Count = 0
		}
	}

	return nil
}

/**
Description:计算手续费（float64）

Parameter:
    fee：原始手续费数量
    decimal：精度

Return:
    手续费（float64转string）
	错误
*/
//计算手续费（float64）
func computeSubmitFee(fee *big.Int, decimal uint) (string, error) {
	a := new(big.Int)
	a.Mul(fee, math2.BigPow(10, 8))
	a.Div(a, math2.BigPow(10, int64(decimal)))
	f := a.Int64()
	fe := float64(f) / math.Pow(10, 8)
	fees := strconv.FormatFloat(fe, 'f', -1, 64)

	return fees, nil

}

/**
Description:判断订单交易数量是否太少

Parameter:
    sells：卖方剩余数量
    sell：当次交易数量
    selldecimal：精度
Return:
    手续费（float64转string）
	错误
*/
//判断订单交易数量是否太少
func computeOrderIsTooSmall(sells uint64, sell uint64, selldecimal uint) bool {
	if sells == sell {
		return false
	}
	var a float64

	a = float64(sells - sell)
	a = a / configs.FLOAT_MATH_RATE
	a = a * 10000
	if a <= 1 {
		return true
	}
	return false
}
