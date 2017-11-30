package controllers

import (
	"strings"

	"ApiServ/models"

	"time"

	"github.com/TruthHun/gotil/cryptil"
	"github.com/TruthHun/gotil/validatil"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

//用户注册
func (this *UserController) Reg() {
	this.handleHasLoginRedirect()
	if this.Ctx.Input.IsPost() {
		var rules = map[string][]string{
			"Username":   []string{"required", "mincount:1", "maxcount:30", "alphanumeric"}, //必填、1-30个字符、数字和字母
			"Nickname":   []string{"required", "mincount:1", "maxcount:20"},                 //必填、1-30个字符、数字和字母
			"Email":      []string{"required", "email", "maxcount:100"},                     //必填、1-100个字符
			"Password":   []string{"required"},                                              //必填
			"Repassword": []string{"required"},                                              //必填
			"InviteCode": []string{},                                                        //非必填
		}
		if params, errs := validatil.Valid(this.Ctx.Request.Form, rules); len(errs) > 0 {
			this.Response(0, "请按表单要求提示填写注册表单", errs)
		} else {
			if InviteCode := strings.TrimSpace(beego.AppConfig.String("invite_code")); len(InviteCode) > 0 {
				if ic, ok := params["InviteCode"]; !ok || ic != InviteCode {
					this.Response(0, "注册邀请码不正确，请联系管理员获取")
				}
			}
			if params["Password"].(string) != params["Repassword"] {
				this.Response(0, "两次输入密码不一致，请重新输入")
			}
			var user = models.User{
				Username:   params["Username"].(string),
				Nickname:   params["Nickname"].(string),
				Email:      params["Email"].(string),
				Password:   cryptil.Sha1Crypt(params["Password"].(string)), //密码
				TimeCreate: int(time.Now().Unix()),                         //注册时间
			}
			if created, id, err := models.O.ReadOrCreate(&user, "Username"); err != nil || !created || id == 0 {
				this.Response(0, "注册失败，用户名已存在")
			} else {
				//设置session
				user.Id = int(id)
				this.SetSession("LoginUser", user)
				this.LoginUser = user
				this.Response(1, "注册成功")
			}

		}
	} else {
		this.TplName = "reg.html"
	}
}

//用户登录
func (this *UserController) Login() {
	this.handleHasLoginRedirect()
	if this.Ctx.Input.IsPost() {

	} else {
		this.TplName = "login.html"
	}
}

//用户退出
func (this *UserController) Logout() {
	this.Redirect(beego.URLFor("IndexController.Index"), 302)
}

//接口列表
func (this *UserController) Apis() {
	this.TplName = "apis.html"
}

//接口创建、更新、编辑
func (this *UserController) ApiCreate() {
	this.TplName = "api_edit.html"
}

//删除接口
func (this *UserController) ApiDel() {

}
