package cmd

import (
	ethContract "bdex.in/bdex/bdex-ex-backend/contract/ethereum"
	"bdex.in/bdex/bdex-ex-backend/controller"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/core/matchpool"
	"bdex.in/bdex/bdex-ex-backend/deleorder"
	"bdex.in/bdex/bdex-ex-backend/errorder"
	"bdex.in/bdex/bdex-ex-backend/ethsellorder"
	"bdex.in/bdex/bdex-ex-backend/log"
	"bdex.in/bdex/bdex-ex-backend/news"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/txrecords"
	"bdex.in/bdex/bdex-ex-backend/ws"
)

func InitLogDbSet() {
	log.Init()

	core.SetupLog()

	ethContract.SetupLog()

	ws.SetUpLog()

	matchpool.SetupLog()
	controller.SetupLog()

	deleorder.SetUpLog()
	deleorder.SetupDB()

	tokens.SetUpLog()
	tokens.SetupDB()

	txrecords.SetUpLog()
	txrecords.SetupDB()

	ethsellorder.SetUpLog()
	ethsellorder.SetupDB()

	news.SetUpLog()
	news.SetupDB()

	errorder.SetUpLog()
	errorder.SetupDB()

	core.Setup()
}
