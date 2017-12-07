package controllers

import (
	"ApiServ/models"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

//未登录用户，直接跳转登录页面
//用户登录了，才能访问下面的控制器方法
func (this *UserController) Prepare() {
	this.BaseController.Prepare() //继承基类初始化函数
	this.handleNotLoginRedirect() //没登录用户跳转登录页面
}

//接口列表
func (this *UserController) Apis() {
	var apis []models.Api
	ApiDomain := beego.AppConfig.DefaultString("api_domain", fmt.Sprintf("http://localhost:%v", beego.AppConfig.String("httpport")))
	models.O.QueryTable(models.TableApi).Filter("Uid", this.LoginUser.Id).OrderBy("-Id").All(&apis)
	this.TplName = "apis.html"
	this.Data["Apis"] = apis
	this.Data["ApiDomain"] = strings.TrimRight(ApiDomain, "/ ")
}

//接口创建、更新、编辑
func (this *UserController) ApiCreate() {
	//注意：这里的参数id，在query请求中，"id"首字母是小写的，在form表单中，"id"是首字母大写的
	var id int
	if this.Ctx.Input.IsPost() { //更新或创建
		var form = this.Ctx.Request.Form
		var apidata models.Api
		var params = make(map[string]interface{})
		var resp = make(map[string]interface{})
		id, _ = this.GetInt("Id") //接口id
		if apidata.Title = strings.TrimSpace(form.Get("Title")); apidata.Title == "" {
			this.Response(0, "接口名称不能为空")
		}
		if methods, ok := form["Methods"]; ok && len(methods) > 0 {
			apidata.Methods = strings.Join(methods, ",")
		} else {
			this.Response(0, "请选择您的接口请求方法")
		}
		if apidata.Api = strings.TrimSpace(form.Get("Api")); apidata.Api == "" {
			this.Response(0, "请填写您的接口API")
		} else {
			apidata.Api = "/" + strings.Trim(apidata.Api, " /") //去除斜线
		}
		apidata.JsonSucc = this.Ctx.Request.Form.Get("JsonSucc")
		apidata.JsonErr = this.Ctx.Request.Form.Get("JsonErr")
		if val, ok := form["ParamsName"]; ok {
			params["ParamsName"] = val
		}
		if val, ok := form["ParamsType"]; ok {
			params["ParamsType"] = val
		}
		if val, ok := form["ParamsState"]; ok {
			params["ParamsState"] = val
		}
		if b, err := json.Marshal(params); err == nil {
			apidata.Params = string(b)
		}

		if val, ok := form["ResponseField"]; ok {
			resp["ResponseField"] = val
		}
		if val, ok := form["ResponseState"]; ok {
			resp["ResponseState"] = val
		}
		if b, err := json.Marshal(resp); err == nil {
			apidata.Resp = string(b)
		}

		apidata.Uid = this.LoginUser.Id
		if id > 0 { //更新
			apidata.Id = id
			if affetd, _ := models.O.Update(&apidata); affetd > 0 {
				this.Response(1, "更新成功")
			} else {
				if apidata.Id > 0 {
					this.Response(0, "更新失败，您没有对内容进行更改")
				} else {
					this.Response(0, "更新失败，您要更换的API接口已存在")
				}
			}
		} else { //判断接口是否已存在
			if created, _, err := models.O.ReadOrCreate(&apidata, "Api", "Uid"); created && err == nil {
				this.Response(1, "新增接口成功")
			} else {
				this.Response(0, "新增接口失败，您要创建的API接口已存在")
			}
		}
	} else {
		idstr := this.Ctx.Request.URL.Query().Get("id")
		id, _ = strconv.Atoi(idstr)
		this.Data["HasData"] = false
		if id > 0 {
			var api models.Api
			models.O.QueryTable(models.TableApi).Filter("Id", id).Filter("Uid", this.LoginUser.Id).One(&api)
			this.Data["Api"] = api
			if api.Id > 0 {
				this.Data["HasData"] = true
			}
		}
		this.TplName = "api_edit.html"
	}

}

//删除接口
func (this *UserController) ApiDel() {
	id, _ := this.GetInt("id")
	if id > 0 {
		if affected, _ := models.O.QueryTable(models.TableApi).Filter("Uid", this.LoginUser.Id).Filter("Id", id).Delete(); affected > 0 {
			this.Response(1, "删除成功")
		}
	}
	this.Response(0, "删除失败，您要删除的接口已不存在")
}
