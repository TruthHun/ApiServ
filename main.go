package main

import (
	"ApiServ/models"
	_ "ApiServ/routers"

	_ "github.com/go-sql-driver/mysql"

	"encoding/gob"

	"strings"

	"encoding/json"

	"fmt"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	SetTplFunc()
}

func main() {
	gob.Register(models.User{}) //注册user缓存
	beego.Run()
}

type ParamsStruct struct {
	ParamsName  []string `json:"ParamsName"`
	ParamsState []string `json:"ParamsState"`
	ParamsType  []string `json:"ParamsType"`
}

//模板函数直接写在这里了
func SetTplFunc() {
	//分割字符串成数组
	beego.AddFuncMap("Split", func(str interface{}, seg string) []string {
		return strings.Split(fmt.Sprintf("%v", str), seg)
	})

	//分割字符串成数组
	beego.AddFuncMap("TrimLeft", func(str, sutset string) string {
		return strings.TrimLeft(str, sutset)
	})
	//{"ParamsName":["username","password"],"ParamsState":["用户名","密码"],"ParamsType":["string","string"]}
	//操作参数
	beego.AddFuncMap("HandleParams", func(paramsjson string) []map[string]string {
		var (
			params ParamsStruct
			data   []map[string]string
		)
		json.Unmarshal([]byte(paramsjson), &params)
		l1 := len(params.ParamsName)
		l2 := len(params.ParamsState)
		l3 := len(params.ParamsType)
		if l1 == l2 && l1 == l3 {
			for i := 0; i < l1; i++ {
				pn := params.ParamsName[i]
				ps := params.ParamsState[i]
				pt := params.ParamsType[i]
				if len(pn) > 0 && len(ps) > 0 && len(pt) > 0 {
					data = append(data, map[string]string{
						"ParamsName":  params.ParamsName[i],
						"ParamsState": params.ParamsState[i],
						"ParamsType":  params.ParamsType[i],
					})
				}
			}
		}
		return data
	})

	//分割字符串成数组
	beego.AddFuncMap("InMethods", func(method string, methods []string) bool {
		for _, v := range methods {
			if v == method {
				return true
			}
		}
		return false
	})
}
