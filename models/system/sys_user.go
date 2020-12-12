package system

import (
	"github.com/astaxie/beego/orm"
	"glory-golang/common"
	"time"
)

type SysUser struct {
	Id         string    `orm:"column(id);pk" description:"主键id" json:"id"`
	LoginName  string    `orm:"column(login_name);size(100);null" description:"登录账号" json:"login_name"`
	Username   string    `orm:"column(username);size(100);null" description:"真实姓名" json:"username"`
	Password   string    `orm:"column(password);size(255);null" description:"密码" json:"password"`
	Salt       string    `orm:"column(salt);size(45);null" description:"md5密码盐" json:"salt"`
	Photo      string    `orm:"column(photo);size(255);null" description:"头像" json:"photo"`
	Birthday   time.Time `orm:"column(birthday);type(datetime);null" description:"生日" json:"birthday"`
	Sex        string    `orm:"column(sex);size(1);null" description:"性别(0-默认未知,1-男,2-女)" json:"sex"`
	Email      string    `orm:"column(email);size(45);null" description:"电子邮件"  json:"email"`
	Phone      string    `orm:"column(phone);size(45);null" description:"电话"  json:"phone"`
	OrgCode    string    `orm:"column(org_code);size(64);null" description:"机构编码" json:"org_code"`
	Status     string    `orm:"column(status);size(1);null" description:"用户状态(1-正常,2-禁用)" json:"status"`
	DelFlag    string    `orm:"column(del_flag);size(1);null" description:"删除状态" json:"del_flag"`
	EmpNo      string    `orm:"column(emp_no);size(100);null" description:"工号，唯一键" json:"emp_no"`
	Telephone  string    `orm:"column(telephone);size(45);null" description:"座机号" json:"telephone"`
	UserType   string    `orm:"column(user_type);size(10);null" description:"用户类型" json:"user_type"`
	CreateBy   string    `orm:"column(create_by);size(32);null" description:"创建人" json:"create_by"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间" json:"create_time"`
	UpdateBy   string    `orm:"column(update_by);size(32);null" description:"更新人" json:"update_by"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间" json:"update_time"`
	Remarks    string    `orm:"column(remarks);size(255);null" json:"remarks"`
	DeleteTime time.Time `orm:"column(delete_time);type(datetime);null" description:"删除时间" json:"delete_time"`
	common.DataEntity
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(SysUser))
}
