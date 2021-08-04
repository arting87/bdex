package main

import (
	_ "bdex.in/bdex/bdex-admin-backend/models"
	_ "bdex.in/bdex/bdex-admin-backend/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
