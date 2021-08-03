package model

type TeacherSchedule struct {
	TeacherId uint16
	Schedule  []string // 这里存放class  class格式为 Day + starttime + endtime
}
