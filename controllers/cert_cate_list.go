package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllCert_Cate_Lists(c *gin.Context) {
	cert_cate_lists, err := models.GetAllCert_Cate_Lists()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cate_lists)
	}
}

func GetCert_Cate_ListByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	cert_cate_list, err := models.GetCert_Cate_ListById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cate_list)
	}
}

func CreateCert_Cate_List(c *gin.Context) {
	var cert_cate_list *models.Cert_cate_list
	c.BindJSON(&cert_cate_list)
	err := models.AddCert_Cate_List(cert_cate_list)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cert_cate_list)
	}
}

func UpdateCert_Cate_List(c *gin.Context) {
	name := c.Query("name")
	abbreviation := c.Query("abbreviation")
	institution := c.Query("institution")
	cate_list_level := c.Query("cate_list_level")
	cate_id, _ := strconv.Atoi(c.Query("cate_id"))
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	cert_cate_lists, _ := models.GetCert_Cate_ListById(id)
	cert_cate_list := &models.Cert_cate_list{Name: cert_cate_lists[0].Name, Institution: cert_cate_lists[0].Institution, Cate_list_level: cert_cate_lists[0].Cate_list_level, Cate_id: cert_cate_lists[0].Cate_id, Abbreviation: cert_cate_lists[0].Abbreviation}
	cert_cate_list.Id = int(id)
	c.ShouldBind(&cert_cate_list)
	if name != "" {
		cert_cate_list.Name = name
	}
	if abbreviation != "" {
		cert_cate_list.Abbreviation = abbreviation
	}
	if cate_id >= 0 {
		cert_cate_list.Cate_id = cate_id
	}
	if cate_list_level != "" {
		cert_cate_list.Cate_list_level = cate_list_level
	}
	if institution != "" {
		cert_cate_list.Institution = institution
	}
	err = cert_cate_list.UpdateCert_Cate_List()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, cert_cate_list)
}

func DeleteCert_Cate_List(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	err := models.DeleteCert_Cate_List(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"id #" + Id: "deleted"})
}
