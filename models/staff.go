package models

import (
	"fmt"
	"gormtest/method"
	"strconv"
)

type Staff struct {
	Id       int    `gorm:"size:8;primary_key"`
	Pwd      string `gorm:"size:15;not null"`
	Idcard   string `gorm:"size:18;staff_idcard;not null"`
	Name     string `gorm:"size:30;not null"`
	Dept     string `gorm:"size:20;null"`
	Position string `gorm:"size:10"`
	Cap      string `gorm:"size:2"`
	Phone    string `gorm:"size:11"`
}

func (Staff) TableName() string {
	return "staff"
}

func GetAllStaffs() []*Staff {
	db := method.DBOpen()
	defer db.Close()
	var staffs []*Staff
	db.Find(&staffs)
	return staffs
}

func GetStaffById(id int64) []*Staff {
	db := method.DBOpen()
	defer db.Close()
	var staffs []*Staff
	db.First(&staffs, id)
	return staffs
}

func AddStaff(staff *Staff) int {
	db := method.DBOpen()
	defer db.Close()
	if len(strconv.Itoa(staff.Id)) != 8 {
		fmt.Println("ID的长度应为8，输入错误！！！")
		return 403
	} else if len(staff.Idcard) != 18 {
		fmt.Println("IDcard的长度为18，输入错误！！！")
		return 403
	} else {
		db.Create(&staff)
		return 200
	}
}

func UpdateStaff(staff *Staff) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&staff).Updates(map[string]interface{}{"id": staff.Id, "idcard": staff.Idcard, "name": staff.Name, "dept": staff.Dept, "phone": staff.Phone, "position": staff.Cap, "cap": staff.Cap})
}

func DeleteStaff(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Student{}, "id=?", id)
}
