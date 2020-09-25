package controllers

import (
	"beeproject/db_mysql"
	"beeproject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type Register struct {
	beego.Controller
}

func (r *Register) Post() {
	var user models.User
	bodydata, err := ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("接受数据失败")
		return
	}
	err = json.Unmarshal(bodydata, &user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("解析数据失败，请重试")
		return
	}
	id, err := db_mysql.InterUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败")
		return
	}
	r.Ctx.WriteString("用户保存成功")
	fmt.Println(id)
	result := models.Responseuser{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
}
