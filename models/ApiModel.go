package models

//API接口表
type Api struct {
	Id         int    `orm:"column(Id)"`                            //自增主键
	Api        string `orm:"column(Api);default();size(100)"`       //API接口，跟Uid组成唯一键
	Uid        int    `orm:"column(Uid);index;default(0)"`          //用户ID(数据哪个接口组的，目前功能先不做)
	Title      string `orm:"column(Title);default()"`               //接口标题
	Params     string `orm:"column(Params);type(text);default()"`   //接口请求参数
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
