package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type OrderController struct {
	beego.Controller
}

func (c OrderController) GetAllOrder() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var orderRequest models.GetOrderRequest
	var response Response
	err := CheckUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderRequest)
		logs.Info("******* GetAllOrder ******* ; json.Unmarshal:", orderRequest)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			orderResponse, err := models.FindOrder(orderRequest)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = *orderResponse
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

func (c OrderController) DeleteOrder() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var deleteOrder DeleteOrderRequest
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败:" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &deleteOrder)
		logs.Info("******* DeleteOrder ******* ; json.Unmarshal:", deleteOrder)
		if err != nil {
			response.Data = "json解析失败：" + err.Error()
			response.ErrorCode = JsonError
		} else {
			if deleteOrder.OrderId == "" {
				response.Data = "订单id不能为空"
				response.ErrorCode = ExternalError
			} else {
				err := models.DeleteOrderById(deleteOrder.OrderId)
				if err != nil {
					response.Data = "数据库操作失败：" + err.Error()
					response.ErrorCode = InternalError
				} else {
					response.Data = NormalResponse{Success: true}
				}
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}
