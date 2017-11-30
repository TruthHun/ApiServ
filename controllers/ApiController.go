package controllers

import (
	"fmt"

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
	//TODO 是否带preview参数，如果带preview参数，则表示web访问接口

	//设置允许跨域
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	//用户名
	user := this.GetString(":user")
	//请求方法:GET、POST、PUT、DELETE、HEAD、OPTIONS
	method := this.Ctx.Request.Method
	//是否请求返回成功的json数据，否则返回失败的json数据
	opt, _ := this.GetBool("RetSucc", true)
	//请求API
	api := this.GetString(":splat")

	fmt.Println(user, method, api, opt)
}
