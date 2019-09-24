package models

import (
	"gormtest/method"
)

type Cert_cate_list struct {
	Id              int    `gorm:"size:4;primary_key"`
	Name            string `gorm:"size:30;not null"`
	Abbreviation    string `gorm:"size:4;default:-1"`
	Cate_list_level string `gorm:"size:2;default:null"`
	Cate_id         int    `gorm:"size:4;index:Cert_cate_id;not null;default:-1"`
	Institution     string `gorm:"size:20;default:null"`
}

func (Cert_cate_list) TableName() string {
	return "Cert_cate_list"
}

func GetAllCert_Cate_Lists() ([]*Cert_cate_list, error) {
	db := method.DBOpen()
	defer db.Close()
	var cert_cate_lists []*Cert_cate_list
	err := db.Find(&cert_cate_lists).Error
	return cert_cate_lists, err
}

func GetCert_Cate_ListById(id int64) ([]*Cert_cate_list, error) {
	db := method.DBOpen()
	defer db.Close()
	var cert_cate_lists []*Cert_cate_list
	err := db.First(&cert_cate_lists, id).Error
	return cert_cate_lists, err
}

func AddCert_Cate_List(cert_cate_list *Cert_cate_list) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Debug().Create(&cert_cate_list).Error
	return err
}

func (cert_cate_list *Cert_cate_list) UpdateCert_Cate_List() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&cert_cate_list).Updates(map[string]interface{}{"id": cert_cate_list.Id, "cate_id": cert_cate_list.Cate_id, "name": cert_cate_list.Name, "institution": cert_cate_list.Institution, "cate_list_level": cert_cate_list.Cate_list_level, "abbreviation": cert_cate_list.Abbreviation}).Error
	return err
}

func DeleteCert_Cate_List(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Cert_cate_list{}, "id=?", id).Error
	return err
}
