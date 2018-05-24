package routers

import (
	"idhub/saml_idhub_api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
