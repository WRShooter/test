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

func GetAllCert_Cates() ([]*Cert_cate, error) {
	db := method.DBOpen()
	defer db.Close()
	var cert_cates []*Cert_cate
	err := db.Debug().Preload("Cert_cate_lists").Find(&cert_cates).Error
	return cert_cates, err
}

func GetCert_CateById(id int64) ([]*Cert_cate, error) {
	db := method.DBOpen()
	defer db.Close()
	var Cert_cates []*Cert_cate
	err := db.Preload("Cert_cate_lists").First(&Cert_cates, id).Error
	return Cert_cates, err
}

func ExistsCert_CateById(id int64) bool {
	db := method.DBOpen()
	defer db.Close()
	var Cert_cates []*Cert_cate
	if db.Preload("Cert_cate_lists").First(&Cert_cates, id).RowsAffected == 0 {
		return false
	}
	return true
}

func AddCert_Cate(cert_cate *Cert_cate) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Debug().Create(&cert_cate).Error
	db.Delete(&Cert_cate_list{}, "Cate_id=?", cert_cate.Id)
	db.Delete(&Cert_save{}, "Cert_id=?", cert_cate.Id)
	return err
}

func (cert_cate *Cert_cate) UpdateCert_Cate() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&cert_cate).Updates(map[string]interface{}{"id": cert_cate.Id, "name": cert_cate.Name}).Error
	return err
}

func DeleteCert_Cate(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Cert_cate{}, "id=?", id).Error
	return err
}
