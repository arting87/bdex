package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ExchangeController struct {
	beego.Controller
}

func (c ExchangeController) ExchangeStatus() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var status ExchangeStatusRequest
	var response Response
	err := CheckUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &status)
		logs.Info("******* ExchangeStatus ******* ; json.Unmarshal:", status)
		if err != nil {
			response.Data = err.Error()
			response.ErrorCode = JsonError
		} else {
			res, err := models.FindManyCount(status.Day)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = *res
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}
