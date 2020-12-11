package routers

import (
	"github.com/astaxie/beego"
	"glory-golang/controllers/system"
)

func init() {
	//常用的路由方式，简单直接明了
	beego.Router("/api/user/list", &system.UserController{}, "post:GetUserList")
}
