package models

import "gormtest/method"

type Cert_cate struct {
	Id              int              `gorm:"size:4;primary_key"`
	Name            string           `gorm:"size:30;not null"`
	Cert_cate_lists []Cert_cate_list `gorm:"ForeignKey:cate_id;AssociationForeignKey:Id"`
	Cert_saves      []Cert_save      `gorm:"ForeignKey:cert_id;AssociationForeignKey:Id"`
}

func (Cert_cate) TableName() string {
	return "cert_cate"
}

func GetAllCert_Cates() []*Cert_cate {
	db := method.DBOpen()
	defer db.Close()
	var cert_cates []*Cert_cate
	db.Debug().Preload("Cert_cate_lists").Find(&cert_cates)
	return cert_cates
}

func GetCert_CateById(id int64) []*Cert_cate {
	db := method.DBOpen()
	defer db.Close()
	var Cert_cates []*Cert_cate
	db.Preload("Cert_cate_lists").First(&Cert_cates, id)
	return Cert_cates
}

func AddCert_Cate(cert_cate *Cert_cate) Cert_cate {
	db := method.DBOpen()
	defer db.Close()
	db.Debug().Create(&cert_cate)
	db.Delete(&Cert_cate_list{}, "Cate_id=?", cert_cate.Id)
	db.Delete(&Cert_save{}, "Cert_id=?", cert_cate.Id)
	return *cert_cate
}

func UpdateCert_Cate(cert_cate *Cert_cate) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&cert_cate).Updates(map[string]interface{}{"id": cert_cate.Id, "name": cert_cate.Name})
}

func DeleteCert_Cate(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Cert_cate{}, "id=?", id)
}
