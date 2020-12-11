package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	jsoniter "github.com/json-iterator/go"
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

/**
获取请求体json数据
*/

func (baseController *BaseController) GetRequestBodyJson(obj interface{}) {
	//接受requestBody
	data := baseController.Ctx.Input.RequestBody
	//json数据封装到LoginForm对象中
	var json = jsoniter.ConfigCompatibleWithStandardLibrary //号称全宇宙最快的json解析包
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

}
