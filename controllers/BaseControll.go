package controllers

import (
	"ApiServ/models"

	"strings"
	"time"

	"github.com/TruthHun/gotil/cryptil"
	"github.com/TruthHun/gotil/validatil"
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

//用户注册
func (this *BaseController) Reg() {
	this.handleHasLoginRedirect() //已登录，则跳转接口管理页面
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
					this.Response(0, "目前站点处于邀请注册状态，您的注册邀请码不正确，可联系管理员获取")
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
				this.Response(0, "注册失败，用户名已存在或邮箱已被注册")
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
func (this *BaseController) Login() {
	this.handleHasLoginRedirect() //已登录，则跳转接口管理页面
	if this.Ctx.Input.IsPost() {
		email := this.GetString("Email")
		if err := validatil.ExecValid(email, "email"); err != nil {
			this.Response(1, "邮箱格式不正确")
		}
		var user models.User
		models.O.QueryTable(models.TableUser).Filter("Email", email).Filter("Password", cryptil.Sha1Crypt(this.GetString("Password"))).One(&user)
		if user.Id > 0 {
			this.LoginUser = user
			this.SetSession("LoginUser", user)
			this.Response(1, "登录成功")
		} else {
			this.Response(0, "登录失败，邮箱或密码错误")
		}
	} else {
		this.TplName = "login.html"
	}
}

//处理登录后的跳转
func (this *BaseController) handleHasLoginRedirect() {
	if this.LoginUser.Id > 0 {
		this.Redirect(beego.URLFor("UserController.Apis"), 302)
		this.StopRun()
		return
	}
}

//用户退出
func (this *BaseController) Logout() {
	this.SetSession("LoginUser", models.User{}) //设置空值
	this.Redirect(beego.URLFor("IndexController.Index"), 302)
}

//处理未登录的跳转
func (this *BaseController) handleNotLoginRedirect() {
	if this.LoginUser.Id == 0 {
		this.Redirect(beego.URLFor("BaseController.Login"), 302)
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
