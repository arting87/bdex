package ethsellorder

import (
	database "bdex.in/bdex/bdex-ex-backend/db"

	"github.com/spf13/viper"
)

var db *database.DB

func SetupDB() {
	db = database.NewDB(viper.GetString("bdex-ex-backend.dsn"))
	err := db.MapCacher(EthSellOrder{}, nil)

	err = db.Sync2(EthSellOrder{})
	if err != nil {
		panic(err)
	}

	db.ShowSQL(false)
}
