package models

import (
	"fmt"
	"gormtest/method"
	"time"
)

type Stuorder struct {
	W_id      int       `gorm:"size:11;primary_key"`
	Stu_id    int       `gorm:"size:11;null"`
	Question  []uint8   `gorm:"null"`
	CreatedAt time.Time `gorm:"null;column:w_date"`
	W_flag    int       `gorm:"size:2"`
}

func (Stuorder) TableName() string {
	return "stuorder"
}

func GetAllStuoders() []*Stuorder {
	db := method.DBOpen()
	defer db.Close()
	var stuorders []*Stuorder
	db.Find(&stuorders)
	return stuorders
}

func GetStuodersByWId(id int64) []*Stuorder {
	db := method.DBOpen()
	defer db.Close()
	var stuorders []*Stuorder
	db.First(&stuorders, id)
	return stuorders
}

func AddStuoders(stuorder *Stuorder) int {
	db := method.DBOpen()
	defer db.Close()
	var students []Student
	if db.First(&students, stuorder.Stu_id).RowsAffected == 0 {
		fmt.Println("该学生不存在！！！")
		return 403
	}
	db.Create(&stuorder)
	return 200
}

func UpdateStuoders(stuorder *Stuorder) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&stuorder).Updates(map[string]interface{}{"wid": stuorder.W_id, "stu_id": stuorder.Stu_id, "question": stuorder.Question, "w_date": stuorder.CreatedAt, "w_flag": stuorder.W_flag})
}

func DeleteStuoders(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Stuorder{}, "id=?", id)
}
