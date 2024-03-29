package main

import (
	_ "fruit/docs"
	_ "fruit/routers"
	_ "fruit/models"
	
	"github.com/astaxie/beego"
)



func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
