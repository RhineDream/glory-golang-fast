package impl

import (
	system2 "glory-golang/dao/system"
	"glory-golang/models/system"
)

type UserServiceImpl struct {
}

func (d *UserServiceImpl) GetUserList(user *system.SysUser) []system.SysUser {
	users := []system.SysUser{}
	//sysUser1 := system.SysUser{Id: result.GetUUID(),Username:"zhangsansss"}

	sysUser, err := system2.GetSysUserById(user.Id)
	if err != nil {

	}
	users = append(users, *sysUser)
	return users
}
