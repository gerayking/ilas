package model

// 与数据库表信息一一对应
import (
	"gorm.io/gorm"
	"time"
)

type StudentInfo struct {
	gorm.Model
	Id         int
	StudentUid int64
	Name       string
	Sex        int8
	Birthday   time.Time
	CreateTime time.Time
	ModifyTime time.Time
}
type SeacherInfo struct {
	gorm.Model
	Id         int
	TeacherUid int64
	Name       string
	Sex        int8
	CreateTime time.Time
	ModifyTime time.Time
}

type TeacherSchedule struct {
	Id             int64
	TeacherUid     int64
	TeacherClassId int64
	BeginTime      int64
	EndTime        int64
}

type StudentSchedule struct {
	Id          int `gorm:"primaryKey"`
	StudentUid  int64
	TimeSet     string
	TeacherList string
	CreateTime  time.Time
}

func (StudentSchedule) TableName() string {
	return "plan"
}
