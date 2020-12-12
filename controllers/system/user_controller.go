package system

import (
	"glory-golang/common"
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

	//  声明一个UserService的接口 = 实例化接口实现类
	var service system2.UserService = new(impl.UserServiceImpl)

	// 使用GetUserListService接口进行数据写入
	pageInfo := service.GetUserList(&form)

	c.SendResponse(result.OkResult(pageInfo))
}

/**
 * 根据id查询
 */
func (c *UserController) GetUserById() {
	//json方式提交
	var form system.SysUser
	c.GetRequestBodyJson(&form)

	//  声明一个UserService的接口 = 实例化接口实现类
	var service system2.UserService = new(impl.UserServiceImpl)

	// 使用GetUserListService接口进行数据写入
	results := service.GetUserById(&form)

	c.SendResponse(result.OkResult(results))
}

/**
 * 插入
 */
func (c *UserController) InsertUser() {
	//json方式提交
	var form system.SysUser
	c.GetRequestBodyJson(&form)

	//  声明一个UserService的接口 = 实例化接口实现类
	var service system2.UserService = new(impl.UserServiceImpl)

	// 使用GetUserListService接口进行数据写入
	res := service.InsertUser(&form)

	c.SendResponse(res)
}

/**
 * 插入
 */
func (c *UserController) UpdateUser() {
	//json方式提交
	var form system.SysUser
	c.GetRequestBodyJson(&form)

	//  声明一个UserService的接口 = 实例化接口实现类
	var service system2.UserService = new(impl.UserServiceImpl)

	// 使用GetUserListService接口进行数据写入
	res := service.UpdateUser(&form)

	c.SendResponse(res)
}

/**
 * 删除
 */
func (c *UserController) DeleteUser() {
	//json方式提交
	var form system.SysUser
	c.GetRequestBodyJson(&form)

	//  声明一个UserService的接口 = 实例化接口实现类
	var service system2.UserService = new(impl.UserServiceImpl)

	// 使用GetUserListService接口进行数据写入
	res := service.DeleteUser(&form)

	c.SendResponse(res)
}
