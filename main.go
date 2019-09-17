package main

import (
	"gormtest/config"
	"gormtest/method"
	"gormtest/models"
	_ "gormtest/routers"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	db = method.DBOpen()
	config.HttpConfig()
}

func main() {
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
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
