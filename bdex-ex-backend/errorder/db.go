package errorder

import (
	database "bdex.in/bdex/bdex-ex-backend/db"
	"time"

	"github.com/spf13/viper"
)

var db *database.DB

func SetupDB() {
	db = database.NewDB(viper.GetString("bdex-ex-backend.dsn"))
	err := db.MapCacher(ErrorOrder{}, nil)

	err = db.Sync2(ErrorOrder{})
	if err != nil {
		panic(err)
	}

	db.ShowSQL(false)
}

type ErrorOrder struct {
	OrderId           string    `xorm:"'order_id' varchar(255) pk" json:"i"`
	Name              string    `xorm:"'name' varchar(255)"`
	NameChainCode     string    `xorm:"'name_chain_code' varchar(10)"`
	TradeType         string    `xorm:"'trade_type' varchar(8)" json:"a"` //1 is sell, 0 is buy
	MarketType        int       `xorm:"'market_type' int" `
	MatchSymbol       string    `xorm:"'match_symbol' varchar(8)"` //1AAA   2DDD
	DeleVol           float64   `xorm:"'dele_vol' Double"`         //除EOS或ETH以外的数量
	SellTokenSymbol   string    `xorm:"'sell_token_symbol' varchar(20)"`
	SellTokenAmount   string    `xorm:"'sell_token_amount' varchar(20)"`
	SellTokenContract string    `xorm:"'sell_token_contract' varchar(255)"`
	BuyTokenSymbol    string    `xorm:"'buy_token_symbol' varchar(20)" json:"bs"`
	BuyTokenAmount    string    `xorm:"'buy_token_amount' varchar(20)"`
	BuyTokenContract  string    `xorm:"'buy_token_contract' varchar(255)" json:"bc"`
	Price             string    `xorm:"'price' varchar(255)" json:"p"`
	Fee               string    `xorm:"'fee' varchar(20)"`
	WithdrawAble      bool      `xorm:"'withdraw_able' bool"`
	ExSuccess         bool      `xorm:"'ex_success' bool"`
	TargetAddress     string    `xorm:"'target_address' varchar(255)" json:"ta"`
	TargetChainCode   string    `xorm:"'target_chain_code' varchar(10)"`
	CreateAt          time.Time `xorm:"'create_at' index"`
	UpdateAt          time.Time `xorm:"'update_at'"`
	TxHash            string    `xorm:"'tx_hash' varchar(255)"`
	Status            int       `xorm:"'status' int"`
	Err               string    `xorm:"'err' Text"`
}
