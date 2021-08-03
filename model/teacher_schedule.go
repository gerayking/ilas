package model

type Teacher struct {
	TeacherId int
	Sch       []Schedule
}

type Schedule struct {
	ClassId int  //课程的唯一ID
	Class   string  // class格式为 Day + starttime + endtime
}
