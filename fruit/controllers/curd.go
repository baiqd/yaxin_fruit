package controllers

import (
	"fmt"
	"fruit/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CurdController struct {
	beego.Controller
}

type PostPackagesRequest struct {
	Name  string `json:"name"`
	Unit  string `json:"unit"`
	Price string `json:"price"`
}
//应该写在配置文件中
var url = "http://10.160.0.89:8089/v1/fruit/"

// @router / [get]
func (c *CurdController) GetFruitALL() {
	var inparm PostPackagesRequest

	ob, err := models.GetFruitAll()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = ob
	}

	c.ServeJSON()
	fmt.Println(url)
	noitce_to_service("get", url, inparm)
}

// @router /:name [get]
func (c *CurdController) Get() {
	var inparm PostPackagesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &inparm)
	name := c.Ctx.Input.Param(":name")
	if name != "" {
		ob, err := models.GetFruitByName(name)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			c.Data["json"] = ob
		}
	}

	c.ServeJSON()

	noitce_to_service("get", url+name, inparm)
}

// @router / [post]
func (c *CurdController) Post() {
	var inparm PostPackagesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &inparm)

	var f = new(models.FruitInfo)
	f.Name = inparm.Name
	f.Unit = inparm.Unit
	f.Price = inparm.Price

	_, err := models.AddFruit(f)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = f
	}

	c.ServeJSON()
	noitce_to_service("post", url, inparm)
}

// @router /:name [delete]
func (c *CurdController) Delete() {
	var inparm PostPackagesRequest
	name := c.Ctx.Input.Param(":name")
	if name != "" {
		err := models.DeleteFruit(name)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			noitce_to_service("delete", url+name, inparm)
		}
	}
	c.ServeJSON()
}

// @router /:name [put]
func (c *CurdController) Put() {
	var inparm PostPackagesRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &inparm)
	name := c.Ctx.Input.Param(":name")
	if name != "" {
		f, err := models.GetFruitByName(name)
		if err != nil {
			c.Data["json"] = err.Error()
		} else {
			if f != nil{
				f.Name = inparm.Name
				f.Unit = inparm.Unit
				f.Price = inparm.Price
				err := models.UpdateFruit(f)
				if err != nil {
					c.Data["json"] = err.Error()
				} else {
					c.Data["json"] = f
				}
			}
			
		}
	}
	c.ServeJSON()
	noitce_to_service("put", url+name, inparm)
}

func noitce_to_service(method string, url string, para PostPackagesRequest) (err error) {
	
	if method == "" || url == "" {
		return
	}

	body, err := json.Marshal(para)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch method {
	case "get":
		req := httplib.Get(url)
		str, err := req.Response()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	case "post":
		req := httplib.Post(url)
		if para.Name == "" || para.Unit == "" || para.Price == "" {
			return
		}
		req.Body(body)
		str, err := req.Response()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	case "put":
		req := httplib.Put(url)
		if para.Name == "" || para.Unit == "" || para.Price == "" {
			return
		}
		req.Body(body)
		str, err := req.Response()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	case "delete":
		req := httplib.Delete(url)
		str, err := req.Response()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(str)
		}
	default:
		panic(fmt.Sprintf("method fail"))
	}

	return
}
