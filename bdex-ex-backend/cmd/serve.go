package cmd

import (
	"fmt"
	"time"

	"bdex.in/bdex/bdex-ex-backend/controller"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/log"
	"bdex.in/bdex/bdex-ex-backend/tokens"
	"bdex.in/bdex/bdex-ex-backend/ws"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve for the whole rpc service",
	Long: `A longer description that spans multiple lines and likely contains examples
			and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		InitLogDbSet()

		//开启Eth缓冲下单
		//go core.CacherEthSellOrder()

		//account := core.EXCore.EAL.DelFirstLeisureAccount()
		//fmt.Println("获取到的account为",account.PublicKey)
		//account = core.EXCore.EAL.DelFirstLeisureAccount()
		//fmt.Println("获取到的account为",account.PublicKey)
		//account = core.EXCore.EAL.DelFirstLeisureAccount()
		//fmt.Println("获取到的account为",account.PublicKey)
		//account = core.EXCore.EAL.DelFirstLeisureAccount()
		//fmt.Println("获取到的account为",account.PublicKey)

		//开启撮合引擎
		core.Match()

		//开启api get 数据获取
		go func() {
			for {
				quoteDatas, err := controller.GetAllQuoteData()
				if err != nil {
					log.Errorw("cmd GetAllQuoteData", "error", err)
					panic(err)
				}
				controller.AllQuoteData = quoteDatas
				timeChan := time.After(time.Duration(10) * time.Second)
				select {
				case <-timeChan:
					continue
				}
			}
		}()

		////开启syncToken
		go func() {
			now := time.Now()
			for {
				if time.Now().After(now.Add(time.Hour * 12)) {
					_, err := tokens.LoadTokenFromDB()
					if err != nil {
						log.Errorw("client.LoadTokenFromDB()", "error", err)
					}
					now = now.Add(time.Hour * 12)
				}
			}
		}()

		fmt.Println(`
  ____        _                   ______                 _                                    
 |  _ \      | |                 |  ____|               | |                                   
 | |_) |   __| |   ___  __  __   | |__    __  __   ___  | |__     __ _   _ __     __ _    ___ 
 |  _ <   / _  |  / _ \ \ \/ /   |  __|   \ \/ /  / __| | '_ \   / _  | | '_ \   / _  |  / _ \
 | |_) | | (_| | |  __/  >  <    | |____   >  <  | (__  | | | | | (_| | | | | | | (_| | |  __/
 |____/   \__,_|  \___| /_/\_\   |______| /_/\_\  \___| |_| |_|  \__,_| |_| |_|  \__, |  \___|
                                                                                  __/ |
                                                                                 |___/
`)

		//开启websocket服务
		ws.InitWSServer()
		//service.Serve()

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
