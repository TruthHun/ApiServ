package controllers

import (
	"ApiServ/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	LoginUser models.User //登录的用户
}

func (this *BaseController) Prepare() {

	//1、从session中获取用户的登录信息
	if user := this.GetSession("LoginUser"); user != nil {
		this.LoginUser = user.(models.User)
	}

	//2、分配信息到模板变量
	this.Data["LoginUser"] = this.LoginUser
}

//处理登录后的跳转
func (this *BaseController) handleHasLoginRedirect() {
	if this.LoginUser.Id > 0 {
		this.Redirect(beego.URLFor("UserController.Apis"), 302)
		this.StopRun()
		return
	}
}

//处理未登录的跳转
func (this *BaseController) handleNotLoginRedirect() {
	if this.LoginUser.Id == 0 {
		this.Redirect(beego.URLFor("UserController.Login"), 302)
		this.StopRun()
		return
	}
}

//响应
func (this *BaseController) Response(status int, msg string, data ...interface{}) {
	var resp = map[string]interface{}{
		"Status": status,
		"Msg":    msg,
	}
	if len(data) > 0 {
		resp["Data"] = data[0]
	}
	this.Data["json"] = resp
	this.ServeJSON()
	this.StopRun()
}
