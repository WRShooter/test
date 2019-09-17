package method

import (
	"gormtest/config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/jinzhu/gorm"
)

func DBOpen() *gorm.DB {
	var dbCfg config.DBConfig = *((*config.GetConfig()).DB)
	dbURI := config.ReturnUrl()
	db, err := gorm.Open(dbCfg.Dialect, dbURI)
	if err != nil {
		panic(err)
	}
	return db
}

func TestCrossOrigin() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
}
