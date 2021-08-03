package main

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"math"
	"strconv"
	"strings"
	"time"
)

const inf = math.MaxInt32

var teacherplan []TeacherSchedule
var studentplan []StudentSchedule

//var studentInfo []StudentInfo
//var teacherInfo []TeacherInfo

var Plan model.Plan
var Student model.Student
var Teacher model.TeacherSchedule

//var Schedule model.Schedule
var Plans []model.Plan

func readPlan() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/ilas?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	db.Find(&studentplan)
	db.Find(&teacherplan)

}
func str2List(s string) {
	res := strings.Split(s, ",")
	for i := 0; i < len(res); i++ {
		res[i] = strings.Replace(res[i], "'", "", -1)
		res[i] = strings.Replace(res[i], " ", "", -1)
	}
	Plan.Class = res
	Student.Plans = append(Student.Plans, Plan)
}
func dealDataStu() {
	for i := 0; i < 200; i++ {
		l1 := strings.Split(studentplan[i].TimeSet, "],")
		l1[0] = l1[0][3 : len(l1[0])-1]
		str2List(l1[0])
		if len(l1) > 1 {
			l1[len(l1)-1] = l1[len(l1)-1][3 : len(l1[len(l1)-1])-3]
			str2List(l1[len(l1)-1])
		}
		for i := 1; i < len(l1)-1; i++ {
			l1[i] = l1[i][2 : len(l1[i])-1]
			str2List(l1[i])
		}
	}

}

func enToNum(e string) int {
	arr := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for index, item := range arr {
		if item == e {
			return index + 1
		}
	}
	return -1
}
func datetodhm(begin time.Time, end time.Time) string {
	res := ""
	res += strconv.Itoa(enToNum(begin.Weekday().String())) + fmt.Sprintf("%02d", begin.Hour()) + fmt.Sprintf("%02d", begin.Minute()) + fmt.Sprintf("%02d", end.Hour()) + fmt.Sprintf("%02d", end.Minute())
	return res
}
func dealDataTea() {
	for i := 0; i < 10; i++ {
		Teacher.TeacherId = teacherplan[i].TeacherUid
		//Schedule.ClassId = teacherplan[i].TeacherClassId
		Schedule.Class = datetodhm(time.Unix(teacherplan[i].BeginTime, 0), time.Unix(teacherplan[i].EndTime, 0))
		Teacher.Sch = append(Teacher.Sch, Schedule)
	}
}

func main() {
	readPlan()
	dealDataStu()
	dealDataTea()

}
