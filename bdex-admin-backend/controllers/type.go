package controllers

import (
	"time"
)

const (
	Cookie_token     = "token"
	AdminPermission  = 1
	NormalPermission = 2
	ExpireTime       = 3600 //1 hour
)

const (
	InternalError = 1 //内部错误
	ExternalError = 2 //外部错误
	ConnectError  = 3 //连接错误
	DbError       = 4 //数据库错误
	JsonError     = 5 //json错误
	NotFoundError = 6 //未找到错误
)

const (
	MarketCode_ETH    = 2
	MarketCode_EOS    = 1
	MarketCode_BDEX   = 3
	MarketCode_ACROSS = 4
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

type UserResponse struct {
	Success    bool `json:"success"`
	Permission int  `json:"permission"`
}

type LoginResponse struct {
	Success    bool   `json:"success"`
	Permission int    `json:"permission"`
	Token      string `json:"token"`
}

type NormalResponse struct {
	Success bool `json:"success"`
}

type AddTokenRequest struct {
	Symbol        string `json:"symbol"`
	TokenContract string `json:"token_contract"`
	Decimal       int    `json:"decimal"`
	MarketCode    int    `json:"market_code"`
}

type DeleteTokenRequest struct {
	TokenContract string `json:"token_contract"`
	MarketCode    int    `json:"market_code"`
}

type DeleteOrderRequest struct {
	OrderId string `json:"order_id"`
}

type ExchangeStatusRequest struct {
	Day time.Time `json:"day"`
}

type AddNewsRequest struct {
	Title    string `json:"title"`
	TypeCode int    `json:"type_code"`
	Content  string `json:"content"`
}

type loginUserRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"pass_word"`
}

type DeleteNewsRequest struct {
	Id int `json:"id"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type CreateUserRequest struct {
	Name       string `json:"name"`
	PassWord   string `json:"pass_word"`
	Permission int    `json:"permission"`
}

type UserInfoSafeCheck struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Pwd        string `json:"pwd"`
	Permission int    `json:"permission"`
	TimeStamp  int64  `json:"time_stamp"`
}
