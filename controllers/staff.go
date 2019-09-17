package controllers

import (
	"encoding/json"
	"fmt"
	"gormtest/models"

	"github.com/astaxie/beego"
)

type StaffController struct {
	beego.Controller
}

// @Title 获得所有职工
// @Description 返回所有的职工数据
// @Success 200 {object} models.Staff
// @router / [get]
func (u *StaffController) GetAll() {
	ss := models.GetAllStaffs()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}

// @Title 创建职工
// @Description 创建职工
// @Param      body          body   models.Staff true          "body for user content"
// @Success 200 create success
// @Failure 403 body is empty
// @router / [post]
func (u *StaffController) Post() {
	var s models.Staff
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddStaff(&s)
	u.Data["json"] = uid
	fmt.Println(uid)
	u.ServeJSON()
}

// @Title 获得一个职工
// @Description 返回某职工数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Staff
// @router /:id [get]
func (u *StaffController) GetById() {
	id, _ := u.GetInt64(":id")
	fmt.Println(id)
	s := models.GetStaffById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 修改职工
// @Description 修改职工的内容
// @Param      body          body   models.Staff true          "body for user content"
// @Success 200 {int} models.Staff
// @Failure 403 body is empty
// @router / [put]
func (u *StaffController) Update() {
	var s models.Staff
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateStaff(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个职工
// @Description 删除某职工数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Staff
// @router /:id [delete]
func (u *StaffController) Delete() {
	id, _ := u.GetInt64(":id")
	models.DeleteStaff(id)
	u.Data["json"] = true
	u.ServeJSON()
}
