package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Student))
	orm.RegisterModel(new(Problem))
	orm.RegisterModel(new(OAuth))
}
