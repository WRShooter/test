package controllers

import (
	"encoding/json"
	"fmt"
	"gormtest/models"

	"github.com/astaxie/beego"
)

type StuOrderController struct {
	beego.Controller
}

// @Title 获得所有证书
// @Description 返回所有的证书数据
// @Success 200 {object} models.Stuorder
// @router / [get]
func (u *StuOrderController) GetAll() {
	ss := models.GetAllStuoders()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}

// @Title 获得一个学生的学生证书
// @Description 返回某学生学生证书数据
// @Param      wid            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Stuorder
// @router /:wid [get]
func (u *StuOrderController) GetById() {
	id, _ := u.GetInt64(":wid")
	fmt.Println(id)
	s := models.GetStuodersByWId(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 创建学生证书
// @Description 创建学生证书
// @Param      body          body   models.Stuorder true          "body for user content"
// @Success 200 create success
// @Failure 403 body is empty
// @router / [post]
func (u *StuOrderController) Insert() {
	var s models.Stuorder
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	id := models.AddStuoders(&s)
	u.Data["json"] = id
	fmt.Println(id)
	u.ServeJSON()
}

// @Title 修改学生证书
// @Description 修改学生证书的内容
// @Param      body          body   models.Stuorder true          "body for user content"
// @Success 200 {int} models.StuOrder
// @Failure 403 body is empty
// @router / [put]
func (u *StuOrderController) Update() {
	var s models.Stuorder
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateStuoders(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个学生证书
// @Description 删除某学生数据
// @Param      wid            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Stuorder
// @router /:wid [delete]
func (u *StuOrderController) Delete() {
	id, _ := u.GetInt64(":wid")
	models.DeleteStuoders(id)
	u.Data["json"] = true
	u.ServeJSON()
}
