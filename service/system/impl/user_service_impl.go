package impl

import (
	"glory-golang/dao/systemdao"
	"glory-golang/models/system"
	result "glory-golang/utils"
)

type UserServiceImpl struct {
}

func (d *UserServiceImpl) GetUserList(user *system.SysUser) result.PageInfo {
	list := systemdao.GetUserList(user)
	return list
}

func (d *UserServiceImpl) GetUserById(user *system.SysUser) system.SysUser {
	result := systemdao.GetUserById(user)
	return result
}

func (d *UserServiceImpl) InsertUser(user *system.SysUser) interface{} {
	flag, id := systemdao.InsertUser(user)
	if flag {
		return result.OkResult(id)
	}
	return result.Fail(500, "保存失败")
}

func (d *UserServiceImpl) UpdateUser(user *system.SysUser) interface{} {
	flag, id := systemdao.UpdateUser(user)
	if flag {
		return result.OkResult(id)
	}
	return result.Fail(500, "修改失败")
}

func (d *UserServiceImpl) DeleteUser(user *system.SysUser) interface{} {
	flag := systemdao.DeleteUser(user)
	if flag {
		return result.Ok()
	}
	return result.Fail(500, "修改失败")
}
