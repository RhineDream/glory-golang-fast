package system

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type SysUser struct {
	Id         string    `orm:"column(id);pk" description:"主键id"`
	LoginName  string    `orm:"column(login_name);size(100);null" description:"登录账号"`
	Username   string    `orm:"column(username);size(100);null" description:"真实姓名"`
	Password   string    `orm:"column(password);size(255);null" description:"密码"`
	Salt       string    `orm:"column(salt);size(45);null" description:"md5密码盐"`
	Photo      string    `orm:"column(photo);size(255);null" description:"头像"`
	Birthday   time.Time `orm:"column(birthday);type(datetime);null" description:"生日"`
	Sex        string    `orm:"column(sex);size(1);null" description:"性别(0-默认未知,1-男,2-女)"`
	Email      string    `orm:"column(email);size(45);null" description:"电子邮件"`
	Phone      string    `orm:"column(phone);size(45);null" description:"电话"`
	OrgCode    string    `orm:"column(org_code);size(64);null" description:"机构编码"`
	Status     string    `orm:"column(status);size(1);null" description:"用户状态(1-正常,2-禁用)"`
	DelFlag    string    `orm:"column(del_flag);size(1);null" description:"删除状态(0-正常,1-已删除)"`
	EmpNo      string    `orm:"column(emp_no);size(100);null" description:"工号，唯一键"`
	Telephone  string    `orm:"column(telephone);size(45);null" description:"座机号"`
	UserType   string    `orm:"column(user_type);size(10);null" description:"用户类型-1老师2家长3园长"`
	CreateBy   string    `orm:"column(create_by);size(32);null" description:"创建人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateBy   string    `orm:"column(update_by);size(32);null" description:"更新人"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	Remarks    string    `orm:"column(remarks);size(255);null"`
	DeleteTime time.Time `orm:"column(delete_time);type(datetime);null" description:"删除时间"`
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(SysUser))
}
