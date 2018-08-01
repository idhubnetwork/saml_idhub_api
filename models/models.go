package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type Organization struct {
	Id          int
	Key         string
	PrivateKey  string `orm:"type(text)"`
	Certificate string `orm:"type(text)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root(127.0.0.1:3306)/saml_idhub_api?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(Organization))

	// create table
	orm.RunSyncdb("default", false, true)
}
