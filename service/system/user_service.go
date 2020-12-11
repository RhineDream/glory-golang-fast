package system

import (
	"glory-golang/models/system"
)

type UserService interface {
	GetUserList(users *system.SysUser) []system.SysUser
}
