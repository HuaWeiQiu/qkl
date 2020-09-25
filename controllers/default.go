package controllers

import (
	"beeproject/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	username := c.Ctx.Input.Query("username")
	pwd := c.Ctx.Input.Query("pwd")
	//c.GetString("user")
	//c.GetString("pwd")
	if username == "admin" && pwd == "12138" {
		c.Ctx.ResponseWriter.Write([]byte("登陆成功"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("登陆失败"))
	c.Data["Website"] = "http://www.baidu.com"
	c.Data["Email"] = "qhw1987752122@163.com"
	c.TplName = "index.tpl"
}

//func (c *MainController) Post(){
//name := c.Ctx.Request.FormValue("name")
//age := c.Ctx.Request.FormValue("age")
//sex := c.Ctx.Request.FormValue("sex")
//fmt.Println(name,age,sex)
//if name != "admin" && age != "21" && sex != "21"{
//	c.Ctx.WriteString("校验失败")
//	return
//}
//c.Ctx.WriteString("校验成功")
//}
/*func (c *MainController) Post(){
var person models.Person
dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
if err != nil {
	c.Ctx.WriteString("数据接收失败")
	return
}
err = json.Unmarshal(dataBytes,&person)
if err != nil{
	c.Ctx.WriteString("数据解析失败，请重试")
	return
}
c.Ctx.WriteString("请求成功")
fmt.Println("名字",person.Name)
fmt.Println("年龄",person.Age)
fmt.Println("性别",person.Sex)
*/
func (c *MainController) Post() {
	var register models.Register
	dataBytes, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	err = json.Unmarshal(dataBytes, &register)
	if err != nil {
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	c.Ctx.WriteString("请求成功")
	fmt.Println("姓名:", register.Name)
	fmt.Println("生日:", register.Birthday)
	fmt.Println("地址:", register.Address)
	fmt.Println("用户名:", register.Nick)
}
