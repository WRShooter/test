package controllers

import (
	"fmt"
	"gormtest/models"
	"net/http"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	students, err := models.GetAllStudents()
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, students)
	}
}

func GetStudentByID(c *gin.Context) {
	Id := c.Params.ByName("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	student, err := models.GetStudentById(id)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, student)
	}
}

func CreateStudent(c *gin.Context) {
	pwd := c.Query("pwd")
	idcard := c.Query("idcard")
	name := c.Query("name")
	grade := c.Query("grade")
	class := c.Query("class")
	professional := c.Query("professional")
	phone := c.Query("phone")
	dept := c.Query("dept")
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	valid := validation.Validation{}
	valid.Length(strconv.Itoa(id), 10, "id").Message("学生长度为10！！！")
	valid.Length(idcard, 18, "idcard").Message("身份证长度为18！！！")
	if !valid.HasErrors() {
		if models.ExistsStudentById(int64(id)) == false {
			student := &models.Student{Id: id, Idcard: idcard, Name: name, Class: class, Grade: grade, Phone: phone, Professional: professional, Pwd: pwd, Dept: dept}
			err = models.AddStudent(student)
			if err != nil {
				c.AbortWithStatus(404)
				fmt.Println(err)
			} else {
				c.JSON(200, student)
			}
		}
	}
}

func UpdateStudent(c *gin.Context) {
	pwd := c.Query("pwd")
	idcard := c.Query("idcard")
	name := c.Query("name")
	grade := c.Query("grade")
	class := c.Query("class")
	professional := c.Query("professional")
	phone := c.Query("phone")
	dept := c.Query("dept")
	Id := c.Param("id")
	id, err := strconv.ParseInt(Id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		fmt.Println(err)
	}
	valid := validation.Validation{}
	valid.Length(idcard, 18, "idcard").Message("身份证长度为18！！！")
	if !valid.HasErrors() {
		students, _ := models.GetStudentById(id)
		student := &models.Student{Idcard: students[0].Idcard, Name: students[0].Name, Class: students[0].Class, Grade: students[0].Grade, Phone: students[0].Phone, Professional: students[0].Professional, Pwd: students[0].Pwd, Dept: students[0].Dept}
		student.Id = int(id)
		c.ShouldBind(&student)
		if pwd != "" {
			student.Pwd = pwd
		}
		if idcard != "" {
			student.Idcard = idcard
		}
		if name != "" {
			student.Name = name
		}
		if grade != "" {
			student.Grade = grade
		}
		if class != "" {
			student.Class = class
		}
		if professional != "" {
			student.Professional = professional
		}
		if phone != "" {
			student.Phone = phone
		}
		if dept != "" {
			student.Dept = dept
		}
		err = student.UpdateStudent()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(200, student)
	}
}

func DeleteStudent(c *gin.Context) {
	Id := c.Param("id")
	id, _ := strconv.ParseInt(Id, 10, 64)
	if models.ExistsStudentById(int64(id)) == true {
		err := models.DeleteStudent(id)
		if err != nil {
			c.AbortWithStatus(404)
			fmt.Println(err)
		}
		c.JSON(200, gin.H{"id #" + Id: "deleted"})
	}
}
