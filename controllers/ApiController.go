package controllers

import (
	"ApiServ/models"

	"github.com/astaxie/beego"
)

var methods = map[string]string{
	"GET":     "GET",
	"POST":    "POST",
	"PUT":     "PUT",
	"DELETE":  "DELETE",
	"HEAD":    "HEAD",
	"OPTIONS": "OPTIONS",
}

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Api() {

	//设置允许跨域
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	//请求API
	api := this.GetString(":splat")

	//用户名
	user := this.GetString(":user")

	//请求方法:GET、POST、PUT、DELETE、HEAD、OPTIONS
	method := this.Ctx.Request.Method

	//是否请求返回成功的json数据，否则返回失败的json数据
	opt, _ := this.GetBool("as-succ", true)

	if js, err := models.ModelApi.GetToResponse(user, api, method, opt); err == nil {
		this.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
		this.Ctx.WriteString(js)
	} else {
		this.Data["json"] = map[string]interface{}{"ok": false, "msg": err.Error()}
		this.ServeJSON()
	}
	return
}
