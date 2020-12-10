package routers

import (
	"github.com/astaxie/beego"
	"glory-golang/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//常用的路由方式，简单直接明了
	beego.Router("/api/student/list", &controllers.StudentController{}, "get:List")
	beego.Router("/api/student/getById", &controllers.StudentController{}, "post:GetById")
}
