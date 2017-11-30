package models

//用户表
type User struct {
	Id         int    `orm:"column(Id)"`                                 //自增主键
	Username   string `orm:"column(Username);unique;default();size(30)"` //用户名
	Avatar     string `orm:"column(Avatar);default();size(100)"`         //用户头像
	Nickname   string `orm:"column(Nickname);default();size(20)"`        //用户名
	Password   string `orm:"column(Password);default();size(40)"`        //登录密码
	Email      string `orm:"column(Email);default();unique;size(100)"`   //邮箱
	TimeCreate int    `orm:"column(TimeCreate);default(0);size(50)"`     //注册时间
}
