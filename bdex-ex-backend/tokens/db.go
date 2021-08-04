package tokens

import (
	database "bdex.in/bdex/bdex-ex-backend/db"

	"github.com/spf13/viper"
)

var db *database.DB

func SetupDB() {
	db = database.NewDB(viper.GetString("bdex-ex-backend.dsn"))
	err := db.MapCacher(Tokens{}, nil)

	err = db.Sync2(Tokens{})
	if err != nil {
		panic(err)
	}

	db.ShowSQL(false)
}

type Tokens struct {
	ID            int    `xorm:"'id' int pk autoincr unique"`
	TokenID       uint64 `xorm:"'token_id' BigInt index"`
	Symbol        string `xorm:"'symbol' varchar(255)"`
	TokenContract string `xorm:"'token_contract' varchar(42)"`
	MarketCode    int    `xorm:"'market_code' int"`
	Decimal       uint   `xorm:"'decimal' int"`
	OrderDecimal  uint   `xorm:"'order_decimal' int"`
	IsCancel      bool   `xorm:"'iscancel' bool"`
}
