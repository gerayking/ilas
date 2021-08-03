package model

type TeacherSchedule struct {
	TeacherId int64
	Schedule  []string // 这里存放class  class格式为 Day + starttime + endtime
}
