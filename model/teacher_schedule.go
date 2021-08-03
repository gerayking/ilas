package model

type TeacherSchedule struct {
	TeacherId int
	Schedule  []string // 这里存放class  class格式为 Day + starttime + endtime
}
