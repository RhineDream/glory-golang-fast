package utils

import (
	"github.com/xormplus/xorm"
)

// 创建 xorm 客户端
func CreateClient() *xorm.Engine {

	var engine *xorm.Engine
	var err error
	engine, err = xorm.NewEngine("mysql", "root:ZwhzMysql2019@tcp(60.205.202.248:3306)/glory-golang?charset=utf8")
	if err != nil {
		println("open error:", &err)
	}
	//engine.SetMapper(core.samemapper{})      //表示struct的类的名称和数据库中相同
	engine.ShowSQL(true) //显示sql语句
	//engine.Logger().SetLevel(core.LOG_DEBUG)

	return engine
}
