package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExist := this.Input().Get("exit") == "true"
	if isExist {
		this.Ctx.SetCookie("uname", "", 0, "/")
		this.Ctx.SetCookie("pwd", "", 0, "/")
		this.Redirect("/", 302)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	maxAge := 1<<31 - 1
	if uname == beego.AppConfig.String("uname") && pwd == beego.AppConfig.String("pwd") {
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	this.Redirect("/", 302)
	return

}

func checkAccount(ctx *context.Context) bool {
	ck,err := ctx.Request.Cookie("uname")
	if err !=nil {
		return false
	}
	uname := ck.Value

	ck,err = ctx.Request.Cookie("pwd")
	if err !=nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd
}
