package problem

import (
	"github.com/duguying/ojsite/controllers"
)

type ProblemsController struct {
	controllers.BaseController
}

func (this *ProblemsController) Get() {
	this.Data["title"] = this.Lang("title_login")
	this.TplNames = "user/login.tpl"
}