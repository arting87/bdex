package controllers

import (
	"bdex.in/bdex/bdex-admin-backend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

//用户登录验证
func (c UserController) LoginUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var user loginUserRequest
	var response Response
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info("******* LoginUser ******* ; json.Unmarshal:", user)
	if err != nil {
		response.Data = "json解析失败：" + err.Error()
		response.ErrorCode = JsonError
	} else {
		users, err := models.CheckUser(user.Name, user.PassWord)
		if err != nil {
			response.Data = "用户身份验证失败：" + err.Error()
			response.ErrorCode = NotFoundError
		} else {
			token, err := EncoderSecret(*users)
			if err != nil {
				response.Data = "用户信息编码失败：" + err.Error()
				response.ErrorCode = InternalError
			} else {
				response.Data = LoginResponse{Success: true, Permission: users.Permission, Token: token}
				//c.Ctx.SetCookie(Cookie_token, token)
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

//创建用户
func (c UserController) CreateUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var user CreateUserRequest
	var response Response
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info("******* CreateUser ******* ; json.Unmarshal:", user)
	if err != nil {
		response.Data = "json解析失败：" + err.Error()
		response.ErrorCode = JsonError
	} else {
		err := CheckAdminUserInfo(c.Ctx.Input)
		if err != nil {
			response.Data = "用户身份验证失败：" + err.Error()
			response.ErrorCode = ExternalError
		} else {
			err = models.InsertUser(user.Name, user.PassWord, user.Permission)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = UserResponse{Success: true, Permission: user.Permission}
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

//更新用户信息
func (c UserController) UpdateUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var user models.User
	var response Response

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info("******* UpdateUser ******* ; json.Unmarshal:", user)
	if err != nil {
		response.Data = "json解析失败：" + err.Error()
		response.ErrorCode = JsonError
	} else {
		err := CheckAdminUserInfo(c.Ctx.Input)
		if err != nil {
			response.Data = "用户身份验证失败：" + err.Error()
			response.ErrorCode = ExternalError
		} else {
			err = models.UpdateUser(&user)
			if err != nil {
				response.Data = "数据库操作失败：" + err.Error()
				response.ErrorCode = DbError
			} else {
				response.Data = UserResponse{Success: true, Permission: user.Permission}
			}
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}

//删除用户信息
func (c UserController) DeleteUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var user DeleteUserRequest
	var response Response
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	logs.Info("******* DeleteUser ******* ; json.Unmarshal:", user)
	if err != nil {
		response.Data = "json解析失败：" + err.Error()
		response.ErrorCode = JsonError
	} else {
		err := CheckAdminUserInfo(c.Ctx.Input)
		if err != nil {
			response.Data = "用户身份验证失败：" + err.Error()
			response.ErrorCode = ExternalError
		} else {
			err = models.DeleteUser(user.Id)
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

//获取所有用户信息
func (c UserController) GetAllUser() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var response Response
	err := CheckAdminUserInfo(c.Ctx.Input)
	if err != nil {
		response.Data = "用户身份验证失败" + err.Error()
		response.ErrorCode = ExternalError
	} else {
		users, err := models.FindAllUser()
		if err != nil {
			response.Data = "数据库操作失败：" + err.Error()
			response.ErrorCode = DbError
		} else {
			response.Data = &users
		}
	}
	c.Data["json"] = &response
	c.ServeJSON()
}
