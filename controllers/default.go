package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "http://www.baidu.com"
	c.Data["Email"] = "qhw1987752122@163.com"
	c.TplName = "index.tpl"
}
