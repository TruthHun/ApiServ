package main

import (
	"ApiServ/models"
	_ "ApiServ/routers"

	_ "github.com/go-sql-driver/mysql"

	"encoding/gob"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}

func main() {
	gob.Register(models.User{}) //注册user缓存
	beego.Run()
}
