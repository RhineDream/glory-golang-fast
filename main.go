package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "glory-golang/routers"
)

func main() {
	logs.Info("glory-golang启动成功....")
	beego.Run()
}
