package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func readPlan() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/ilas?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	var studentplan [NODENUMBER]studentSchedule
	var teacherplan [NODENUMBER]teacherNode
	var student [NODENUMBER]studentInfo
	var teacher [NODENUMBER]teacherInfo
	db.Select("student_uid", "time_set", "teacher_list").Find(&studentplan)
	db.Select("teacher_class_id", "begin_time", "end_time").Find(&teacherplan)
}
