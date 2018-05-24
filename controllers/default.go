package controllers

import (
	"idhub/saml_idhub_api/keypairs"
	"idhub/saml_idhub_api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var privateKeyDirPrefix = "./privateKey"
var certificateDirPrefix = "./certificate"

type OrganizationController struct {
	beego.Controller
}

type MetadataController struct {
	beego.Controller
}

type SamlResponseController struct {
	beego.Controller
}

func (c *OrganizationController) Get() {
	id := c.Ctx.Input.Param(":id")
	privateKeyDir := privateKeyDirPrefix + id + ".pem"
	certificateDir := certificateDirPrefix + id + ".crt"
	s, err := keypairs.GenKeyAndCert(privateKeyDir, certificateDir)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)

	privateKey, err := ioutil.ReadFile(privateKeyDir)
	if err != nil {
		return
	}
	fmt.Println(privateKey)

	certificate, err := ioutil.ReadFile(certificateDir)
	if err != nil {
		return
	}
	fmt.Println(certificate)

	var organization models.Organization
	organization.Key = id
	organization.PrivateKey = privateKey
	organization.Certificate = certificate

	o := orm.NewOrm()
	id, err := o.Insert(&organization)
	if err != nil {
		// 查询数据库失败
	}
}

func (c *MetadataController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *SamlResponseController) Post() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
