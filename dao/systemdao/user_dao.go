package systemdao

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"glory-golang/models/system"
	"glory-golang/utils"
)

type UserDao struct {
}

func GetUserList(user *system.SysUser) utils.PageInfo {
	o := orm.NewOrm()

	var list []system.SysUser
	sql := "select * from sys_user where del_flag = 0"
	count, pageSql := GetPageData(user.PageNo, user.PageSize, sql, list)
	_, err := o.Raw(pageSql).QueryRows(&list)
	if err != nil {
		fmt.Println("分页失败")
	}
	return utils.PageInfo{user.PageNo, user.PageSize, count, list}
}

func GetUserById(user *system.SysUser) system.SysUser {
	o := orm.NewOrm()

	_temp := system.SysUser{Id: user.Id}
	//先查询是否存在	Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	err := o.Read(&_temp)
	if err != nil {
		return system.SysUser{}
	}

	return _temp
}

func InsertUser(user *system.SysUser) (flag bool, id string) {
	o := orm.NewOrm()
	//设置id，插入时间
	user.Id = utils.GetUUID()
	user.CreateTime = utils.GetCurrentTime()

	_, err := o.Insert(user)
	if err != nil {
		fmt.Print("插入失败")
		return false, ""
	}
	return true, user.Id
}

func UpdateUser(user *system.SysUser) (flag bool, id string) {
	o := orm.NewOrm()

	_temp := system.SysUser{Id: user.Id}
	//先查询是否存在	Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	err := o.Read(&_temp)
	if err == nil {
		_temp.Username = user.Username
		_temp.Password = user.Password
		_temp.Email = user.Email
		_temp.UpdateTime = utils.GetCurrentTime()

		//Update 默认更新所有的字段，可以更新指定的字段：	// 指定多个字段  o.Update(&user, "Field1", "Field2", ...)
		_, err := o.Update(&_temp)
		if err != nil {
			return false, ""
		}
	}
	return true, user.Id
}

func DeleteUser(user *system.SysUser) bool {
	o := orm.NewOrm()

	res, err := o.Raw("update sys_user set del_flag = 1,update_time = ? where id = ?", utils.GetCurrentTimeString(), user.Id).Exec()
	if err != nil {
		return false
	}
	//返回执行成功条数
	num, _ := res.RowsAffected()
	fmt.Println("mysql row delete nums: ", num)
	return true
}
