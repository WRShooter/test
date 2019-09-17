package controllers

import (
	"encoding/json"
	"fmt"
	"gormtest/models"

	"github.com/astaxie/beego"
)

type Cert_Cate_ListController struct {
	beego.Controller
}

// @Title 获得所有证书
// @Description 返回所有的证书数据
// @Success 200 {object} models.Cert_cate_list
// @router / [get]
func (u *Cert_Cate_ListController) GetAll() {
	ss := models.GetAllCert_Cate_Lists()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}

// @Title 创建职工
// @Description 创建职工
// @Param      body          body   models.Cert_cate_list true          "body for user content"
// @Success 200 create success
// @Failure 403 body is empty
// @router / [post]
func (u *Cert_Cate_ListController) Post() {
	var s models.Cert_cate_list
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	id := models.AddCert_Cate_List(&s)
	u.Data["json"] = id
	fmt.Println(id)
	u.ServeJSON()
}

// @Title 获得一个证书
// @Description 返回某证书数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Cert_cate_list
// @router /:id [get]
func (u *Cert_Cate_ListController) GetById() {
	id, _ := u.GetInt64(":id")
	fmt.Println(id)
	s := models.GetCert_Cate_ListById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 修改证书
// @Description 修改证书的内容
// @Param      body          body   models.Cert_cate_list true          "body for user content"
// @Success 200 {int} models.Cert_cate_list
// @Failure 403 body is empty
// @router / [put]
func (u *Cert_Cate_ListController) Update() {
	var s models.Cert_cate_list
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateCert_Cate_List(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个证书
// @Description 删除某证书数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Cert_cate_list
// @router /:id [delete]
func (u *Cert_Cate_ListController) Delete() {
	id, _ := u.GetInt64(":id")
	models.DeleteCert_Cate_List(id)
	u.Data["json"] = true
	u.ServeJSON()
}
