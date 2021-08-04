package ethaccountlist

import (
	"bdex.in/bdex/bdex-ex-backend/utils"
	"container/list"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strconv"
	"sync"
	"time"

	"bdex.in/bdex/bdex-ex-backend/log"
)

type EthAccountList struct {
	Leisure list.List
	Pending list.List
	l       sync.RWMutex //single write ,more read  slock
	//记录每个节点指针，用于直接删除
	IDMap map[string]*list.Element // k:orderId  v:*order
}

type Account struct {
	PublicKey  string             `json:"address"`
	PrivateKey string             `json:"private"`
	AuthCall   *bind.CallOpts     `json:"-"`
	Auth       *bind.TransactOpts `json:"-"`
}

func (etl *EthAccountList) InitEAL() {
	etl.IDMap = make(map[string]*list.Element)
}

//向空闲队列中插入Account
func (eal *EthAccountList) InsertLeisureAccount(account Account) {
	//fmt.Println("insert account.PublicKey----:"+account.PublicKey)
	eal.l.Lock()
	defer eal.l.Unlock()
	oElement := eal.Leisure.PushBack(account)
	eal.IDMap[account.PublicKey] = oElement
	//log.Infow("当前可用账户队列", "大小", eal.Leisure.Len())
}

//向Pending队列中插入Account
func (eal *EthAccountList) InsertPendingAccount(account *Account) {
	eal.l.Lock()
	defer eal.l.Unlock()
	oElement := eal.Pending.PushBack(account)
	eal.IDMap[account.PublicKey] = oElement
	//log.Infow("当前可用账户队列", "大小", eal.Leisure.Len())
}

//获取账号空闲队列第一个账号
func (eal *EthAccountList) GetFirstLeisureAccount() (account Account) {
	for {
		if eal.Leisure.Len() == 0 {
			log.Infow("GetOrder", "error", "当前无可用账号")
			time.Sleep(time.Second)
		} else {
			eal.l.Lock()
			defer eal.l.Unlock()
			return eal.Leisure.Front().Value.(Account)
		}
	}
}

//删除买或卖队列第一个订单
func (eal *EthAccountList) DelFirstLeisureAccount() (account Account) {
	eal.l.RLock()
	defer eal.l.RUnlock()
	for {
		if eal.Leisure.Len() == 0 {
			log.Infow("DelFirstLeisureAccount", "error", "addountList is empty-------")
		} else {
			b := eal.Leisure.Front()
			//fmt.Println("b::"+b.Value.(Account).PublicKey)

			eal.Leisure.Remove(b)

			delete(eal.IDMap, b.Value.(Account).PublicKey)
			//fmt.Println("delete account.PublicKey"+b.Value.(Account).PublicKey)
			//fmt.Println(eal.Leisure.Len())
			return b.Value.(Account)
		}
	}
}

func InitLeisureAccounts() (*EthAccountList, error) {
	accountArr := make([]Account, 0)
	//读取JSON文件bdex/bdex-ex-backend/configs/bulk_ether_accounts.json
	//           bdex/bdex-ex-backend/core/ethaccountlist/ethaccountlist.go
	err := utils.LoadJsonFile("./././configs/bulk_ether_accounts.json", &accountArr)
	if err != nil {
		log.Error("加载JSON文件失败", "error", err)
		return nil, err
	}
	//加载公钥/私钥
	log.Infow("加载" + strconv.Itoa(len(accountArr)) + "个账户公钥/私钥中...")
	for k, account := range accountArr {
		//加载账户私钥
		privateKey, err := crypto.HexToECDSA(account.PrivateKey)
		if err != nil {
			log.Errorw("加载权限账户"+account.PublicKey+"的私钥失败：", "error:", err)
			return nil, err
		}
		//构造参与者的签名对象
		funderAuth := bind.NewKeyedTransactor(privateKey)

		funderAuth.GasLimit = 4000000
		funderAuth.GasPrice = big.NewInt(10000000000)

		funderAuthCall := bind.CallOpts{Pending: true, From: common.HexToAddress(account.PublicKey)}

		accountArr[k].Auth = funderAuth
		accountArr[k].AuthCall = &funderAuthCall
	}
	//for _, account := range accountArr {
	//	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
	//	fmt.Println(account)
	//}

	//插入到账号队列中
	accountList, err := InsertAccounts(accountArr)
	if err != nil {
		log.Error("插入空闲账号数组失败", "error", err)
		return nil, err
	}
	return accountList, nil
}

func InsertAccounts(accounts []Account) (*EthAccountList, error) {
	eal := EthAccountList{}
	eal.InitEAL()

	for _, account := range accounts {
		eal.InsertLeisureAccount(account)
	}
	return &eal, nil
}

////通过OrderId获取订单
//func (etl *EthPutList) GetEthPutOrderByOid(oid string) (c *list.Element) {
//	return etl.IDMap[oid]
//}
//
////删除订单通过订单
//func (etl *EthPutList) RemoveOrderByObj(c ethsellorder.EthSellOrder) bool {
//	etl.l.Lock()
//	defer etl.l.Unlock()
//	for {
//		//根据订单号查询出待删除的订单
//		delObj := etl.GetEthPutOrderByOid(c.OrderId)
//		if delObj == nil {
//			return true
//		}
//		if etl.List.Len() == 0 {
//			break
//		} else {
//			o := etl.List.Remove(delObj)
//			delete(etl.IDMap, c.OrderId)
//			log.Infof("删除EthPutSellOrder队列中元素:%s %s:%v", o.(*ethsellorder.EthSellOrder).OrderId)
//			log.Infof("当前EthPutSellOrder队列:%s %s:%d", "大小", etl.List.Len())
//			log.Infof("当前IDMap:%s %s:%s", "内容", etl.IDMap)
//			return true
//		}
//
//	}
//	return true
//}
//
//func (etl *EthPutList) RemoveAll() {
//	var n *list.Element
//	for e := etl.List.Front(); e != nil; e = n {
//		n = e.Next()
//		etl.List.Remove(e)
//	}
//	for k, _ := range etl.IDMap {
//		delete(etl.IDMap, k)
//	}
//
//	log.Infow("buy/sell queue  clean empty-------")
//}
