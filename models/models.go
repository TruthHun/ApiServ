package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var TablePrefix = beego.AppConfig.DefaultString("db_prefix", "hc_")
var O orm.Ormer

//model
var (
	ModelApi  = new(Api)
	ModelUser = new(User)
)

//table
var (
	TableUser = TablePrefix + "user"
	TableApi  = TablePrefix + "api"
)

func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=%v",
		beego.AppConfig.String("db_user"),
		beego.AppConfig.String("db_password"),
		beego.AppConfig.String("db_host"),
		beego.AppConfig.String("db_port"),
		beego.AppConfig.String("db_database"),
		beego.AppConfig.String("db_charset"),
	))
	if beego.AppConfig.String("runmode") == "dev" {
		//开发模式下，打印SQL语句
		orm.Debug = true
	}
	orm.RegisterModelWithPrefix(
		TablePrefix, //表前缀
		ModelUser,   //用户表
		ModelApi,    //API表
	)
	if err := orm.RunSyncdb("default", false, true); err != nil {
		beego.Error(err)
	}
	O = orm.NewOrm()
}
