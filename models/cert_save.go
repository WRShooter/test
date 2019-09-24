package models

import (
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

func GetAllCert_Saves() ([]*Cert_save, error) {
	db := method.DBOpen()
	defer db.Close()
	var cert_saves []*Cert_save
	err := db.Debug().Preload("Cert_cate_lists").Find(&cert_saves).Error
	return cert_saves, err
}

func GetCert_SaveById(id int64) ([]*Cert_save, error) {
	db := method.DBOpen()
	defer db.Close()
	var Cert_saves []*Cert_save
	err := db.Preload("Cert_cate_lists").First(&Cert_saves, id).Error
	return Cert_saves, err
}

func ExistsCert_SaveById(id int64) bool {
	db := method.DBOpen()
	defer db.Close()
	var Cert_saves []*Cert_save
	if db.Preload("Cert_cate_lists").First(&Cert_saves, id).RowsAffected == 0 {
		return false
	}
	return true
}

func AddCert_Save(cert_save *Cert_save) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Debug().Create(&cert_save).Error
	return err
}

func (cert_save *Cert_save) UpdateCert_Save() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&cert_save).Updates(map[string]interface{}{"id": cert_save.Id, "input_id": cert_save.Input_id, "num": cert_save.Num, "cert_id": cert_save.Cert_id, "time": cert_save.Time, "stu_id": cert_save.Stu_id}).Error
	return err
}

func DeleteCert_Save(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Cert_save{}, "id=?", id).Error
	return err
}
