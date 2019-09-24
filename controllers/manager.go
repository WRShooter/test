package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllManagers(c *gin.Context) {
	managers, err := models.GetAllManagers()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, managers)
	}
}

func GetManagerByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	manager, err := models.GetManagerById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, manager)
	}
}

func CreateManager(c *gin.Context) {
	var manager *models.Manager
	c.BindJSON(&manager)
	err := models.AddManager(manager)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, manager)
	}
}

func UpdateManager(c *gin.Context) {
	pwd := c.Query("pwd")
	phone := c.Query("phone")
	name := c.Query("name")
	idcard := c.Query("idcard")
	department := c.Query("department")
	cap := c.Query("cap")
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	managers, _ := models.GetManagerById(id)
	manager := &models.Manager{Pwd: managers[0].Pwd, Phone: managers[0].Phone, Name: managers[0].Name, Idcard: managers[0].Idcard, Department: managers[0].Department, Cap: managers[0].Cap}
	manager.Id = int(id)
	c.ShouldBind(&manager)
	if pwd != "" {
		manager.Pwd = pwd
	}
	if department != "" {
		manager.Department = department
	}
	if phone != "" {
		manager.Phone = phone
	}
	if name != "" {
		manager.Name = name
	}
	if idcard != "" {
		manager.Idcard = idcard
	}
	if cap != "" {
		manager.Cap = cap
	}
	err = manager.UpdateManager()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, manager)
}

func DeleteManager(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	err := models.DeleteManager(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"id #" + Id: "deleted"})
}
