package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["fruit/controllers:CurdController"] = append(beego.GlobalControllerRouter["fruit/controllers:CurdController"],
		beego.ControllerComments{
			"GetFruitALL",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["fruit/controllers:CurdController"] = append(beego.GlobalControllerRouter["fruit/controllers:CurdController"],
		beego.ControllerComments{
			"Get",
			`/:name`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["fruit/controllers:CurdController"] = append(beego.GlobalControllerRouter["fruit/controllers:CurdController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["fruit/controllers:CurdController"] = append(beego.GlobalControllerRouter["fruit/controllers:CurdController"],
		beego.ControllerComments{
			"Delete",
			`/:name`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["fruit/controllers:CurdController"] = append(beego.GlobalControllerRouter["fruit/controllers:CurdController"],
		beego.ControllerComments{
			"Put",
			`/:name`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["fruit/controllers:TotalController"] = append(beego.GlobalControllerRouter["fruit/controllers:TotalController"],
		beego.ControllerComments{
			"Get",
			`/total/:name`,
			[]string{"get"},
			nil})

}
