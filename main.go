package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ics042/go-blog/controllers"
	"github.com/ics042/go-blog/models"
	_ "github.com/ics042/go-blog/routers"
)

func init() {
	models.InitDb()
	controllers.RegisterTemplateFuncs()
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
