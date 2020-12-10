package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"glory-golang/models"
)

type StudentController struct {
	beego.Controller
}

func (c *StudentController) List() {
	c.Ctx.WriteString("1")
}

func (c *StudentController) GetById() {
	//json方式提交
	var form models.Student
	//接受requestBody
	data := c.Ctx.Input.RequestBody
	//json数据封装到LoginForm对象中
	var json = jsoniter.ConfigCompatibleWithStandardLibrary //号称全宇宙最快的json解析包
	err := json.Unmarshal(data, &form)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	fmt.Println(form)

	logs.Info("this %s cat is %v years old", "yellow", 3)
	logs.Info("ParseLoginForm:", &form)

	//c.Data["json"] = result.SuccessResult(&form)
	//c.ServeJSON()
}
