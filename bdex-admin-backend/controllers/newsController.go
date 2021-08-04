package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type NewsController struct {
	beego.Controller
}

func (c NewsController) GetExchangeNews() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var response Response
	err := CheckUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		news, err := models.FindAllExchangeNews()
		if err != nil {
			response.Data = "数据库操作失败：" + err.Error()
			response.ErrorCode = DbError
		} else {
			response.Data = &news
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

func (c NewsController) AddExchangeNews() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var newsRequest AddNewsRequest
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败：" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &newsRequest)
		logs.Info("******* AddExchangeNews ******* ; json.Unmarshal:", newsRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			err := models.AddExchangeNews(newsRequest.Title, newsRequest.TypeCode, newsRequest.Content)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = NormalResponse{Success: true}
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

func (c NewsController) DeleteExchangeNews() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var newsRequest DeleteNewsRequest
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败：" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &newsRequest)
		logs.Info("******* DeleteExchangeNews ******* ; json.Unmarshal:", newsRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			err := models.DeleteExchangeNews(int64(newsRequest.Id))
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = NormalResponse{Success: true}
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}
