package controllers

import (
	"github.com/astaxie/beego"
	"monitortools/models"
)

type AddHostsController struct {
	beego.Controller
}

func (this *AddHostsController) Get() {
	this.Data["IsAddHost"] = true

	if checkAccount(this.Ctx) {
	} else {
		this.Redirect("/", 302)
		return
	}
	this.TplName = "addHosts.html"
}

func (this *AddHostsController) Post() {

	hostIp := this.Input().Get("hostip")
	hostPort:=this.Input().Get("hostport")
	if len(hostIp) != 0 {
		err := models.AddHosts(hostIp,hostPort)
		if err != nil {
			beego.Error(err)
		}
	}

	this.Redirect("/", 302)
	return
}
