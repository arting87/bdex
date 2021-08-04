package wallet

import "time"

type User struct {
	UserName   string    `xorm:"'user_name' index" json:"user_name"`
	UserID     string    `xorm:"user_id pk" json:"user_id"`
	Phone      int16     `xorm:"phone" json:"phone"`
	EosAccount string    `xorm:"eos_account" json:"eos_account"`
	EthAccount string    `xorm:"eth_account" json:"eth_account"`
	EosToken   string    `xorm:"eos_token" json:"eos_token"`
	EthToken   string    `xorm:"eth_token" json:"eth_token"`
	CrossToken string    `xorm:"cross_token" json:"cross_token"`
	CreateTime time.Time `xorm:"create_time" json:"create_time"`
	UpdateTime time.Time `xorm:"update_time" json:"update_time"`
}
