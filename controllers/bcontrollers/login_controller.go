package bcontrollers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Post() {
	username := this.Input().Get("uname")
	password := this.Input().Get("pwd")
	if beego.AppConfig.String("uname") == username &&
		beego.AppConfig.String("pwd") == password {
		this.SetSession("loginKey", beego.AppConfig.String("key"))
		this.Redirect("/admin", 302)
	} else {
		this.Redirect("/login", 302)
	}

}

func (this *LoginController) Get() {
	action := this.Input().Get("action")
	switch action {
	case "logout":
		this.Logout()
	}
	this.Data["PageTitle"] = "Blog"
	this.TplName = "admin/layout/login.html"

}

func (this *LoginController) Logout() {
	this.DelSession("isLogin")
}
