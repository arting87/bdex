package routers

import (
	"bdex.in/bdex/bdex-admin-backend/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	//用户请求
	beego.Router("/api/admin/login", &controllers.UserController{}, "post:LoginUser")
	beego.Router("/api/admin/user/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/api/admin/user/update", &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/api/admin/user/delete", &controllers.UserController{}, "post:DeleteUser")
	beego.Router("api/admin/user/getUser", &controllers.UserController{}, "get:GetAllUser")

	//交易对
	beego.Router("/api/admin/getToken", &controllers.TokenController{}, "post:GetToken")
	beego.Router("/api/admin/addToken", &controllers.TokenController{}, "post:AddToken")
	beego.Router("/api/admin/deleteToken", &controllers.TokenController{}, "post:DeleteToken")

	//订单
	beego.Router("/api/admin/getAllDeleteOrder", &controllers.OrderController{}, "post:GetAllOrder")
	beego.Router("/api/admin/dbDeleteOrder", &controllers.OrderController{}, "post:DeleteOrder")

	//交易状态
	beego.Router("/api/admin/getDataInDay", &controllers.ExchangeController{}, "post:ExchangeStatus")

	//新闻
	beego.Router("/api/admin/getNews", &controllers.NewsController{}, "get:GetExchangeNews")
	beego.Router("/api/admin/addNews", &controllers.NewsController{}, "post:AddExchangeNews")
	beego.Router("/api/admin/deleteNews", &controllers.NewsController{}, "post:DeleteExchangeNews")

	//auth 验证
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowOrigins: []string{"222.67.190.136"}, //cookie
		AllowAllOrigins:  true, //auth
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type, Content-Length", "Authorization", "Token", "Accept", "X-Requested-With", "yourHeaderFeild"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

}
