package service

import (
	"awesomeProject/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strconv"
	"strings"
	"time"
)

var teacherplan []model.TeacherSchedule
var studentplan []model.StudentSchedule

//var studentInfo []StudentInfo
//var teacherInfo []TeacherInfo

var Plan model.Plan
var Student model.Student
var Teacher model.TeacherPlan
var TeacherList []model.TeacherPlan
var StudentList []model.Student

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
	db.Limit(10000).Find(&studentplan)
	db.Limit(10000).Find(&teacherplan)
	//db.Find(&studentplan)
	//db.Find(&teacherplan)
}
func str2List(s string) {
	res := strings.Split(s, ",")
	for i := 0; i < len(res); i++ {
		res[i] = strings.Replace(res[i], "'", "", -1)
		res[i] = strings.Replace(res[i], " ", "", -1)
		res[i] = strings.Replace(res[i], "]", "", -1)
	}
	Plan.Class = res
	Student.Plans = append(Student.Plans, Plan)
}
func str2list(s string) []uint {
	s = s[1 : len(s)-1]
	sl := strings.Split(s, ",")
	for i := 0; i < len(sl); i++ {
		sl[i] = strings.Replace(sl[i], " ", "", -1)
	}
	var rl []uint
	for _, item := range sl {
		number, _ := strconv.Atoi(item)
		rl = append(rl, uint(number))
	}
	return rl
}
func dealDataStu() []model.Student {
	for i := 0; i < len(studentplan); i++ {
		Student.StuId = uint(studentplan[i].StudentUid)
		Student.Teachers = str2list(studentplan[i].TeacherList)
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
		StudentList = append(StudentList, Student)
		Student = model.Student{}
	}
	return StudentList
}

func enToNum(e string) int {
	arr := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for index, item := range arr {
		if item == e {
			return index
		}
	}
	return -1
}
func Datetodhm(begin time.Time, end time.Time) string {
	res := ""
	res += strconv.Itoa(enToNum(begin.Weekday().String())) + fmt.Sprintf("%d", begin.Hour()) + fmt.Sprintf("%02d", begin.Minute()) + fmt.Sprintf("%d", end.Hour()) + fmt.Sprintf("%02d", end.Minute())
	return res
}
func dealDataTea() []model.TeacherPlan {
	mp := make(map[int64]model.TeacherPlan)
	for i := 0; i < len(teacherplan); i++ {
		Teacher.TeacherId = teacherplan[i].TeacherUid
		Teacher.Schedule = append(Teacher.Schedule, Datetodhm(time.Unix(teacherplan[i].BeginTime, 0), time.Unix(teacherplan[i].EndTime, 0)))
		if _, ok := mp[teacherplan[i].TeacherUid]; ok {
			tmp := model.TeacherPlan{TeacherId: teacherplan[i].TeacherUid, Schedule: append(mp[teacherplan[i].TeacherUid].Schedule, Datetodhm(time.Unix(teacherplan[i].BeginTime, 0), time.Unix(teacherplan[i].EndTime, 0)))}
			mp[teacherplan[i].TeacherUid] = tmp
		} else {
			mp[teacherplan[i].TeacherUid] = Teacher
		}
		Teacher = model.TeacherPlan{}
	}
	TeacherList = TeacherList[0:0]
	for _, v := range mp {
		TeacherList = append(TeacherList, v)
	}
	return TeacherList
}

func CreateData() ([]model.Student, []model.TeacherPlan) {
	readPlan()
	return dealDataStu(), dealDataTea()
}
