package controllers

import (
	"github.com/astaxie/beego/logs"
	"glory-golang/models"
	"glory-golang/service"
	result "glory-golang/utils"
)

type StudentController struct {
	BaseController
}

func (c *StudentController) List() {
	c.Ctx.WriteString("1")
}

func (c *StudentController) GetById() {
	//json方式提交
	var form models.Student
	c.GetRequestBodyJson(&form)

	logs.Info("ParseLoginForm:", &form)
	stu := service.GetStudentById(form)
	stu2 := service.GetTeacherById(form)
	logs.Info("ParseLoginForm:", &stu2)
	c.SendResponse(result.OkResult(&stu))
}
