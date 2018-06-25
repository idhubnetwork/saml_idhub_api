package routers

import (
	"idhub/saml_idhub_api/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	beego.Router("/setOrganizations/:id", &controllers.OrganizationController{})
	beego.Router("/getMetadata/:id", &controllers.MetadataController{})
	beego.Router("/getSamlResponse", &controllers.SamlResponseController{})
}
