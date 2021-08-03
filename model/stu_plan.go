package model

type Student struct {
	StuId    uint
	Plans    []Plan
	Teachers []uint //备选老师的ID
}

type Plan struct {
	Status bool     //默认为false即可
	Class  []string // class格式为 Day + starttime + endtime
}
