package models

import (
	"gormtest/method"
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

func GetAllStaffs() ([]*Staff, error) {
	db := method.DBOpen()
	defer db.Close()
	var staffs []*Staff
	err := db.Find(&staffs).Error
	return staffs, err
}

func GetStaffById(id int64) ([]*Staff, error) {
	db := method.DBOpen()
	defer db.Close()
	var staffs []*Staff
	err := db.First(&staffs, id).Error
	return staffs, err
}

func ExistsStaffById(id int64) bool {
	db := method.DBOpen()
	defer db.Close()
	var staffs []*Staff
	if db.First(&staffs, id).RowsAffected == 0 {
		return false
	}
	return true
}

func AddStaff(staff *Staff) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Create(&staff).Error
	return err
}

func (staff *Staff) UpdateStaff() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&staff).Updates(map[string]interface{}{"id": staff.Id, "idcard": staff.Idcard, "name": staff.Name, "dept": staff.Dept, "phone": staff.Phone, "position": staff.Cap, "cap": staff.Cap}).Error
	return err
}

func DeleteStaff(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Student{}, "id=?", id).Error
	return err
}
