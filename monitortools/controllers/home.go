package controllers

import (
	"github.com/astaxie/beego"
	"monitortools/models"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
	this.Data["IsLogin"]=checkAccount(this.Ctx)

	hostStatus1, err := models.GetHostStatus()
	if err != nil {
		beego.Error()
	} else {
		this.Data["HostStatus"] = hostStatus1
	}

}
