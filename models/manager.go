package models

import (
	"fmt"
	"gormtest/method"
	"strconv"
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

func GetAllManagers() []*Manager {
	db := method.DBOpen()
	defer db.Close()
	var managers []*Manager
	db.Find(&managers)
	return managers
}

func GetManagerById(id int64) []*Manager {
	db := method.DBOpen()
	defer db.Close()
	var managers []*Manager
	db.First(&managers, id)
	return managers
}

func AddManager(manager *Manager) int {
	db := method.DBOpen()
	defer db.Close()
	if len(strconv.Itoa(manager.Id)) != 8 {
		fmt.Println("ID的长度应为8，输入错误！！！")
		return 403
	} else if (len(manager.Idcard)) != 18 {
		fmt.Println("身份证的长度应该为18位！！！")
		return 403
	} else {
		db.Create(&manager)
		return 200
	}
}

func UpdateManager(manager *Manager) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&manager).Updates(map[string]interface{}{"id": manager.Id, "idcard": manager.Idcard, "name": manager.Name, "department": manager.Department, "cap": manager.Cap, "pwd": manager.Pwd, "phone": manager.Phone})
}

func DeleteManager(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Manager{}, "id=?", id)
}
