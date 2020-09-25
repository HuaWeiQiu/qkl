package main

import (
	_ "beeproject/db_mysql"
	_ "beeproject/routers"
	"github.com/astaxie/beego"
)

func main() {
	/*config := beego.AppConfig
	appname := config.String("appname")
	fmt.Println("应用名称:",appname)
	port ,err := config.Int("httpport")
	if err != nil{
		panic("项目配置文件解析失败，请检查配置文件")
	}
	fmt.Println(port)
	beego.Run(":8090")

	*/
	beego.Run(":8090")
}
