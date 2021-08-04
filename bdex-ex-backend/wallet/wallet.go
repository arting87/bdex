package wallet

import "time"

type EosWallet struct {
	UserName   string    `xorm:"user_name" json:"user_name"`
	UserId     string    `xorm:"user_id" json:"user_id"`
	OwnerKey   string    `xorm:"owner_key" json:"owner_key"`
	ActiveKey  string    `xorm:"active_key" json:"active_key"`
	CreateTime time.Time `xorm:"create_time" json:"create_time"`
	UpdateTime time.Time `xorm:"update_time" json:"update_time"`
}

type EthWallet struct {
	Address    string    `xorm:"address" json:"address"`
	UserId     string    `xorm:"user_id" json:"user_id"`
	CreateTime time.Time `xorm:"create_time" json:"create_time"`
	UpdateTime time.Time `xorm:"update_time" json:"update_time"`
}
