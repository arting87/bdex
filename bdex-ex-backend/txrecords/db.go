package txrecords

import (
	database "bdex.in/bdex/bdex-ex-backend/db"

	"github.com/spf13/viper"
)

var db *database.DB

func SetupDB() {
	db = database.NewDB(viper.GetString("bdex-ex-backend.dsn"))
	err := db.MapCacher(TransactionRecords{}, nil)

	err = db.Sync2(TransactionRecords{})
	if err != nil {
		panic(err)
	}

	db.ShowSQL(false)
}
