// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gormtest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/students",
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
		beego.NSNamespace("/stuorders",
			beego.NSInclude(
				&controllers.StuOrderController{},
			),
		),
		beego.NSNamespace("/staffs",
			beego.NSInclude(
				&controllers.StaffController{},
			),
		),
		beego.NSNamespace("/manager",
			beego.NSInclude(
				&controllers.ManagerController{},
			),
		),
		beego.NSNamespace("/cert_cate",
			beego.NSInclude(
				&controllers.Cert_CateController{},
			),
		),
		beego.NSNamespace("/cert_cate_list",
			beego.NSInclude(
				&controllers.Cert_Cate_ListController{},
			),
		),
		beego.NSNamespace("/cert_save",
			beego.NSInclude(
				&controllers.Cert_SaveController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
