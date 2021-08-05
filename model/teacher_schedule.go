package model

// 老师计划
type TeacherPlan struct {
	TeacherId int64
	Schedule  []string // 这里存放class  class格式为 Day + starttime + endtime
}
