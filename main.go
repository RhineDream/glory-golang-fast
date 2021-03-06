package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "glory-golang/routers"
)

func main() {
	logs.Info("glory-golang启动成功....")
	beego.Run()
}

func init() {
	//配置日志模块
	l := logs.GetLogger()
	l.Println("this is a message of http")
	logs.GetLogger("ORM").Println("this is a message of orm")

	//注册数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("username")+":"+beego.AppConfig.String("password")+"@tcp("+beego.AppConfig.String("address")+":"+beego.AppConfig.String("port")+")/"+beego.AppConfig.String("dbname")+"?charset=utf8&parseTime=True&loc=Local", 30, 30)
	orm.Debug = true

	orm.SetMaxIdleConns("default", 20)
	orm.SetMaxOpenConns("default", 100)

	////访问接口前验证token
	//var filterUser = func(ctx *context.Context) {
	//	url := ctx.Request.RequestURI
	//	flag := result.CheckStringInList(url, *anonList)
	//	if flag {
	//		logs.Info("-----当前请求是匿名请求，无需拦截----")
	//	} else {
	//		token := ctx.Input.Header("Authorization")
	//		if token == "" {
	//			logs.Info("token为空！")
	//			result.ResponseWithJson(ctx.ResponseWriter, 400, result.FailedResult("token为空"))
	//			return
	//		}
	//		appid, err := user_encode.Token_auth(token, "secret")
	//		if err != nil {
	//			//http.Error(ctx.ResponseWriter, "Token verification not pass", http.StatusBadRequest)
	//			result.ResponseWithJson(ctx.ResponseWriter, 400, result.FailedResult("token无效"))
	//			return
	//		}
	//		logs.Info("appid---------", appid)
	//
	//		logs.Info("Request token---------", token)
	//	}
	//
	//}
	//
	//beego.InsertFilter("/*", beego.BeforeRouter, filterUser)
}
