package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/controllers/chainConn"
	"bdex.in/bdex/bdex-admin-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type TokenController struct {
	beego.Controller
}

func (c TokenController) AddToken() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var addTokenRequest AddTokenRequest
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &addTokenRequest)
		logs.Info("******* AddToken ******* ; json.Unmarshal:", addTokenRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			if addTokenRequest.MarketCode == MarketCode_EOS {
				err := chainConn.AddEosToken(addTokenRequest.TokenContract, addTokenRequest.Symbol, uint64(addTokenRequest.Decimal))
				if err != nil {
					response.Data = "添加代币上链失败" + err.Error()
					response.ErrorCode = ConnectError
				} else {
					response.Data = NormalResponse{Success: true}
				}
			} else {
				err := chainConn.AddEthToken(addTokenRequest.Symbol, addTokenRequest.TokenContract, addTokenRequest.Decimal, addTokenRequest.MarketCode)
				if err != nil {
					response.Data = "添加代币上链失败" + err.Error()
					response.ErrorCode = ConnectError
				} else {
					response.Data = NormalResponse{Success: true}
				}
			}

		}
	}

	c.Data["json"] = &response
	c.ServeJSON()
}

func (c TokenController) DeleteToken() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var deleteTokenRequest DeleteTokenRequest
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &deleteTokenRequest)
		logs.Info("******* DeleteToken ******* ; json.Unmarshal:", deleteTokenRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {

			if deleteTokenRequest.MarketCode == MarketCode_EOS {
				err := chainConn.DeleteEosToken(deleteTokenRequest.TokenContract)
				if err != nil {
					response.Data = "链上删除代币失败：" + err.Error()
					response.ErrorCode = ConnectError
				} else {
					response.Data = NormalResponse{Success: true}
				}
			} else {
				err := chainConn.DeleteEthToken(deleteTokenRequest.TokenContract, deleteTokenRequest.MarketCode)
				if err != nil {
					response.Data = "链上删除代币失败：" + err.Error()
					response.ErrorCode = ConnectError
				} else {
					response.Data = NormalResponse{Success: true}
				}

			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

func (c TokenController) GetToken() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var tokenRequest models.GetTokenRequest
	var response Response
	err := CheckUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &tokenRequest)
		logs.Info("******* GetToken ******* ; json.Unmarshal:", tokenRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			tokens, err := models.FindTokenInfo(tokenRequest)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = *tokens
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}
