package main

import (
	"ApiServ/models"
	_ "ApiServ/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}
