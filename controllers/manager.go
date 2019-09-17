package controllers

import (
	"encoding/json"
	"fmt"
	"gormtest/models"

	"github.com/astaxie/beego"
)

type ManagerController struct {
	beego.Controller
}

// @Title 获得所有管理员
// @Description 返回所有的管理员数据
// @Success 200 {object} models.Manager
// @router / [get]
func (u *ManagerController) GetAll() {
	ss := models.GetAllManagers()
	u.Data["json"] = ss
	fmt.Println(ss)
	u.ServeJSON()
}

// @Title 创建管理员
// @Description 创建管理员
// @Param      body          body   models.Manager true          "body for user content"
// @Success 200 create success
// @Failure 403 body is empty
// @router / [post]
func (u *ManagerController) Post() {
	var s models.Manager
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	uid := models.AddManager(&s)
	u.Data["json"] = uid
	fmt.Println(uid)
	u.ServeJSON()
}

// @Title 获得一个管理员
// @Description 返回某管理员数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Manager
// @router /:id [get]
func (u *ManagerController) GetById() {
	id, _ := u.GetInt64(":id")
	fmt.Println(id)
	s := models.GetManagerById(id)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 修改管理员
// @Description 修改管理员的内容
// @Param      body          body   models.Manager true          "body for user content"
// @Success 200 {int} models.Manager
// @Failure 403 body is empty
// @router / [put]
func (u *ManagerController) Update() {
	var s models.Manager
	json.Unmarshal(u.Ctx.Input.RequestBody, &s)
	models.UpdateManager(&s)
	u.Data["json"] = s
	u.ServeJSON()
}

// @Title 删除一个管理员
// @Description 删除某管理员数据
// @Param      id            path   int    true          "The key for staticblock"
// @Success 200 {object} models.Manager
// @router /:id [delete]
func (u *ManagerController) Delete() {
	id, _ := u.GetInt64(":id")
	models.DeleteManager(id)
	u.Data["json"] = true
	u.ServeJSON()
}
