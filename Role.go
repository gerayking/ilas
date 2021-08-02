package main

import (
	"gorm.io/gorm"
	"time"
)

type studentInfo struct {
	gorm.Model
	id         int
	studentUid int64
	name       string
	sex        int8
	birthday   time.Time
	createTime time.Time
	modifyTime time.Time
}
type teacherInfo struct {
	gorm.Model
	id         int
	teacherUid int64
	name       string
	sex        int8
	createTime time.Time
	modifyTime time.Time
}

type teacherSchedule struct {
	id             int64
	teacherUid     int64
	teacherClassId int64
	beginTime      time.Time
	endTime        time.Time
}

type studentSchedule struct {
	gorm.Model
	id          int
	studentUid  int64
	timeSet     string
	teacherList []int64
	createTime  time.Time
}
