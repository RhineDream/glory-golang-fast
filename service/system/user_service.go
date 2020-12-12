package system

import (
	"glory-golang/models/system"
	result "glory-golang/utils"
)

type UserService interface {
	GetUserList(users *system.SysUser) result.PageInfo

	GetUserById(users *system.SysUser) system.SysUser

	InsertUser(s *system.SysUser) interface{}

	UpdateUser(s *system.SysUser) interface{}

	DeleteUser(s *system.SysUser) interface{}
}
