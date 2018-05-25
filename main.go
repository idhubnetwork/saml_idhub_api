package main

import (
	_ "idhub/saml_idhub_api/models"
	_ "idhub/saml_idhub_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	beego.Run()
}
