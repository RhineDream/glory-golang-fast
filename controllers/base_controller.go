package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

/**
统一返回
*/
func (baseController *BaseController) SendResponse(data interface{}) {
	baseController.Data["json"] = data
	baseController.ServeJSON()
}
