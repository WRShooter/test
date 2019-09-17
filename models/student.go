package models

import (
	"fmt"
	"gormtest/method"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Student struct {
	Id           int        `gorm:"primary_key;size:11"`
	Pwd          string     `gorm:"size:15;not null"`
	Idcard       string     `gorm:"size:18;index:student_idcard;not null"`
	Name         string     `gorm:"size:30;not null"`
	Dept         string     `gorm:"size:20;null"`
	Grade        string     `gorm:"size:4;null"`
	Class        string     `gorm:"size:2;null"`
	Professional string     `gorm:"size:20;null"`
	Phone        string     `gorm:"size:11;null"`
	Stuorders    []Stuorder `gorm:"ForeignKey:Stu_id;AssociationForeignKey:Id"`
}

func (Student) TableName() string {
	return "student"
}

var db *gorm.DB

func GetAllStudents() []*Student {
	db := method.DBOpen()
	defer db.Close()
	var students []*Student
	db.Debug().Preload("Stuorders").Find(&students)
	return students
}

func GetStudentById(id int64) []*Student {
	db := method.DBOpen()
	defer db.Close()
	var students []*Student
	db.Preload("Stuorders").First(&students, id)
	return students
}

func AddStudent(student *Student) int {
	db := method.DBOpen()
	defer db.Close()
	fmt.Println(len(strconv.Itoa(student.Id)))
	if len(strconv.Itoa(student.Id)) != 10 {
		fmt.Println("ID的长度应为10，输入错误！！！")
		return 403
	} else if len(student.Idcard) != 18 {
		fmt.Println("IDcard的长度为18，输入错误！！！")
		return 403
	} else {
		db.Debug().Create(&student)
		db.Delete(&Stuorder{}, "Stu_id=?", student.Id)
		return 200
	}
}

func UpdateStudent(student *Student) {
	db := method.DBOpen()
	defer db.Close()
	db.Model(&student).Updates(map[string]interface{}{"id": student.Id, "idcard": student.Idcard, "name": student.Name, "grade": student.Grade, "class": student.Class, "dept": student.Dept, "phone": student.Phone, "professional": student.Professional, "pwd": student.Pwd})
}

func DeleteStudent(id int64) {
	db := method.DBOpen()
	defer db.Close()
	db.Delete(&Student{}, "id=?", id)
}
