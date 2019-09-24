package models

import (
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

func GetAllStuoders() ([]*Stuorder, error) {
	db := method.DBOpen()
	defer db.Close()
	var stuorders []*Stuorder
	err := db.Find(&stuorders).Error
	return stuorders, err
}

func GetStuodersByWId(wid int64) ([]*Stuorder, error) {
	db := method.DBOpen()
	defer db.Close()
	var stuorders []*Stuorder
	err := db.First(&stuorders, wid).Error
	return stuorders, err
}

func AddStuoders(stuorder *Stuorder) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Create(&stuorder).Error
	return err
}

func (stuorder *Stuorder) UpdateStuoders() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&stuorder).Updates(map[string]interface{}{"w_id": stuorder.W_id, "stu_id": stuorder.Stu_id, "question": stuorder.Question, "w_date": stuorder.CreatedAt, "w_flag": stuorder.W_flag}).Error
	return err
}

func DeleteStuoders(wid int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Stuorder{}, "w_id=?", wid).Error
	return err
}
