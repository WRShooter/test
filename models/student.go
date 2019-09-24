package models

import (
	"gormtest/method"

	"github.com/jinzhu/gorm"
)

type Student struct {
	Id           int         `gorm:"primary_key;size:11"`
	Pwd          string      `gorm:"size:15;not null"`
	Idcard       string      `gorm:"size:18;index:student_idcard;not null"`
	Name         string      `gorm:"size:30;not null"`
	Dept         string      `gorm:"size:20;null"`
	Grade        string      `gorm:"size:4;null"`
	Class        string      `gorm:"size:2;null"`
	Professional string      `gorm:"size:20;null"`
	Phone        string      `gorm:"size:11;null"`
	Stuorders    []Stuorder  `gorm:"ForeignKey:Stu_id;AssociationForeignKey:Id"`
	Cert_saves   []Cert_save `gorm:"ForeignKey:stu_id;AssociationForeignKey:Id"`
}

func (Student) TableName() string {
	return "student"
}

var db *gorm.DB

func GetAllStudents() ([]*Student, error) {
	db := method.DBOpen()
	defer db.Close()
	var students []*Student
	err := db.Debug().Preload("Stuorders").Find(&students).Error
	return students, err
}

func GetStudentById(id int64) ([]*Student, error) {
	db := method.DBOpen()
	defer db.Close()
	var students []*Student
	err := db.Preload("Stuorders").First(&students, id).Error
	return students, err
}

func ExistsStudentById(id int64) bool {
	db := method.DBOpen()
	defer db.Close()
	var students []*Student
	if db.Preload("Stuorders").First(&students, id).RowsAffected == 0 {
		return false
	}
	return true
}

func AddStudent(student *Student) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Debug().Create(&student).Error
	db.Delete(&Stuorder{}, "Stu_id=?", student.Id)
	return err
}

func (student *Student) UpdateStudent() error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Model(&student).Updates(map[string]interface{}{"id": student.Id, "idcard": student.Idcard, "name": student.Name, "grade": student.Grade, "class": student.Class, "dept": student.Dept, "phone": student.Phone, "professional": student.Professional, "pwd": student.Pwd}).Error
	return err
}

func DeleteStudent(id int64) error {
	db := method.DBOpen()
	defer db.Close()
	err := db.Delete(&Student{}, "id=?", id).Error
	return err
}
