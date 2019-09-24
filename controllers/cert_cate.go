package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCert_Cates(c *gin.Context) {
	cert_cates, err := models.GetAllCert_Cates()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cates)
	}
}

func GetCert_CateByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	cert_cate, err := models.GetCert_CateById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cate)
	}
}

func CreateCert_Cate(c *gin.Context) {
	var cert_cate *models.Cert_cate
	c.BindJSON(&cert_cate)
	err := models.AddCert_Cate(cert_cate)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cate)
	}
}

func UpdateCert_Cate(c *gin.Context) {
	name := c.Query("name")
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	cert_cates, _ := models.GetCert_CateById(id)
	cert_cate := &models.Cert_cate{Name: cert_cates[0].Name}
	cert_cate.Id = int(id)
	c.ShouldBind(&cert_cate)
	if name != "" {
		cert_cate.Name = name
	}
	err = cert_cate.UpdateCert_Cate()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, cert_cate)
}

func DeleteCert_Cate(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	err := models.DeleteCert_Cate(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"id #" + Id: "deleted"})
}
