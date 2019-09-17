package models

import (
	"fmt"
	"gormtest/method"
	"time"
)

type Cert_save struct {
	Id        int       `gorm:"primary_key"`
	Stu_id    int       `gorm:"size:11;not null"`
	Cert_id   int       `gorm:"size:11;not null"`
	Num       string    `gorm:"size:4;not null"`
	Time      string    `gorm:"size:8;default:0;not null"`
	CreatedAt time.Time `gorm:"column:create_time;not null"`
	UpdatedAt time.Time `gorm:"column:modify_time;not null"`
	Input_id  int       `gorm:"8;not null"`
	Hash      string    `gorm:"not null"`
}

func (Cert_save) TableName() string {
	return "cert_save"
}

func GetAllCert_Saves() []*Cert_save {
	db := method.DBOpen()
	defer db.Close()
	var cert_saves []*Cert_save
	db.Debug().Preload("Cert_cate_lists").Find(&cert_saves)
	return cert_saves
}

func GetCert_SaveById(id int64) []*Cert_save {
	db := method.DBOpen()
	defer db.Close()
	var Cert_saves []*Cert_save
	db.Preload("Cert_cate_lists").First(&Cert_saves, id)
	return Cert_saves
}

func AddCert_Save(cert_save *Cert_save) int {
	db := method.DBOpen()
	defer db.Close()
	var cert_cates []*Cert_cate
	if db.First(&cert_cates, cert_save.Cert_id).RowsAffected == 0 {
		fmt.Println("该证书类别不存在！！！")
		return 403
	}
	db.Debug().Create(&cert_save)
	return 200
}

func UpdateCert_Save(cert_save *Cert_save) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&cert_save).Updates(map[string]interface{}{"id": cert_save.Id, "input_id": cert_save.Input_id, "num": cert_save.Num, "cert_id": cert_save.Cert_id, "time": cert_save.Time, "stu_id": cert_save.Stu_id})
}

func DeleteCert_Save(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Cert_save{}, "id=?", id)
}
