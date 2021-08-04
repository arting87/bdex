package ws

import (
	"errors"
	"strconv"
	"time"

	"bdex.in/bdex/bdex-ex-backend/controller"
	"bdex.in/bdex/bdex-ex-backend/core"
	"bdex.in/bdex/bdex-ex-backend/utils"

	news2 "bdex.in/bdex/bdex-ex-backend/news"
	"github.com/mitchellh/mapstructure"
)

const (
	ping = "ping"
	api  = "api"
	user = "user"
	news = "news"

	update      = "update"
	send        = "send"
	get         = "get"
	sendeos     = "sendeos"
	sendeth     = "sendeth"
	sendbde     = "sendbde"
	withdraweos = "withdraweos"
	withdraweth = "withdraweth"
	withdrawbde = "withdrawbde"
	getid       = "getid"

	newslist   = "newsList"
	onenews    = "onenews"
	typenews   = "typenews"
	insertnews = "insertnews"
	updatenews = "updatenews"
	deletenews = "deletenews"

	k = "k"
)

// 每隔1秒, 检查一次连接是否健康
func (wsConnection *WSConnection) heartbeatChecker() {
	for {
		log.Infow("--------心跳检测--------", "connAddr", wsConnection.connAddr)
		timer := time.After(time.Duration(WsHeartbeatInterval) * time.Second)
		select {
		case <-timer:
			if !wsConnection.IsAlive() {
				wsConnection.Close()
				goto EXIT
			}
		case <-wsConnection.closeChan:
			wsConnection.Close()
			goto EXIT
		}
	}

EXIT:
	// 确保连接被关闭
}

// 处理PING请求
func (wsConnection *WSConnection) handlePing(bizReq *BizMessage) (bizResp *BizMessage) {
	wsConnection.KeepAlive()

	bizResp = &BizMessage{
		//	Data: "PONG",
		Namespace: "PONG",
		Action:    "PONG",
		Data:      nil,
	}
	return
}

// 处理websocket请求
func (wsConnection *WSConnection) WSHandle() {
	var (
	//message BizMessage
	//bizResp *BizMessage
	//err error
	)

	// 心跳检测线程
	go wsConnection.heartbeatChecker()

	userGetContexts := make([]chan int, 0)
	apiGetKContexts := make([]chan int, 0)

	// 请求处理协程
	for {
		message, err := wsConnection.ReadMessage()
		if err != nil {
			log.Errorw("读取信息失败", "error", err)
			goto ERR
		}
		//log.Infow("读取到的信息", "connAddr", wsConnection.connAddr, "message", message)

		//bizResp = nil
		// 请求串行处理
		switch message.Namespace {
		case ping:
			switch message.Action {
			case send:
				bizResp := wsConnection.handlePing(&message)
				_ = wsConnection.sendResp(bizResp)
			}
			//处理新闻开始
		case news:
			switch message.Action {
			case newslist:
				go func() {
					bizResp := BizMessage{
						Namespace: news,
						Action:    newslist,
					}
					newresp, err := news2.GetNewsList()
					if err != nil {
						log.Errorw("GetNewsList error:", err)
						bizResp.Data = controller.NewsError{Error: "GetNewsList error", Status: false}
					} else {
						bizResp.Data = newresp
					}
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("news get GetNewsList", "error", err)
						return
					}
				}()
			case onenews:
				var oneNewsData controller.GetOneNewsData
				reqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(reqMap, &oneNewsData)
				if err != nil {
					log.Errorw("mapstructure.Decode", "error", err)
					return
				}
				go func() {
					bizResp := BizMessage{
						Namespace: news,
						Action:    onenews,
					}
					newresp, err, istrue := news2.GetNewsById(oneNewsData.Id)

					if err != nil {
						log.Errorw("GetNewsById error", err)
						bizResp.Data = controller.NewsError{Error: "GetNewsById error", Status: false}
					}
					if !istrue {
						log.Debug("this news is not exit")
						bizResp.Data = controller.NewsError{Error: "this news is not exit", Status: false}
					}
					if (err == nil && istrue) {
						bizResp.Data = newresp
					}
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("news get onenews", "error", err)
						return
					}
				}()
			case typenews:
				var typeNewsData controller.GetTypeNewsData
				reqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(reqMap, &typeNewsData)
				if err != nil {
					log.Errorw("mapstructure.Decode", "error", err)
					return
				}
				go func() {
					bizResp := BizMessage{
						Namespace: news,
						Action:    typenews,
					}
					newresp, err := news2.GetNewsByTypes(typeNewsData.TypeCode)
					if err != nil {
						log.Errorw("GetNewsByTypes error", err)
						bizResp.Data = controller.NewsError{Error: "GetNewsByTypes error", Status: false}
					} else {
						bizResp.Data = newresp
					}
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("news get typenews", "error", err)
						return
					}
				}()
			case insertnews:
				var insertNewsData controller.InsertNewsData
				reqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(reqMap, &insertNewsData)
				if err != nil {
					log.Errorw("mapstructure.Decode", "error", err)
					return
				}
				go func() {
					bizResp := BizMessage{
						Namespace: news,
						Action:    insertnews,
					}
					err := news2.InsertNews(insertNewsData.Title, insertNewsData.TypeCode, insertNewsData.Content)
					if err != nil {
						log.Errorw("InsertNews", "error", err)
						bizResp.Data = controller.NewsError{Error: "InsertNews error", Status: false}
					} else {
						bizResp.Data = controller.NewsSuccess{Status: true}
					}
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("news insertnews", "error", err)
						return
					}
				}()
			/*case updatenews:
			var updataNewsData controller.UpdataNewsData
			reqMap := message.Data.(map[string]interface{})
			err = mapstructure.Decode(reqMap, &updataNewsData)
			if err != nil {
				log.Errorw("mapstructure.Decode", "error", err)
				return
			}
			go func() {
				bizResp := BizMessage{
					Namespace: news,
					Action:    updatenews,
				}
				err := news2.UpdateNewsById(updataNewsData.Id, updataNewsData.Title, updataNewsData.Type, updataNewsData.Content)
				if err != nil {
					log.Errorw("UpdateNewsById", "error", err)
					bizResp.Data = controller.NewsError{Error: "UpdateNewsById error", Status: false}
				} else {
					bizResp.Data = controller.NewsSuccess{Status: true}
				}
				err = wsConnection.sendResp(&bizResp)
				if err != nil {
					log.Errorw("news updatenews", "error", err)
					return
				}
			}()*/
			case deletenews:
				var deleteNewsData controller.DeleteNewsData
				reqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(reqMap, &deleteNewsData)
				if err != nil {
					log.Errorw("mapstructure.Decode", "error", err)
					return
				}
				go func() {
					bizResp := BizMessage{
						Namespace: news,
						Action:    deletenews,
					}
					err := news2.DeleteNewsById(deleteNewsData.Id)
					if err != nil {
						log.Errorw("DeleteNewsById", "error", err)
						bizResp.Data = controller.NewsError{Error: "DeleteNewsById error", Status: false}
					} else {
						bizResp.Data = controller.NewsSuccess{Status: true}
					}
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("news deletenews", "error", err)
						return
					}

				}()

			}
			//处理新闻结束
		case api:
			switch message.Action {
			case get:
				go func() {
					for {
						bizResp := BizMessage{
							Namespace: api,
							Action:    get,
							Data:      controller.AllQuoteData,
						}
						err = wsConnection.sendResp(&bizResp)
						if err != nil {
							log.Errorw("api get sendResp", "error", err)
							return
						}
						timeChan := time.After(time.Duration(10) * time.Second)
						select {
						case <-timeChan:
							continue
						}
					}
				}()
			case k: //请求K线图数据
				go func() {
					var kLineReqData controller.KLineReqData
					a := make([]controller.KLineRespData, 0)
					bizResp := BizMessage{
						Namespace: api,
						Action:    k,
						Data:      a,
					}
					if message.Data == nil {
						err = errors.New("api k data is nil")
						bizResp.Data = err
						_ = wsConnection.sendResp(&bizResp)
						log.Errorw("api k sendResp", "error", err)
						return
					}

					kLineReqMap := message.Data.(map[string]interface{})
					if kLineReqMap["timeType"] == nil || kLineReqMap["symbol"] == nil || kLineReqMap["marketCode"] == nil {
						err = errors.New("api k data primary has nil")
						bizResp.Data = err
						_ = wsConnection.sendResp(&bizResp)
						return
					}

					newChan := make(chan int)
					//关闭上一个协程
					for _, lastChan := range apiGetKContexts {
						lastChan <- 1
					}
					apiGetKContexts = make([]chan int, 0)
					apiGetKContexts = append(apiGetKContexts, newChan)

					//解析请求数据
					err = mapstructure.Decode(kLineReqMap, &kLineReqData)
					if err != nil {
						log.Errorw("api k Decode userReqMap", "error", err)
						bizResp.Data = err
						_ = wsConnection.sendResp(&bizResp)
						return
					}

					if kLineReqData.TimeType <= 0 {
						err = errors.New("api k data is illegality")
						bizResp.Data = err
						_ = wsConnection.sendResp(&bizResp)
						return
					}

					//修改返回的action
					var ktype string
					switch kLineReqData.MarketCode {
					case 1:
						ktype = kLineReqData.Symbol + "EOS" + "-" + strconv.Itoa(kLineReqData.TimeType)
					case 2:
						ktype = kLineReqData.Symbol + "ETH" + "-" + strconv.Itoa(kLineReqData.TimeType)
					case 3:
						ktype = kLineReqData.Symbol + "EOS" + "-" + strconv.Itoa(kLineReqData.TimeType)
					case 4:
						ktype = kLineReqData.Symbol + "EOS" + "-" + strconv.Itoa(kLineReqData.TimeType)
					default:
						ktype = "Unknown"
						log.Errorw("apicline req data", "error", err)
						err = errors.New("api k unknown marketCode")
						bizResp.Data = err
						_ = wsConnection.sendResp(&bizResp)
						return
					}
					bizResp.Action = ktype + "-INIT"

					//获取间隔对应时间所有的交易记录,从最早拿到当前时间
					resp, baseTime, err := controller.GetKLineData(kLineReqData, time.Now().Add(time.Minute), -1)
					if err != nil {
						return
					}
					if resp == nil {
						log.Errorw("api k sendResp", "error", err)
						return
					}
					bizResp.Data = resp
					err = wsConnection.sendResp(&bizResp)
					if err != nil {
						log.Errorw("api k sendResp", "error", err)
						return
					}
					//根据时间获取最近的一次间隔时间对应的Kline数据,从传入时间拿到当前时间
					var lastClosePrice float64
					if len(resp) > 0 {
						lastClosePrice = resp[len(resp)-1].ClosePrice
					} else {
						lastClosePrice = -10
					}
					for {
						if time.Now().After(baseTime.Add(time.Duration(kLineReqData.TimeType) * time.Minute)) {
							resp, baseTime, err = controller.GetKLineData(kLineReqData, baseTime, lastClosePrice)
							if err != nil {
								log.Errorw("apicline GetKLineData", "error", err)
								bizResp.Data = err
								err = wsConnection.sendResp(&bizResp)
								if err != nil {
									log.Errorw("api k sendResp", "error", err)
									return
								}
								return
							}
							if resp == nil {
								err = errors.New("this token tx is nil")
								bizResp.Data = err
								err = wsConnection.sendResp(&bizResp)
								if err != nil {
									log.Errorw("api k sendResp", "error", err)
									return
								}
								return
							}
							bizResp.Action = ktype + "-UPDATE"
							bizResp.Data = resp
							err = wsConnection.sendResp(&bizResp)
							if err != nil {
								log.Errorw("api k sendResp", "error", err)
								return
							}
						} else {
							select {
							case <-newChan:
								log.Infow("this chan api k is close", "connAddr", wsConnection.connAddr)
								return
							default:
								continue
							}
						}
					}
				}()
			}
		case user:
			switch message.Action {
			case get: //用户获取历史订单
				newChan := make(chan int)
				go func() {
					//关闭上一个协程
					for _, lastChan := range userGetContexts {
						lastChan <- 1
					}
					userGetContexts = make([]chan int, 0)
					userGetContexts = append(userGetContexts, newChan)
					//LoopUserGet = false
					var userReqData controller.UserReqData
					userReqSlice := message.Data.([]interface{})
					if len(userReqSlice) == 0 {
						newChan <- 1
						userGetContexts = make([]chan int, 0)
					}
					for {
						select {
						case <-newChan:
							return
						default:
							datas := make([]controller.UserRespDatas, 0)
							for _, userReq := range userReqSlice {
								userReqMap := userReq.(map[string]interface{})
								err = mapstructure.Decode(userReqMap, &userReqData)
								if err != nil {
									log.Errorw("user get Decode userReqMap", "error", err)
									return
								}
								data, err := controller.GetUserOrderRecords(&userReqData)
								if err != nil {
									log.Errorw("GetUserOrderRecords", "error", err)
									return
								}
								datas = append(datas, controller.UserRespDatas{
									Account:   userReqData.Account,
									ChainCode: userReqData.ChainCode,
									Datas:     data,
								})
							}
							bizResp := BizMessage{
								//	Data: "PONG",
								Namespace: user,
								Action:    get,
								Data:      datas,
							}
							err = wsConnection.sendResp(&bizResp)
							if err != nil {
								return
							}
							timeChan := time.After(time.Duration(1) * time.Second)
							select {
							case <-timeChan:
								continue
							}
						}
					}
				}()
			case sendeos: //EOS下单
				var eosPutOrderData controller.EosPutOrderData
				userReqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(userReqMap, &eosPutOrderData)
				if err != nil {
					log.Errorw("user get Decode userReqMap", "error", err)
					goto ERR
				}

				orderId, err := controller.UserSendEOSOrder(&eosPutOrderData)
				resp := controller.EosPutOrderResp{
					OrderId: orderId,
				}

				if err != nil {
					log.Errorw("UserSendEOSOrder", "error", err)
					resp.Status = false
				} else {
					resp.Status = true
				}
				bizResp := BizMessage{
					Namespace: user,
					Action:    sendeos,
					Data:      resp,
				}
				_ = wsConnection.sendResp(&bizResp)
			case sendeth: //ETH下单
				go func() {

					var ethPutOrderData controller.EthPutOrderData
					userReqMap := message.Data.(map[string]interface{})
					err = mapstructure.Decode(userReqMap, &ethPutOrderData)
					if err != nil {
						log.Errorw("user get Decode userReqMap", "error", err)
						//goto ERR
						return
					}
					var orderId string

						orderId, err = controller.UserSendEthOrder(&ethPutOrderData)


					resp := controller.EthPutOrderResp{
						OrderId: orderId,
					}

					if err != nil {
						log.Errorw("UserSendETHOrder", "error", err.Error())
						resp.Status = false
					} else {
						resp.Status = true
					}

					bizResp := BizMessage{
						Namespace: user,
						Action:    sendeth,
						Data:      resp,
					}
					_ = wsConnection.sendResp(&bizResp)
				}()

			case sendbde:
				go func() {

					var ethPutOrderData controller.EthPutOrderData
					userReqMap := message.Data.(map[string]interface{})
					err = mapstructure.Decode(userReqMap, &ethPutOrderData)
					if err != nil {
						log.Errorw("user get Decode userReqMap", "error", err)
						//goto ERR
						return
					}
					var orderId string

					orderId, err = controller.UserSendEthOrder(&ethPutOrderData)


					resp := controller.EthPutOrderResp{
						OrderId: orderId,
					}

					if err != nil {
						log.Errorw("UserSendETHOrder", "error", err.Error())
						resp.Status = false
					} else {
						resp.Status = true
					}

					bizResp := BizMessage{
						Namespace: user,
						Action:    sendbde,
						Data:      resp,
					}
					_ = wsConnection.sendResp(&bizResp)
				}()

			case withdraweos:
				var withdrawOrderData controller.WithdrawOrderData
				userReqMap := message.Data.(map[string]interface{})
				err = mapstructure.Decode(userReqMap, &withdrawOrderData)
				if err != nil {
					log.Errorw("user get Decode userReqMap", "error", err)
					goto ERR
				}

				_, err = controller.UserRecallEosOrder(&withdrawOrderData)
				resp := controller.EosWithdrawOrderResp{
					OrderId: withdrawOrderData.OrderId,
				}

				if err != nil {
					log.Errorw("UserRecallEosOrder", "error", err)
					resp.Status = false
				} else {
					resp.Status = true
				}

				bizResp := BizMessage{
					Namespace: user,
					Action:    withdraweos,
					Data:      resp,
				}
				_ = wsConnection.sendResp(&bizResp)
			case withdraweth:
				go func() {

					var withdrawOrderData controller.WithdrawOrderData
					userReqMap := message.Data.(map[string]interface{})
					err = mapstructure.Decode(userReqMap, &withdrawOrderData)
					if err != nil {
						log.Errorw("user get Decode userReqMap", "error", err)
						//goto ERR
						return
					}

					_, err = controller.UserRecallEthOrder(&withdrawOrderData)
					resp := controller.EthWithdrawOrderResp{
						OrderId: withdrawOrderData.OrderId,
					}

					if err != nil {
						log.Errorw("UserRecallEthOrder", "error", err)
						resp.Status = false
					} else {
						resp.Status = true
					}

					bizResp := BizMessage{
						//	Data: "PONG",
						Namespace: user,
						Action:    withdraweth,
						Data:      resp,
					}
					_ = wsConnection.sendResp(&bizResp)
				}()
			case withdrawbde:
				go func() {

					var withdrawOrderData controller.WithdrawOrderData
					userReqMap := message.Data.(map[string]interface{})
					err = mapstructure.Decode(userReqMap, &withdrawOrderData)
					if err != nil {
						log.Errorw("user get Decode userReqMap", "error", err)
						//goto ERR
						return
					}

					_, err = controller.UserRecallEthOrder(&withdrawOrderData)
					resp := controller.EthWithdrawOrderResp{
						OrderId: withdrawOrderData.OrderId,
					}

					if err != nil {
						log.Errorw("UserRecallEthOrder", "error", err)
						resp.Status = false
					} else {
						resp.Status = true
					}

					bizResp := BizMessage{
						//	Data: "PONG",
						Namespace: user,
						Action:    withdrawbde,
						Data:      resp,
					}
					_ = wsConnection.sendResp(&bizResp)
				}()

			case getid:
				//获取当前线程号
				threadId := utils.GetGID()
				threadIdstr := strconv.FormatUint(threadId, 10)
				if len(threadIdstr) >= 7 {
					threadIdstr = threadIdstr[len(threadIdstr)-7:]
				}
				i, err := strconv.Atoi(threadIdstr)
				if err != nil {
					log.Errorw("threadIdstr AtiI", "error", err)
					goto ERR
				}
				threadIdstr = strconv.Itoa(i)
				ids := make([]string, 0)
				for i := 0; i < 3; i++ {
					core.ID++
					idstr := strconv.FormatUint(core.ID, 10)
					idstrLen := len(idstr)
					if idstrLen <= 12 {
						for j := 0; j < 12-idstrLen; j++ {
							idstr = "0" + idstr
						}
					} else {
						panic("ID已用完请维护")
					}
					idstr = threadIdstr + idstr
					ids = append(ids, idstr)
				}
				bizResp := BizMessage{
					//	Data: "PONG",
					Namespace: user,
					Action:    getid,
					Data:      ids,
				}
				_ = wsConnection.sendResp(&bizResp)
			}

		}

	}

ERR:
	// 确保连接关闭
	//if err != nil {
	//	wsConnection.SendMessage(BizMessage{
	//		Namespace: message.Namespace,
	//		Action:    message.Action,
	//		Data:      err.Error(),
	//	})
	//	time.Sleep(time.Second / 2)
	//}
	wsConnection.Close()
	return
}

func (wsConnection *WSConnection) sendResp(bizResp *BizMessage) error {
	if bizResp != nil {
		if err := wsConnection.SendMessage(*bizResp); err != nil {
			if err != errors.New("ERR_SEND_MESSAGE_FULL") {
				// 确保连接关闭
				wsConnection.Close()
				return err
			}
		}
	}
	return nil
}
