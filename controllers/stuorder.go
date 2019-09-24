package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStuorders(c *gin.Context) {
	stuorders, err := models.GetAllStuoders()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, stuorders)
	}
}

func GetStuorderByID(c *gin.Context) {
	WId := c.Params.ByName("wid")
	wid, _ := strconv.ParseInt(WId, 10, 64)
	stuorders, err := models.GetStuodersByWId(wid)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, stuorders)
	}
}

func CreateStuorder(c *gin.Context) {
	var stuorder *models.Stuorder
	c.BindJSON(&stuorder)
	if models.ExistsStuoderByWId(int64(stuorder.W_id)) == false {
		err := models.AddStuoders(stuorder)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			c.JSON(200, stuorder)
		}
	}
}

func UpdateStuorder(c *gin.Context) {
	w_flag, _ := strconv.Atoi(c.Query("w_flag"))
	WId := c.Param("wid")
	wid, err := strconv.ParseInt(WId, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	stuorders, _ := models.GetStuodersByWId(wid)
	stuorder := &models.Stuorder{CreatedAt: stuorders[0].CreatedAt, Question: stuorders[0].Question, Stu_id: stuorders[0].Stu_id, W_flag: stuorders[0].W_flag}
	stuorder.W_id = int(wid)
	c.ShouldBind(&stuorder)
	if w_flag >= 0 {
		stuorder.W_flag = w_flag
	}
	err = stuorder.UpdateStuoders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(200, stuorder)
}

func DeleteStuorder(c *gin.Context) {
	WId := c.Param("wid")
	wid, _ := strconv.ParseInt(WId, 10, 64)
	if models.ExistsStuoderByWId(wid) == true {
		err := models.DeleteStuoders(wid)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"wid #" + WId: "deleted"})
	}
}
