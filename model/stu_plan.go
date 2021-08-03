package model

type student struct {
	StuId    int
	Plans    []plan
	Teachers []int //备选老师的ID
}

type plan struct {
	Status bool //默认为false即可
	Cla    []Class
}
