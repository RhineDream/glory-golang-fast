package routers

import (
	"github.com/astaxie/beego"
	"glory-golang/controllers/system"
)

func init() {

	//常用的路由方式，简单直接明了
	beego.Router("/api/user/list", &system.UserController{}, "post:GetUserList")
	beego.Router("/api/user/getById", &system.UserController{}, "post:GetUserById")
	beego.Router("/api/user/save", &system.UserController{}, "post:InsertUser")
	beego.Router("/api/user/updateById", &system.UserController{}, "post:UpdateUser")
	beego.Router("/api/user/delete", &system.UserController{}, "post:DeleteUser")

}
