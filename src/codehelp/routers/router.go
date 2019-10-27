package routers

import (
	"codehelp/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.SetStaticPath("a", "/home/siwei/github/go_angular/src/codehelp/views")

	beego.Router("/test", &controllers.MainController{})
}
