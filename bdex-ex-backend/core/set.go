package core

import (
	"bdex.in/bdex/bdex-ex-backend/core/ethaccountlist"
	"strconv"

	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/core/matchquote"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/tokens"
)

var ID uint64

var EXCore *ExService

var BdeAddress string

func Setup() {

	exService, err := NewService()
	if err != nil {
		log.Errorw("NewService", "error", err)
		panic(err)
	}

	initApp(exService)
	EXCore = exService
}

func initApp(exService *ExService) {
	e := new(Engine)

	//初始化ID
	newestOrder, err := deleorder.GetNewestDeleOrder()
	if err != nil {
		log.Errorw("init ID", "error", err)
		panic(err)
	}

	if newestOrder.OrderId == "" {
		ID = 0
	} else {
		ID, err = strconv.ParseUint(newestOrder.OrderId[len(newestOrder.OrderId)-12:], 10, 64)
		if err != nil {
			log.Errorw("ParseID", "error", err)
			panic(err)
		}
	}

	//从数据库恢复token
	bdAddress, err := tokens.LoadTokenFromDB()
	if err != nil {
		log.Errorw("LoadTokenFromDB", "error", err)
		panic(err)
	}
	BdeAddress = bdAddress
	log.Infow("从数据库恢复Token成功...")

	//初始化订单池
	e.Handle = matchpool.NewMatchPool()
	log.Infow("初始化订单池成功...")

	//初始化行情池
	quoteUtil := matchquote.NewDynaQuoteUtil()
	e.Quote = quoteUtil
	log.Infow("初始化行情池成功...")

	//从数据库恢复订单
	err = matchpool.LoadOrderFromDB(e.Handle)
	if err != nil {
		log.Errorw("LoadOrderFromDB", "error", err)
		panic(err)
	}
	log.Infow("从数据库恢复订单池成功...")

	////从EthSellOrder数据库恢复队列
	//exService.Epl, err = ethputlist.LoadEthSellOrderFromDB()
	//if err != nil {
	//	log.Errorw("LoadEthSellOrderFromDB", "error", err)
	//	panic(err)
	//}
	//log.Infow("从数据库恢复EthSellOrder成功...")

	//从数据库恢复种子

	//初始化100个账号
	accountList, err := ethaccountlist.InitLeisureAccounts()
	if err != nil {
		log.Errorw("InitLeisureAccounts", "error", err)
		panic(err)
	}
	exService.EAL = accountList
	log.Infow("加载账号成功...")

	exService.Engine = e
}
