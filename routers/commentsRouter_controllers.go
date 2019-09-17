package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_CateController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_Cate_ListController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"] = append(beego.GlobalControllerRouter["gormtest/controllers:Cert_SaveController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gormtest/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gormtest/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gormtest/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gormtest/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:ManagerController"] = append(beego.GlobalControllerRouter["gormtest/controllers:ManagerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StaffController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StaffController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StaffController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StaffController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StaffController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StaffController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StaffController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StaffController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StaffController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StaffController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:wid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StuOrderController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:wid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StudentController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StudentController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StudentController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StudentController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetById",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gormtest/controllers:StudentController"] = append(beego.GlobalControllerRouter["gormtest/controllers:StudentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
