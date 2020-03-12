package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"monitortools/controllers"
	"monitortools/models"
	_ "monitortools/routers"
)

func init ()  {
	models.RegisterDb()
}



func main() {
	orm.Debug=true
	beego.Debug()

	orm.RunSyncdb("default",false, true)

	beego.Router("/",&controllers.HomeController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Router("/addhosts",&controllers.AddHostsController{})

	beego.Run()

}
