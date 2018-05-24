package routers

import (
	"idhub/saml_idhub_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/setOrganizations/:id", &controllers.OrganizationController{})
	beego.Router("/getMetadata/:id", &controllers.MetadataController{})
	beego.Router("/getSamlResponse", &controllers.SamlResponseController{})
}
