package routers

import (
	"tax/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/after", &controllers.MainController{}, "post:After")
}
