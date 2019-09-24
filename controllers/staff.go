package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStaffs(c *gin.Context) {
	staffs, err := models.GetAllStaffs()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, staffs)
	}
}

func GetStaffByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	stuorders, err := models.GetStaffById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, stuorders)
	}
}

func CreateStaff(c *gin.Context) {
	var staff *models.Staff
	c.BindJSON(&staff)
	if models.ExistsStaffById(int64(staff.Id)) == false {
		err := models.AddStaff(staff)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, staff)
		}
	}
}

func UpdateStaff(c *gin.Context) {
	pwd := c.Query("pwd")
	position := c.Query("postion")
	phone := c.Query("phone")
	name := c.Query("name")
	idcard := c.Query("idcard")
	dept := c.Query("dept")
	cap := c.Query("cap")
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	staffs, _ := models.GetStaffById(id)
	staff := &models.Staff{Pwd: staffs[0].Pwd, Position: staffs[0].Position, Phone: staffs[0].Phone, Name: staffs[0].Name, Idcard: staffs[0].Idcard, Dept: staffs[0].Dept, Cap: staffs[0].Cap}
	staff.Id = int(id)
	c.ShouldBind(&staff)
	if pwd != "" {
		staff.Pwd = pwd
	}
	if position != "" {
		staff.Position = position
	}
	if phone != "" {
		staff.Phone = phone
	}
	if name != "" {
		staff.Name = name
	}
	if idcard != "" {
		staff.Idcard = idcard
	}
	if dept != "" {
		staff.Dept = dept
	}
	if cap != "" {
		staff.Cap = cap
	}
	err = staff.UpdateStaff()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, staff)
}

func DeleteStaff(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	if models.ExistsStaffById(id) == true {

		err := models.DeleteStaff(id)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"id #" + Id: "deleted"})
	}
}
