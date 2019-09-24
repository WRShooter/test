package models

import (
	"gormtest/method"
)

type Manager struct {
	Id         int    `gorm:"size:8;primary_key"`
	Pwd        string `gorm:"size:15;not null"`
	Idcard     string `gorm:"size:18;index:manager_idcard;not null"`
	Name       string `gorm:"size:30;not null"`
	Department string `gorm:"size:20"`
	Cap        string `gorm:"size:2"`
	Phone      string `gorm:"size:11"`
}

func (Manager) TableName() string {
	return "manager"
}

func GetAllManagers() ([]*Manager, error) {
	db := method.DBOpen()
	defer db.Close()
	var managers []*Manager
	err := db.Find(&managers).Error
	return managers, err
}

func GetManagerById(id int64) ([]*Manager, error) {
	db := method.DBOpen()
	defer db.Close()
	var managers []*Manager
	err := db.First(&managers, id).Error
	return managers, err
}

func ExistsManagerById(id int64) bool {
	db := method.DBOpen()
	defer db.Close()
	var managers []*Manager
	if db.First(&managers, id).RowsAffected == 0 {
		return false
	}
	return true
}

func AddManager(manager *Manager) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Create(&manager).Error
	return err
}

func (manager *Manager) UpdateManager() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&manager).Updates(map[string]interface{}{"id": manager.Id, "idcard": manager.Idcard, "name": manager.Name, "department": manager.Department, "cap": manager.Cap, "pwd": manager.Pwd, "phone": manager.Phone}).Error
	return err
}

func DeleteManager(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Manager{}, "id=?", id).Error
	return err
}
