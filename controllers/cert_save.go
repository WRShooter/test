package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCert_Saves(c *gin.Context) {
	cert_saves, err := models.GetAllCert_Saves()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_saves)
	}
}

func GetCert_SaveByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	cert_save, err := models.GetCert_SaveById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_save)
	}
}

func CreateCert_Save(c *gin.Context) {
	var cert_save *models.Cert_save
	c.BindJSON(&cert_save)
	if models.ExistsCert_SaveById(int64(cert_save.Id)) == false {
		err := models.AddCert_Save(cert_save)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, cert_save)
		}
	}
}

func UpdateCert_Save(c *gin.Context) {
	time := c.Query("time")
	stu_id, _ := strconv.ParseInt(c.Query("stu_id"), 10, 64)
	num := c.Query("num")
	cert_id, _ := strconv.Atoi(c.Query("cert_id"))
	input_id, _ := strconv.Atoi(c.Query("input_id"))
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	cert_saves, _ := models.GetCert_SaveById(id)
	cert_save := &models.Cert_save{Time: cert_saves[0].Time, Stu_id: cert_saves[0].Stu_id, Num: cert_saves[0].Num, Input_id: cert_saves[0].Input_id, CreatedAt: cert_saves[0].CreatedAt}
	cert_save.Id = int(id)
	c.ShouldBind(&cert_save)
	if time != "" {
		cert_save.Time = time
	}
	if stu_id >= 0 {
		cert_save.Stu_id = int(stu_id)
	}
	if cert_id >= 0 {
		cert_save.Cert_id = cert_id
	}
	if num != "" {
		cert_save.Num = num
	}
	if input_id >= 0 {
		cert_save.Input_id = input_id
	}
	err = cert_save.UpdateCert_Save()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, cert_save)
}

func DeleteCert_Save(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	if models.ExistsCert_SaveById(id) == true {
		err := models.DeleteCert_Save(id)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"id #" + Id: "deleted"})
	}
}
