package system

import (
	"glory-golang/common"
	"glory-golang/models/business"
	"glory-golang/models/system"
	system2 "glory-golang/service/system"
	"glory-golang/service/system/impl"
	result "glory-golang/utils"
)

type UserController struct {
	common.BaseController
}

/**
 * 列表分页查询
 */
func (c *UserController) GetUserList() {
	var form system.SysUser
	c.GetRequestBodyJson(&form)

	// 实例化接口实现类
	f := new(impl.UserServiceImpl)
	// 声明一个UserService的接口
	var service system2.UserService
	// 将接口赋值f，也就是*file类型
	service = f
	// 使用GetUserListService接口进行数据写入
	listService := service.GetUserList(&form)

	c.SendResponse(result.OkResult(listService))
}

func (c *UserController) GetById() {
	//json方式提交
	var form business.Student
	c.GetRequestBodyJson(&form)

	//sysUser := system.SysUser{Id: result.GetUUID(),Username:"zhangsansss"}

	//_, err := system.AddSysUser(&sysUser)
	//if(err != nil){
	//	logs.Info("异常")
	//}
	//
	//logs.Info("ParseLoginForm:", &form)
	//stu := business2.GetStudentById(form)
	//stu2 := business2.GetTeacherById(form)
	//logs.Info("ParseLoginForm:", &stu2)
	//c.SendResponse(result.OkResult(&stu))
}
