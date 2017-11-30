package models

//API接口表
type Tags struct {
	Id         int    `orm:"column(Id)"`                   //自增主键
	Uid        int    `orm:"column(Uid);index;default(0)"` //用户ID(数据哪个接口组的，目前功能先不做)
	Tag        string `orm:"column(Tag);default()"`
	TimeCreate int    `orm:"column(TimeCreate);default(0)"` //标签创建时间
}

// 多字段唯一键
func (this *Tags) TableUnique() [][]string {
	return [][]string{
		[]string{"Tag", "Uid"},
	}
}
