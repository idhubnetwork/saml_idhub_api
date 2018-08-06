package models

import (
    "fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"github.com/astaxie/beego/config"
)

// Model Struct
type Organization struct {
	Id          int
	Key         string
	PrivateKey  string `orm:"type(text)"`
	Certificate string `orm:"type(text)"`
}

func init() {

	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
    }

    mysql_link := iniconf.String("mysql_user")+":"+iniconf.String("mysql_pass")+"@"+"("+
    iniconf.String("mysql_host")+":"+iniconf.String("mysql_port")+")/saml_idhub_api?charset=utf8"

	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", mysql_link, 30)

	// register model
	orm.RegisterModel(new(Organization))

	// create table
	orm.RunSyncdb("default", false, true)
}
