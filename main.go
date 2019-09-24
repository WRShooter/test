package main

import (
	"gormtest/controllers"
	"gormtest/method"
	"gormtest/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db = method.DBOpen()
}

func main() {
	router := gin.Default()
	defer db.Close()
	method.TestCrossOrigin()
	if !db.HasTable(&models.Student{}) {
		db.CreateTable(&models.Student{})
	}
	if !db.HasTable(&models.Stuorder{}) {
		db.CreateTable(&models.Stuorder{})
	}
	if !db.HasTable(&models.Staff{}) {
		db.CreateTable(&models.Staff{})
	}
	if !db.HasTable(&models.Manager{}) {
		db.CreateTable(&models.Manager{})
	}
	if !db.HasTable(&models.Cert_cate_list{}) {
		db.CreateTable(&models.Cert_cate_list{})
	}
	if !db.HasTable(&models.Cert_cate{}) {
		db.CreateTable(&models.Cert_cate{})
	}
	if !db.HasTable(&models.Cert_save{}) {
		db.CreateTable(&models.Cert_save{})
	}
	router.GET("/student", controllers.GetAllStudents)
	router.GET("/student/:id", controllers.GetStudentByID)
	router.POST("/student", controllers.CreateStudent)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("/student/:id", controllers.DeleteStudent)

	router.GET("/stuorder", controllers.GetAllStuorders)
	router.GET("/stuorder/:wid", controllers.GetStuorderByID)
	router.POST("/stuorder", controllers.CreateStuorder)
	router.PUT("/stuorder/:wid", controllers.UpdateStuorder)
	router.DELETE("/stuorder/:wid", controllers.DeleteStuorder)

	router.GET("/staff", controllers.GetAllStaffs)
	router.GET("/staff/:id", controllers.GetStaffByID)
	router.POST("/staff", controllers.CreateStaff)
	router.PUT("/staff/:id", controllers.UpdateStaff)
	router.DELETE("/staff/:id", controllers.DeleteStaff)

	router.GET("/manager", controllers.GetAllManagers)
	router.GET("/manager/:id", controllers.GetManagerByID)
	router.POST("/manager", controllers.CreateManager)
	router.PUT("/manager/:id", controllers.UpdateManager)
	router.DELETE("/manager/:id", controllers.DeleteManager)

	router.GET("/cert_save", controllers.GetAllCert_Saves)
	router.GET("/cert_save/:id", controllers.GetCert_SaveByID)
	router.POST("/cert_save", controllers.CreateCert_Save)
	router.PUT("/cert_save/:id", controllers.UpdateCert_Save)
	router.DELETE("/cert_save/:id", controllers.DeleteCert_Save)

	router.GET("/cert_cate", controllers.GetAllCert_Cates)
	router.GET("/cert_cate/:id", controllers.GetCert_CateByID)
	router.POST("/cert_cate", controllers.CreateCert_Cate)
	router.PUT("/cert_cate/:id", controllers.UpdateCert_Cate)
	router.DELETE("/cert_cate/:id", controllers.DeleteCert_Cate)

	router.GET("/cert_cate_list", controllers.GetAllCert_Cate_Lists)
	router.GET("/cert_cate_list/:id", controllers.GetCert_Cate_ListByID)
	router.POST("/cert_cate_list", controllers.CreateCert_Cate_List)
	router.PUT("/cert_cate_list/:id", controllers.UpdateCert_Cate_List)
	router.DELETE("/cert_cate_list/:id", controllers.DeleteCert_Cate_List)
	router.Run(":8080")
}
