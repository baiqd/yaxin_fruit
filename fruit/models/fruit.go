package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	FruitInfos map[string]*FruitInfo
)

type FruitInfo struct {
	Id         int64     `orm:"pk;unique;auto"`
	Name       string    `orm:"size(64)"`
	Unit       string    `orm:"size(64)"`
	Price      string    `orm:"size(64)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTIme time.Time `orm:"auto_now;type(datetime)"`
}

func createTable() {
   	name := "default"
  	force := false
   	verbose := true

   	err := orm.RunSyncdb(name, force, verbose)
   	if err != nil {
       	beego.Error(err)
   	}
}

func dbInit() {
	appcfg := beego.AppConfig
	dbUser := appcfg.String("DB_user")
	dbPasswd := appcfg.String("DB_passwd")
	dbIp := appcfg.String("DB_ip")
	dbPort := appcfg.String("DB_port")
	dbName := appcfg.String("DB_name")

	var param string = dbUser + ":" + dbPasswd + "@tcp(" + dbIp + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia%2FShanghai"
	fmt.Println("DB Info : " + param)

	orm.RegisterDriver("mysql", orm.DRMySQL)

    // set default database
    orm.RegisterDataBase("default", "mysql", "root:123456@/default?charset=utf8", 30)
    
    // register model
    orm.RegisterModel(new(FruitInfo))

	orm.SetMaxOpenConns("default", 100)

	createTable()
}

func init() {
	dbInit()
}

func AddFruit(f *FruitInfo) (id int64, err error) {

	o := orm.NewOrm()
	o.Using("default")

	o.Begin()

	id, err = o.Insert(f)

	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}

	return
}

func DeleteFruit(name string) (err error) {
	
	o := orm.NewOrm()

	v := FruitInfo{Name: name}

	if err = o.Read(&v, "Name"); err == nil {
		var num int64
		if num, err = o.Delete(&FruitInfo{Id: v.Id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}

	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}

	return
}

func UpdateFruit(f *FruitInfo) (err error) {
	
	o := orm.NewOrm()

	v := FruitInfo{Id: f.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(f); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}

	if err != nil {
		o.Rollback()
	} else {
		o.Commit()
	}

	return
}

func GetFruitByName(name string) (v *FruitInfo, err error) {	
	o := orm.NewOrm()

	v = &FruitInfo{Name: name}
	err = o.Read(v, "Name")

	if err != nil {
		v = nil
		o.Rollback()
	} else {
		o.Commit()
	}

	return
}

func GetFruitAll() (v []*FruitInfo, err error) {
	
	o := orm.NewOrm()
	o.Using("default")
	
	_, e := o.QueryTable("fruit_info").All(&v)

	if  e == nil {
		return v,nil
	}

	return nil,err
}
