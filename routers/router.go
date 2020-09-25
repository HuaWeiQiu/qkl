package routers

import (
	"beeproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/register", &controllers.Register{})
}
