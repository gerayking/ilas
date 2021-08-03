package model

type Teacher struct {
	TeacherId int
	Sch       []Schedule
}

type Schedule struct {
	ClassId int
	Cla     *Class
}
