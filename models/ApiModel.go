package models

import (
	"fmt"

	"strings"

	"errors"

	"github.com/astaxie/beego/orm"
)

//API接口表
type Api struct {
	Id         int    `orm:"column(Id)"`                            //自增主键
	Api        string `orm:"column(Api);default();size(100)"`       //API接口，跟Uid组成唯一键
	Uid        int    `orm:"column(Uid);index;default(0)"`          //用户ID(数据哪个接口组的，目前功能先不做)
	Title      string `orm:"column(Title);default()"`               //接口标题
	Params     string `orm:"column(Params);type(text);default()"`   //接口请求参数
	Resp       string `orm:"column(Resp);type(text);default()"`     //响应说明
	JsonSucc   string `orm:"column(JsonSucc);type(text);default()"` //接口响应的成功的json数据
	JsonErr    string `orm:"column(JsonErr);type(text);default()"`  //接口响应的失败的json数据
	Methods    string `orm:"column(Methods);default()"`             //允许的请求方法，GET、POST、PUT、DELETE、HEAD、OPTIONS，多个方法之间用英文逗号分隔
	TimeCreate int    `orm:"column(TimeCreate);default(0)"`         //接口创建时间
}

// 多字段唯一键
func (this *Api) TableUnique() [][]string {
	return [][]string{
		[]string{"Api", "Uid"},
	}
}

//返回响应的json数据
//@param		username		用户名
//@param		api				请求的api
//@param		method			请求方法
//@param		succ			是否返回成功的数据，默认为true
//@return		js				json数据字符串
func (this *Api) GetToResponse(username string, api string, method string, succ ...bool) (js string, err error) {
	retsucc := true
	sql := ""
	api = "/" + strings.Trim(api, " /") //清除空格和斜线
	var params []orm.Params
	if len(succ) > 0 {
		retsucc = succ[0]
	}
	if retsucc {
		sql = fmt.Sprintf("select a.JsonSucc,a.Methods from %v a left join %v u on u.Id=a.Uid where u.Username=? and a.Api=? limit 1", TableApi, TableUser)
	} else {
		sql = fmt.Sprintf("select a.JsonErr,a.Methods from %v a left join %v u on u.Id=a.Uid where u.Username=? and a.Api=? limit 1", TableApi, TableUser)
	}
	O.Raw(sql, username, api).Values(&params)
	if len(params) > 0 {
		if !strings.Contains(params[0]["Methods"].(string), method) {
			return "", errors.New(fmt.Sprintf("请求方法（%v）不符合接口要求，当前接口允许的请求方法为：%v", method, params[0]["Methods"].(string)))
		}
		if retsucc {
			return params[0]["JsonSucc"].(string), nil
		} else {
			return params[0]["JsonErr"].(string), nil
		}
	}
	return
}
