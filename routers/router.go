package routers

import (
	"beegoTest/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/students", &controllers.StudentsController{})
}


